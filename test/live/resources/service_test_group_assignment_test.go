package resource_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestServiceTestGroupAssignmentResource(t *testing.T) {
	const (
		groupName  = "tf_acceptance_test_service_test_group_assignment"
		group2Name = "tf_acceptance_test_service_test_group_assignment_two"
	)

	var (
		resourceIdBeforeRecreate string
		resourceIdAfterRecreate  string
	)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "` + config.ServiceTestName + `"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "` + config.ServiceTestUid + `"
					}

					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						config.ServiceTestUid,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupName).Id)
							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "uxi_service_test_group_assignment.my_service_test_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIdBeforeRecreate = rs.Primary.ID
						return util.CheckStateAgainstServiceTestGroupAssignment(
							t,
							"uxi_service_test_group_assignment.my_service_test_group_assignment",
							util.GetServiceTestGroupAssignment(rs.Primary.ID),
						)(s)
					},
				),
			},
			// ImportState
			{
				ResourceName:      "uxi_service_test_group_assignment.my_service_test_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "` + groupName + `"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "` + config.ServiceTestName + `"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from service_test/group to service_test/group_2
					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						config.ServiceTestUid,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(group2Name).Id)
							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "uxi_service_test_group_assignment.my_service_test_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIdAfterRecreate = rs.Primary.ID
						return util.CheckStateAgainstServiceTestGroupAssignment(
							t,
							"uxi_service_test_group_assignment.my_service_test_group_assignment",
							util.GetServiceTestGroupAssignment(rs.Primary.ID),
						)(s)
					},
					// Check that resource has been recreated
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"id",
						func(value string) error {
							assert.NotEqual(t, value, resourceIdBeforeRecreate)
							return nil
						},
					),
				),
			},
			// Remove serviceTests from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_service_test.my_service_test

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
		CheckDestroy: func(s *terraform.State) error {
			assert.Equal(t, util.GetGroupByName(groupName), nil)
			assert.Equal(t, util.GetGroupByName(group2Name), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIdBeforeRecreate), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIdAfterRecreate), nil)
			return nil
		},
	})
}
