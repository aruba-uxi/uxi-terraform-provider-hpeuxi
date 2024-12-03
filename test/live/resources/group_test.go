/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
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

	var resourceIDBeforeRecreate string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				// Node without parent (attached to root)
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "parent" {
					name = "` + groupNameParent + `"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(
						"hpeuxi_group.parent",
						"id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupNameParent).Id)

							return nil
						},
					),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.parent", "name", groupNameParent,
					),
					resource.TestCheckNoResourceAttr("hpeuxi_group.parent", "parent_group_id"),
				),
			},
			// ImportState
			{
				ResourceName:      "hpeuxi_group.parent",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update that does not trigger a recreate
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "parent" {
						name = "` + groupNameParentUpdated + `"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(
						"hpeuxi_group.parent",
						"id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupNameParentUpdated).Id)

							return nil
						},
					),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.parent",
						"name",
						groupNameParentUpdated,
					),
					resource.TestCheckNoResourceAttr(
						"hpeuxi_group.parent",
						"parent_group_id",
					),
				),
			},
			// Create nodes attached to non root node
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "parent" {
						name = "` + groupNameParentUpdated + `"
					}

					resource "hpeuxi_group" "child" {
						name            = "` + groupNameChild + `"
						parent_group_id = hpeuxi_group.parent.id
					}

					resource "hpeuxi_group" "grandchild" {
						name            = "` + groupNameGrandChild + `"
						parent_group_id = hpeuxi_group.child.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(
						"hpeuxi_group.child",
						"id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupNameChild).Id)

							return nil
						},
					),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.child",
						"name",
						groupNameChild,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_group.child",
						"parent_group_id",
						func(parentGroupId string) error {
							return checkGroupIsChildOfNode(parentGroupId, groupNameParentUpdated)
						},
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_group.grandchild",
						"id",
						func(value string) error {
							resourceIDBeforeRecreate = util.GetGroupByName(
								groupNameGrandChild,
							).Id
							assert.Equal(t, value, resourceIDBeforeRecreate)

							return nil
						},
					),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.grandchild",
						"name",
						groupNameGrandChild,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_group.grandchild",
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
					resource "hpeuxi_group" "parent" {
						name = "` + groupNameParentUpdated + `"
					}

					resource "hpeuxi_group" "child" {
						name            = "` + groupNameChild + `"
						parent_group_id = hpeuxi_group.parent.id
					}

					# move grandchild from child to parent
					resource "hpeuxi_group" "grandchild" {
						name            = "` + groupNameGrandChildMovedToParent + `"
						parent_group_id = hpeuxi_group.parent.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(
						"hpeuxi_group.grandchild",
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
						"hpeuxi_group.grandchild",
						"name",
						groupNameGrandChildMovedToParent,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_group.grandchild",
						"parent_group_id",
						func(parentGroupId string) error {
							return checkGroupIsChildOfNode(parentGroupId, groupNameParentUpdated)
						},
					),
					// Check that resource has been recreated
					resource.TestCheckResourceAttrWith(
						"hpeuxi_group.grandchild",
						"id",
						func(value string) error {
							assert.NotEqual(t, value, resourceIDBeforeRecreate)

							return nil
						},
					),
				),
			},
			// Update non root node group back to the root node by removing parent_group_id
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "parent" {
						name = "` + groupNameParentUpdated + `"
					}

					resource "hpeuxi_group" "child" {
						name            = "` + groupNameChild + `"
						parent_group_id = hpeuxi_group.parent.id
					}

					# move grandchild from parent to root
					resource "hpeuxi_group" "grandchild" {
						name = "` + groupNameGrandChildMovedToRoot + `"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"hpeuxi_group.grandchild",
						"name",
						groupNameGrandChildMovedToRoot,
					),
					resource.TestCheckNoResourceAttr(
						"hpeuxi_group.grandchild",
						"parent_group_id",
					),
				),
			},
			// Delete
			{
				Config: provider.ProviderConfig,
			},
		},
		CheckDestroy: func(s *terraform.State) error {
			assert.Nil(t, util.GetGroupByName(groupNameParent))
			assert.Nil(t, util.GetGroupByName(groupNameParentUpdated))
			assert.Nil(t, util.GetGroupByName(groupNameChild))
			assert.Nil(t, util.GetGroupByName(groupNameGrandChild))
			assert.Nil(t, util.GetGroupByName(groupNameGrandChildMovedToParent))
			assert.Nil(t, util.GetGroupByName(groupNameGrandChildMovedToRoot))

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
				resource "hpeuxi_group" "my_root_group" {
					name = "root"
				}

				import {
					to = hpeuxi_group.my_root_group
					id = "` + config.GroupIDRoot + `"
				}`,
				ExpectError: regexp.MustCompile(`The root group cannot be used as a resource`),
			},
		},
	})
}

func checkGroupIsChildOfNode(actualParentGroupID, expectedParentName string) error {
	expectedParentGroupID := util.GetGroupByName(expectedParentName).GetId()

	if expectedParentGroupID != actualParentGroupID {
		return fmt.Errorf(
			"expected \"%s\", but got \"%s\"",
			expectedParentGroupID,
			actualParentGroupID,
		)
	}

	return nil
}
