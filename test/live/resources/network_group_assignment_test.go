package resource_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
)

func TestNetworkGroupAssignmentResourceForWiredNetwork(t *testing.T) {
	const groupName = "tf_provider_acceptance_test_network_association_test"
	const group2Name = "tf_provider_acceptance_test_network_association_test_two"

	// Test Wired Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a network group assignment
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_wired_network" "my_network" {
						name = "` + config.WiredNetworkName + `"
					}

					import {
						to = uxi_wired_network.my_network
						id = "` + config.WiredNetworkUid + `"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wired_network.my_network.id
						group_id   = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WiredNetworkUid,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							st.Assert(t, group_id, util.GetGroupByName(groupName).Id)
							return nil
						},
					),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_network_group_assignment.my_network_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_wired_network" "my_network" {
						name = "` + config.WiredNetworkName + `"
					}

					import {
						to = uxi_wired_network.my_network
						id = "` + config.WiredNetworkUid + `"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from network/group to network/group_2
					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = uxi_wired_network.my_network.id
						group_id 		 = uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WiredNetworkUid,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							st.Assert(t, group_id, util.GetGroupByName(group2Name).Id)
							return nil
						},
					),
				),
			},
			// Delete network-group-assignments and remove networks from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_wired_network.my_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}

func TestNetworkGroupAssignmentResourceForWirelessNetwork(t *testing.T) {
	const groupName = "tf_provider_acceptance_test_network_association_test"
	const group2Name = "tf_provider_acceptance_test_network_association_test_two"

	// Test Wired Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a network group assignment
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_wireless_network" "my_network" {
						name = "` + config.WirelessNetworkName + `"
					}

					import {
						to = uxi_wireless_network.my_network
						id = "` + config.WirelessNetworkUid + `"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wireless_network.my_network.id
						group_id   = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WirelessNetworkUid,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							st.Assert(t, group_id, util.GetGroupByName(groupName).Id)
							return nil
						},
					),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_network_group_assignment.my_network_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_wireless_network" "my_network" {
						name = "` + config.WirelessNetworkName + `"
					}

					import {
						to = uxi_wireless_network.my_network
						id = "` + config.WirelessNetworkUid + `"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from network/group to network/group_2
					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = uxi_wireless_network.my_network.id
						group_id 		 = uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						config.WirelessNetworkUid,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						func(group_id string) error {
							st.Assert(t, group_id, util.GetGroupByName(group2Name).Id)
							return nil
						},
					),
				),
			},
			// Delete network-group-assignments and remove networks from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_wireless_network.my_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}
