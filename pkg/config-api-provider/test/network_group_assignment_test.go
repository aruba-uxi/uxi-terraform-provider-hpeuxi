package test

import (
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNetworkGroupAssignmentResource(t *testing.T) {

	// Test Wired Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a network group assignment
			{
				PreConfig: func() {
					// required for network import
					resources.GetWiredNetwork = func(uid string) resources.WiredNetworkResponseModel {
						return GenerateWiredNetworkResponseModel(uid, "")
					}

					// required for group create
					resources.CreateGroup = func(request resources.GroupCreateRequestModel) resources.GroupResponseModel {
						return GenerateGroupResponseModel("group_uid", "", "")
					}
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						return GenerateGroupResponseModel("group_uid", "", "")
					}

					// required for network group assignment create
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid", "")
					}
					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						return GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid", "")
					}
				},

				Config: providerConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						alias = "alias"
					}

					import {
						to = uxi_wired_network.my_network
						id = "network_uid"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wired_network.my_network.id
						group_id   = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "network_id", "network_uid"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "group_id", "group_uid"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "id", "network_group_assignment_uid"),
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
				PreConfig: func() {
					resources.GetWiredNetwork = func(uid string) resources.WiredNetworkResponseModel {
						if uid == "network_uid" {
							return GenerateWiredNetworkResponseModel(uid, "")
						} else {
							return GenerateWiredNetworkResponseModel(uid, "_2")
						}
					}

					// required for creating another group
					resources.CreateGroup = func(request resources.GroupCreateRequestModel) resources.GroupResponseModel {
						return GenerateGroupResponseModel("group_uid_2", "_2", "_2")
					}
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						if uid == "group_uid" {
							return GenerateGroupResponseModel("uid", "", "")
						} else {
							return GenerateGroupResponseModel(uid, "_2", "_2")
						}
					}

					// required for network group assignment create
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid_2", "_2")
					}
					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						if uid == "network_group_assignment_uid" {
							return GenerateNetworkGroupAssignmentResponse(uid, "")
						} else {
							return GenerateNetworkGroupAssignmentResponse(uid, "_2")
						}
					}
				},
				Config: providerConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						alias = "alias"
					}

					import {
						to = uxi_wired_network.my_network
						id = "network_uid"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "name_2"
						parent_group_id = "parent_uid_2"
					}

					resource "uxi_wired_network" "my_network_2" {
						alias = "alias_2"
					}

					import {
						to = uxi_wired_network.my_network_2
						id = "network_uid_2"
					}

					// the assignment update, updated from network/group to network_2/group_2
					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = uxi_wired_network.my_network_2.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "network_id", "network_uid_2"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "group_id", "group_uid_2"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "id", "network_group_assignment_uid_2"),
				),
			},
			// Remove networks from state
			{
				Config: providerConfig + `
					removed {
						from = uxi_wired_network.my_network

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = uxi_wired_network.my_network_2

						lifecycle {
							destroy = false
						}
					}`,
			},
			// Delete testing automatically occurs in TestCase
		},
	})

	// Test Wireless Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a network group assignment
			{
				PreConfig: func() {
					// required for network import
					resources.GetWirelessNetwork = func(uid string) resources.WirelessNetworkResponseModel {
						return GenerateWirelessNetworkResponseModel(uid, "")
					}

					// required for group create
					resources.CreateGroup = func(request resources.GroupCreateRequestModel) resources.GroupResponseModel {
						return GenerateGroupResponseModel("group_uid", "", "")
					}
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						return GenerateGroupResponseModel(uid, "", "")
					}

					// required for network group assignment create
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid", "")
					}
					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						return GenerateNetworkGroupAssignmentResponse(uid, "")
					}
				},

				Config: providerConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wireless_network" "my_network" {
						alias = "alias"
					}

					import {
						to = uxi_wireless_network.my_network
						id = "network_uid"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wireless_network.my_network.id
						group_id   = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "network_id", "network_uid"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "group_id", "group_uid"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "id", "network_group_assignment_uid"),
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
				PreConfig: func() {
					resources.GetWirelessNetwork = func(uid string) resources.WirelessNetworkResponseModel {
						if uid == "network_uid" {
							return GenerateWirelessNetworkResponseModel(uid, "")
						} else {
							return GenerateWirelessNetworkResponseModel(uid, "_2")
						}
					}

					// required for creating another group
					resources.CreateGroup = func(request resources.GroupCreateRequestModel) resources.GroupResponseModel {
						return GenerateGroupResponseModel("group_uid_2", "_2", "_2")
					}
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						if uid == "group_uid" {
							return GenerateGroupResponseModel(uid, "", "")
						} else {
							return GenerateGroupResponseModel("group_uid_2", "_2", "_2")
						}
					}

					// required for network group assignment create
					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						if uid == "network_group_assignment_uid" {
							return GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid", "")
						} else {
							return GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid_2", "_2")
						}
					}
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid_2", "_2")
					}
				},
				Config: providerConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wireless_network" "my_network" {
						alias = "alias"
					}

					import {
						to = uxi_wireless_network.my_network
						id = "network_uid"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "name_2"
						parent_group_id = "parent_uid_2"
					}

					resource "uxi_wireless_network" "my_network_2" {
						alias = "alias_2"
					}

					import {
						to = uxi_wireless_network.my_network_2
						id = "network_uid_2"
					}

					// the assignment update, updated from network/group to network_2/group_2
					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = uxi_wireless_network.my_network_2.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "network_id", "network_uid_2"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "group_id", "group_uid_2"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "id", "network_group_assignment_uid_2"),
				),
			},
			// Remove networks from state
			{
				Config: providerConfig + `
					removed {
						from = uxi_wireless_network.my_network

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = uxi_wireless_network.my_network_2

						lifecycle {
							destroy = false
						}
					}`,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
