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

func TestNetworkGroupAssignmentResourceForWiredNetwork(t *testing.T) {
	const (
		groupName  = "tf_provider_acceptance_test_network_assignment_resource"
		group2Name = "tf_provider_acceptance_test_network_assignment_resource_two"
	)

	var (
		resourceIDBeforeRecreate string
		resourceIDAfterRecreate  string
	)

	// Test Wired Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "hpeuxi_wired_network" "my_network" {
						name = "` + config.WiredNetworkName + `"
					}

					import {
						to = hpeuxi_wired_network.my_network
						id = "` + config.WiredNetworkID + `"
					}

					resource "hpeuxi_network_group_assignment" "my_network_group_assignment" {
						network_id = hpeuxi_wired_network.my_network.id
						group_id   = hpeuxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"hpeuxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WiredNetworkID,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							assert.Equal(t, group_id, util.GetGroupByName(groupName).Id)

							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "hpeuxi_network_group_assignment.my_network_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIDBeforeRecreate = rs.Primary.ID

						return util.CheckStateAgainstNetworkGroupAssignment(
							t,
							"hpeuxi_network_group_assignment.my_network_group_assignment",
							util.GetNetworkGroupAssignment(resourceIDBeforeRecreate),
						)(s)
					},
				),
			},
			// ImportState
			{
				ResourceName:      "hpeuxi_network_group_assignment.my_network_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "hpeuxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "hpeuxi_wired_network" "my_network" {
						name = "` + config.WiredNetworkName + `"
					}

					import {
						to = hpeuxi_wired_network.my_network
						id = "` + config.WiredNetworkID + `"
					}

					// the new resources we wanna update the assignment to
					resource "hpeuxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from network/group to network/group_2
					resource "hpeuxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = hpeuxi_wired_network.my_network.id
						group_id 		 = hpeuxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"hpeuxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WiredNetworkID,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							assert.Equal(t, group_id, util.GetGroupByName(group2Name).Id)

							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "hpeuxi_network_group_assignment.my_network_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIDAfterRecreate = rs.Primary.ID

						return util.CheckStateAgainstNetworkGroupAssignment(
							t,
							"hpeuxi_network_group_assignment.my_network_group_assignment",
							util.GetNetworkGroupAssignment(resourceIDAfterRecreate),
						)(s)
					},
					// Check that resource has been recreated
					resource.TestCheckResourceAttrWith(
						"hpeuxi_network_group_assignment.my_network_group_assignment",
						"id",
						func(value string) error {
							assert.NotEqual(t, value, resourceIDBeforeRecreate)

							return nil
						},
					),
				),
			},
			// Delete network-group-assignments and remove networks from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_wired_network.my_network

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

func TestNetworkGroupAssignmentResourceForWirelessNetwork(t *testing.T) {
	const (
		groupName  = "tf_provider_acceptance_test_network_assignment_test"
		group2Name = "tf_provider_acceptance_test_network_assignment_test_two"
	)

	var (
		resourceIDBeforeRecreate string
		resourceIDAfterRecreate  string
	)

	// Test Wired Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "hpeuxi_wireless_network" "my_network" {
						name = "` + config.WirelessNetworkName + `"
					}

					import {
						to = hpeuxi_wireless_network.my_network
						id = "` + config.WirelessNetworkID + `"
					}

					resource "hpeuxi_network_group_assignment" "my_network_group_assignment" {
						network_id = hpeuxi_wireless_network.my_network.id
						group_id   = hpeuxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"hpeuxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WirelessNetworkID,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							assert.Equal(t, group_id, util.GetGroupByName(groupName).Id)

							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "hpeuxi_network_group_assignment.my_network_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIDBeforeRecreate = rs.Primary.ID

						return util.CheckStateAgainstNetworkGroupAssignment(
							t,
							"hpeuxi_network_group_assignment.my_network_group_assignment",
							util.GetNetworkGroupAssignment(resourceIDBeforeRecreate),
						)(s)
					},
				),
			},
			// ImportState
			{
				ResourceName:      "hpeuxi_network_group_assignment.my_network_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "hpeuxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "hpeuxi_wireless_network" "my_network" {
						name = "` + config.WirelessNetworkName + `"
					}

					import {
						to = hpeuxi_wireless_network.my_network
						id = "` + config.WirelessNetworkID + `"
					}

					// the new resources we wanna update the assignment to
					resource "hpeuxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from network/group to network/group_2
					resource "hpeuxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = hpeuxi_wireless_network.my_network.id
						group_id 		 = hpeuxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"hpeuxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WirelessNetworkID,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							assert.Equal(t, group_id, util.GetGroupByName(group2Name).Id)

							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "hpeuxi_network_group_assignment.my_network_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIDAfterRecreate = rs.Primary.ID

						return util.CheckStateAgainstNetworkGroupAssignment(
							t,
							"hpeuxi_network_group_assignment.my_network_group_assignment",
							util.GetNetworkGroupAssignment(resourceIDAfterRecreate),
						)(s)
					},
					// Check that resource has been recreated
					resource.TestCheckResourceAttrWith(
						"hpeuxi_network_group_assignment.my_network_group_assignment",
						"id",
						func(value string) error {
							assert.NotEqual(t, value, resourceIDBeforeRecreate)

							return nil
						},
					),
				),
			},
			// Delete network-group-assignments and remove networks from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_wireless_network.my_network

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
