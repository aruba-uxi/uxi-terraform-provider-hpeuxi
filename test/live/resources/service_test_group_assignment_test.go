/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
)

func TestServiceTestGroupAssignmentResource(t *testing.T) {
	const (
		groupName  = "tf_acceptance_test_service_test_group_assignment_resource"
		group2Name = "tf_acceptance_test_service_test_group_assignment_resource_two"
	)

	var (
		resourceIDBeforeRecreate string
		resourceIDAfterRecreate  string
	)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "hpeuxi_service_test" "my_service_test" {
						name = "` + config.ServiceTestName + `"
					}

					import {
						to = hpeuxi_service_test.my_service_test
						id = "` + config.ServiceTestID + `"
					}

					resource "hpeuxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = hpeuxi_service_test.my_service_test.id
						group_id 		= hpeuxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"hpeuxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						config.ServiceTestID,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupName).Id)

							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "hpeuxi_service_test_group_assignment.my_service_test_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIDBeforeRecreate = rs.Primary.ID

						return util.CheckStateAgainstServiceTestGroupAssignment(
							t,
							"hpeuxi_service_test_group_assignment.my_service_test_group_assignment",
							util.GetServiceTestGroupAssignment(rs.Primary.ID),
						)(s)
					},
				),
			},
			// ImportState
			{
				ResourceName:      "hpeuxi_service_test_group_assignment.my_service_test_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "hpeuxi_group" "my_group" {
						name            = "` + groupName + `"
					}

					resource "hpeuxi_service_test" "my_service_test" {
						name = "` + config.ServiceTestName + `"
					}

					// the new resources we wanna update the assignment to
					resource "hpeuxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from service_test/group to service_test/group_2
					resource "hpeuxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = hpeuxi_service_test.my_service_test.id
						group_id 		= hpeuxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"hpeuxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						config.ServiceTestID,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(group2Name).Id)

							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "hpeuxi_service_test_group_assignment.my_service_test_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIDAfterRecreate = rs.Primary.ID

						return util.CheckStateAgainstServiceTestGroupAssignment(
							t,
							"hpeuxi_service_test_group_assignment.my_service_test_group_assignment",
							util.GetServiceTestGroupAssignment(rs.Primary.ID),
						)(s)
					},
					// Check that resource has been recreated
					resource.TestCheckResourceAttrWith(
						"hpeuxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						func(value string) error {
							assert.NotEqual(t, value, resourceIDBeforeRecreate)

							return nil
						},
					),
				),
			},
			// Remove serviceTests from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_service_test.my_service_test

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
		CheckDestroy: func(s *terraform.State) error {
			assert.Equal(t, util.GetGroupByName(groupName), nil)
			assert.Equal(t, util.GetGroupByName(group2Name), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIDBeforeRecreate), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIDAfterRecreate), nil)

			return nil
		},
	})
}
