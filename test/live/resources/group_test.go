package resource_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

type Fetcher interface {
	FetchData() ([]byte, error)
}

func TestGroupResource(t *testing.T) {
	const (
		groupNameParent                  = "tf_provider_acceptance_test_group_resource_parent"
		groupNameParentUpdated           = groupNameParent + "_updated"
		groupNameChild                   = "tf_provider_acceptance_test_group_resource__child"
		groupNameGrandChild              = "tf_provider_acceptance_test_group_resource__grandchild"
		groupNameGrandChildMovedToParent = groupNameGrandChild + "_moved_to_parent"
		groupNameGrandChildMovedToRoot   = groupNameGrandChild + "_moved_to_root"
	)

	var resourceIdBeforeRecreate string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				// Node without parent (attached to root)
				Config: provider.ProviderConfig + `
				resource "uxi_group" "parent" {
					name = "` + groupNameParent + `"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(
						"uxi_group.parent",
						"id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupNameParent).Id)
							return nil
						},
					),
					resource.TestCheckResourceAttr(
						"uxi_group.parent", "name", groupNameParent,
					),
					resource.TestCheckNoResourceAttr("uxi_group.parent", "parent_group_id"),
				),
			},
			// ImportState
			{
				ResourceName:      "uxi_group.parent",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update that does not trigger a recreate
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "parent" {
						name = "` + groupNameParentUpdated + `"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(
						"uxi_group.parent",
						"id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupNameParentUpdated).Id)
							return nil
						},
					),
					resource.TestCheckResourceAttr(
						"uxi_group.parent",
						"name",
						groupNameParentUpdated,
					),
					resource.TestCheckNoResourceAttr(
						"uxi_group.parent",
						"parent_group_id",
					),
				),
			},
			// Create nodes attached to non root node
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "parent" {
						name = "` + groupNameParentUpdated + `"
					}

					resource "uxi_group" "child" {
						name            = "` + groupNameChild + `"
						parent_group_id = uxi_group.parent.id
					}

					resource "uxi_group" "grandchild" {
						name            = "` + groupNameGrandChild + `"
						parent_group_id = uxi_group.child.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(
						"uxi_group.child",
						"id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupNameChild).Id)
							return nil
						},
					),
					resource.TestCheckResourceAttr(
						"uxi_group.child",
						"name",
						groupNameChild,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_group.child",
						"parent_group_id",
						func(parentGroupId string) error {
							return checkGroupIsChildOfNode(parentGroupId, groupNameParentUpdated)
						},
					),
					resource.TestCheckResourceAttrWith(
						"uxi_group.grandchild",
						"id",
						func(value string) error {
							resourceIdBeforeRecreate = util.GetGroupByName(
								groupNameGrandChild,
							).Id
							assert.Equal(t, value, resourceIdBeforeRecreate)
							return nil
						},
					),
					resource.TestCheckResourceAttr(
						"uxi_group.grandchild",
						"name",
						groupNameGrandChild,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_group.grandchild",
						"parent_group_id",
						func(parentGroupId string) error {
							return checkGroupIsChildOfNode(parentGroupId, groupNameChild)
						},
					),
				),
			},
			// Update that does trigger a recreate (moving group)
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "parent" {
						name = "` + groupNameParentUpdated + `"
					}

					resource "uxi_group" "child" {
						name            = "` + groupNameChild + `"
						parent_group_id = uxi_group.parent.id
					}

					# move grandchild from child to parent
					resource "uxi_group" "grandchild" {
						name            = "` + groupNameGrandChildMovedToParent + `"
						parent_group_id = uxi_group.parent.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(
						"uxi_group.grandchild",
						"id",
						func(value string) error {
							assert.Equal(
								t,
								value,
								util.GetGroupByName(groupNameGrandChildMovedToParent).Id,
							)
							return nil
						},
					),
					resource.TestCheckResourceAttr(
						"uxi_group.grandchild",
						"name",
						groupNameGrandChildMovedToParent,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_group.grandchild",
						"parent_group_id",
						func(parentGroupId string) error {
							return checkGroupIsChildOfNode(parentGroupId, groupNameParentUpdated)
						},
					),
					// Check that resource has been recreated
					resource.TestCheckResourceAttrWith(
						"uxi_group.grandchild",
						"id",
						func(value string) error {
							assert.NotEqual(t, value, resourceIdBeforeRecreate)
							return nil
						},
					),
				),
			},
			// Update non root node group back to the root node by removing parent_group_id
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "parent" {
						name = "` + groupNameParentUpdated + `"
					}

					resource "uxi_group" "child" {
						name            = "` + groupNameChild + `"
						parent_group_id = uxi_group.parent.id
					}

					# move grandchild from parent to root
					resource "uxi_group" "grandchild" {
						name = "` + groupNameGrandChildMovedToRoot + `"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_group.grandchild",
						"name",
						groupNameGrandChildMovedToRoot,
					),
					resource.TestCheckNoResourceAttr(
						"uxi_group.grandchild",
						"parent_group_id",
					),
				),
			},
			// Delete
			{
				Config:  provider.ProviderConfig,
				Destroy: true,
			},
		},
		CheckDestroy: func(s *terraform.State) error {
			assert.Equal(t, util.GetGroupByName(groupNameParent), nil)
			assert.Equal(t, util.GetGroupByName(groupNameParentUpdated), nil)
			assert.Equal(t, util.GetGroupByName(groupNameChild), nil)
			assert.Equal(t, util.GetGroupByName(groupNameGrandChild), nil)
			assert.Equal(t, util.GetGroupByName(groupNameGrandChildMovedToParent), nil)
			assert.Equal(t, util.GetGroupByName(groupNameGrandChildMovedToRoot), nil)
			return nil
		},
	})
}

func TestRootGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Importing the root group does not work
			{
				Config: provider.ProviderConfig + `
				resource "uxi_group" "my_root_group" {
					name = "root"
				}

				import {
					to = uxi_group.my_root_group
					id = "` + config.GroupUidRoot + `"
				}`,
				ExpectError: regexp.MustCompile(`The root group cannot be used as a resource`),
			},
		},
	})
}

func checkGroupIsChildOfNode(actualParentGroupId, expectedParentName string) error {
	expectedParentGroupId := util.GetGroupByName(expectedParentName).GetId()

	if expectedParentGroupId != actualParentGroupId {
		return fmt.Errorf(
			"expected \"%s\", but got \"%s\"",
			expectedParentGroupId,
			actualParentGroupId,
		)
	}

	return nil
}
