package resource_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-configuration/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-configuration/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

type Fetcher interface {
	FetchData() ([]byte, error)
}

var rootGroup = util.GetRoot()

func TestGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				// Node without parent (attached to root)
				Config: provider.ProviderConfig + `
				resource "uxi_group" "parent" {
					name = "tf_provider_acceptance_test_parent"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_group.parent",
						"name",
						"tf_provider_acceptance_test_parent",
					),
					resource.TestCheckResourceAttrPtr(
						"uxi_group.parent",
						"parent_group_id",
						nil,
					),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_group.parent",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update that does not trigger a recreate
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "parent" {
						name            = "tf_provider_acceptance_test_parent_name_updated"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_group.parent",
						"name",
						"tf_provider_acceptance_test_parent_name_updated",
					),
					resource.TestCheckResourceAttrPtr(
						"uxi_group.parent",
						"parent_group_id",
						&rootGroup.Id,
					),
				),
			},
			// Create nodes attached to non root node
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "parent" {
						name            = "tf_provider_acceptance_test_parent_name_updated"
					}

					resource "uxi_group" "child" {
						name            = "tf_provider_acceptance_test_child"
						parent_group_id = uxi_group.parent.id
					}

					resource "uxi_group" "grandchild" {
						name            = "tf_provider_acceptance_test_grandchild"
						parent_group_id = uxi_group.child.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_group.child",
						"name",
						"tf_provider_acceptance_test_child",
					),
					resource.TestCheckResourceAttrWith(
						"uxi_group.child",
						"parent_group_id",
						func(parentGroupId string) error {
							return checkGroupIsChildOfNode(
								parentGroupId,
								"tf_provider_acceptance_test_parent_name_updated",
							)
						},
					),
					resource.TestCheckResourceAttr(
						"uxi_group.grandchild",
						"name",
						"tf_provider_acceptance_test_grandchild",
					),
					resource.TestCheckResourceAttrWith(
						"uxi_group.grandchild",
						"parent_group_id",
						func(parentGroupId string) error {
							return checkGroupIsChildOfNode(
								parentGroupId, "tf_provider_acceptance_test_child",
							)
						},
					),
				),
			},
			// Update that does trigger a recreate (moving group)
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "parent" {
						name            = "tf_provider_acceptance_test_parent_name_updated"
					}

					resource "uxi_group" "child" {
						name            = "tf_provider_acceptance_test_child"
						parent_group_id = uxi_group.parent.id
					}

					# move grandchild from child to parent
					resource "uxi_group" "grandchild" {
						name            = "tf_provider_acceptance_test_grandchild_moved_to_parent"
						parent_group_id = uxi_group.parent.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_group.grandchild",
						"name",
						"tf_provider_acceptance_test_grandchild_moved_to_parent",
					),
					resource.TestCheckResourceAttrWith(
						"uxi_group.grandchild",
						"parent_group_id",
						func(parentGroupId string) error {
							return checkGroupIsChildOfNode(
								parentGroupId,
								"tf_provider_acceptance_test_parent_name_updated",
							)
						},
					),
				),
			},
			// Update non root node group back to the root node by removing parent_group_id
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "parent" {
						name            = "tf_provider_acceptance_test_parent_name_updated"
					}

					resource "uxi_group" "child" {
						name            = "tf_provider_acceptance_test_child"
						parent_group_id = uxi_group.parent.id
					}

					# move grandchild from parent to root
					resource "uxi_group" "grandchild" {
						name            = "tf_provider_acceptance_test_grandchild_moved_to_root"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_group.grandchild",
						"name",
						"tf_provider_acceptance_test_grandchild_moved_to_root",
					),
					resource.TestCheckResourceAttr(
						"uxi_group.grandchild",
						"parent_group_id",
						rootGroup.Id,
					),
				),
			},
			// Deletes happen automatically
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
					id = "` + rootGroup.Id + `"
				}`,
				ExpectError: regexp.MustCompile(`The root group cannot be used as a resource`),
			},
		},
	})
}

func checkGroupIsChildOfNode(actualParentGroupId, expectedParentName string) error {
	expectedParentGroupId := util.GetNodeByName(expectedParentName).GetId()

	if expectedParentGroupId != actualParentGroupId {
		return fmt.Errorf(
			"expected \"%s\", but got \"%s\"",
			expectedParentGroupId,
			actualParentGroupId,
		)
	}

	return nil
}
