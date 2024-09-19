package test

import (
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNetworkGroupAssignmentResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	// Test Wired Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a network group assignment
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid",
						util.GenerateWiredNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWiredNetworkResponse("network_uid", "")}),
						2,
					)

					// required for group create
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")), 1)
					util.MockGetGroup("group_uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						}),
						2,
					)

					// required for network group assignment create
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return util.GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid", "")
					}
					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						return util.GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid", "")
					}
				},

				Config: provider.ProviderConfig + `
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
					util.MockGetWiredNetwork(
						"network_uid_2",
						util.GenerateWiredNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWiredNetworkResponse("network_uid_2", "_2")}),
						2,
					)
					util.MockGetWiredNetwork(
						"network_uid",
						util.GenerateWiredNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWiredNetworkResponse("network_uid", "")}),
						2,
					)
					util.MockGetGroup("group_uid_2", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")),
						}),
						1,
					)
					util.MockGetGroup("group_uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						}),
						3,
					)

					// required for creating another group
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")), 1)

					// required for network group assignment create
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return util.GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid_2", "_2")
					}
					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						if uid == "network_group_assignment_uid" {
							return util.GenerateNetworkGroupAssignmentResponse(uid, "")
						} else {
							return util.GenerateNetworkGroupAssignmentResponse(uid, "_2")
						}
					}
				},
				Config: provider.ProviderConfig + `
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
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid_2",
						util.GenerateWiredNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWiredNetworkResponse("network_uid_2", "_2")}),
						1,
					)
					util.MockGetWiredNetwork(
						"network_uid",
						util.GenerateWiredNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWiredNetworkResponse("network_uid", "")}),
						1,
					)
					util.MockGetGroup("group_uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						}),
						2,
					)
					util.MockGetGroup("group_uid_2", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")),
						}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
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
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a network group assignment
			{
				PreConfig: func() {
					// required for network import
					util.MockGetWirelessNetwork(
						"network_uid",
						util.GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWirelessNetworkResponse("network_uid", "")}),
						2,
					)
					// required for group create
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")), 1)
					util.MockGetGroup("group_uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						}),
						1,
					)

					// required for network group assignment create
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return util.GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid", "")
					}
					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						return util.GenerateNetworkGroupAssignmentResponse(uid, "")
					}
				},

				Config: provider.ProviderConfig + `
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
					util.MockGetWirelessNetwork(
						"network_uid_2",
						util.GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWirelessNetworkResponse("network_uid_2", "_2")}),
						2,
					)
					util.MockGetWirelessNetwork(
						"network_uid",
						util.GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWirelessNetworkResponse("network_uid", "")}),
						2,
					)

					// required for creating another group
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")), 1)
					util.MockGetGroup("group_uid_2", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")),
						}),
						1,
					)

					util.MockGetGroup("group_uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						}),
						3,
					)

					// required for network group assignment create
					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						if uid == "network_group_assignment_uid" {
							return util.GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid", "")
						} else {
							return util.GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid_2", "_2")
						}
					}
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return util.GenerateNetworkGroupAssignmentResponse("network_group_assignment_uid_2", "_2")
					}
				},
				Config: provider.ProviderConfig + `
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
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"network_uid_2",
						util.GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWirelessNetworkResponse("network_uid_2", "_2")}),
						1,
					)
					util.MockGetWirelessNetwork(
						"network_uid",
						util.GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWirelessNetworkResponse("network_uid", "")}),
						1,
					)
					util.MockGetGroup("group_uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						}),
						2,
					)
					util.MockGetGroup("group_uid_2", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")),
						}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
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

	mockOAuth.Mock.Disable()
}
