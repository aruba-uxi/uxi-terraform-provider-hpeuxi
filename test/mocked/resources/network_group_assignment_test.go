package resource_test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNetworkGroupAssignmentResourceForWiredNetwork(t *testing.T) {
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
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						),
						1,
					)
					util.MockGetGroup("group_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						},
					),
						2,
					)

					// required for network group assignment create
					util.MockPostNetworkGroupAssignment(
						util.GenerateNetworkGroupAssignmentRequest(
							"network_group_assignment_uid",
							"",
						),
						util.GenerateNetworkGroupAssignmentResponse(
							"network_group_assignment_uid",
							"",
						),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						name = "name"
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
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						"network_uid",
					),
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						"group_uid",
					),
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"id",
						"network_group_assignment_uid",
					),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
				},
				ResourceName:      "uxi_network_group_assignment.my_network_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid_2", "_2"),
							},
						),
						2,
					)
					util.MockGetWiredNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						2,
					)
					util.MockGetGroup(
						"group_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
							},
						),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						3,
					)

					// required for creating another group
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid_2", "_2", "_2"),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
						),
						1,
					)

					// required for network group assignment create
					util.MockDeleteNetworkGroupAssignment("network_group_assignment_uid", 1)
					util.MockPostNetworkGroupAssignment(
						util.GenerateNetworkGroupAssignmentRequest(
							"network_group_assignment_uid_2",
							"_2",
						),
						util.GenerateNetworkGroupAssignmentResponse(
							"network_group_assignment_uid_2",
							"_2",
						),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid_2",
									"_2",
								),
							},
						),
						2,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						name = "name"
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
						name = "name_2"
					}

					import {
						to = uxi_wired_network.my_network_2
						id = "network_uid_2"
					}

					// the assignment update, updated from network/group to network_2/group_2
					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = uxi_wired_network.my_network_2.id
						group_id 		 = uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						"network_uid_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						"group_uid_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"id",
						"network_group_assignment_uid_2",
					),
				),
			},
			// Delete network-group-assignments and remove networks from state
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid_2", "_2"),
							},
						),
						1,
					)
					util.MockGetWiredNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						1,
					)
					util.MockGetGroup("group_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						},
					),
						2,
					)
					util.MockGetGroup("group_uid_2", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
						},
					),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid_2",
									"_2",
								),
							},
						),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)

					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteGroup("group_uid_2", 1)
					util.MockDeleteNetworkGroupAssignment("network_group_assignment_uid", 1)
					util.MockDeleteNetworkGroupAssignment("network_group_assignment_uid_2", 1)
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
		},
	})

	mockOAuth.Mock.Disable()
}

func TestNetworkGroupAssignmentResourceForWirelessNetwork(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

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
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("network_uid", ""),
							},
						),
						2,
					)
					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						1,
					)

					// required for network group assignment create
					util.MockPostNetworkGroupAssignment(
						util.GenerateNetworkGroupAssignmentRequest(
							"network_group_assignment_uid",
							"",
						),
						util.GenerateNetworkGroupAssignmentResponse(
							"network_group_assignment_uid",
							"",
						),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wireless_network" "my_network" {
						name = "name"
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
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						"network_uid",
					),
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						"group_uid",
					),
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"id",
						"network_group_assignment_uid",
					),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
				},
				ResourceName:      "uxi_network_group_assignment.my_network_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"network_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("network_uid_2", "_2"),
							},
						),
						2,
					)
					util.MockGetWirelessNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("network_uid", ""),
							},
						),
						2,
					)

					// required for creating another group
					util.MockDeleteNetworkGroupAssignment("network_group_assignment_uid", 1)
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid_2", "_2", "_2"),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
						),
						1,
					)
					util.MockGetGroup(
						"group_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
							},
						),
						1,
					)

					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						3,
					)

					// required for network group assignment create
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid_2",
									"_2",
								),
							},
						),
						2,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						2,
					)
					util.MockPostNetworkGroupAssignment(
						util.GenerateNetworkGroupAssignmentRequest(
							"network_group_assignment_uid_2",
							"_2",
						),
						util.GenerateNetworkGroupAssignmentResponse(
							"network_group_assignment_uid_2",
							"_2",
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wireless_network" "my_network" {
						name = "name"
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
						name = "name_2"
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
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						"network_uid_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"group_id",
						"group_uid_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"id",
						"network_group_assignment_uid_2",
					),
				),
			},
			// Delete network-group-assignments and remove networks from state
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"network_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("network_uid_2", "_2"),
							},
						),
						1,
					)
					util.MockGetWirelessNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("network_uid", ""),
							},
						),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						2,
					)
					util.MockGetGroup(
						"group_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
							},
						),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid_2",
									"_2",
								),
							},
						),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)

					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteGroup("group_uid_2", 1)
					util.MockDeleteNetworkGroupAssignment("network_group_assignment_uid", 1)
					util.MockDeleteNetworkGroupAssignment("network_group_assignment_uid_2", 1)
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
		},
	})

	mockOAuth.Mock.Disable()
}

func TestNetworkGroupAssignmentResource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mock429 *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a network group assignment
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						),
						1,
					)
					util.MockGetGroup("group_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						},
					),
						2,
					)

					// required for network group assignment create
					mock429 = gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/network-group-assignments").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPostNetworkGroupAssignment(
						util.GenerateNetworkGroupAssignmentRequest(
							"network_group_assignment_uid",
							"",
						),
						util.GenerateNetworkGroupAssignmentResponse(
							"network_group_assignment_uid",
							"",
						),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						name = "name"
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
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						"network_uid",
					),
					func(s *terraform.State) error {
						assert.Equal(t, mock429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Delete network-group assignment and remove networks from state
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						1,
					)
					util.MockGetGroup("group_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						},
					),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)

					util.MockDeleteGroup("group_uid", 1)
					mock429 = gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/network-group-assignments/network_group_assignment_uid").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockDeleteNetworkGroupAssignment("network_group_assignment_uid", 1)
				},
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_wired_network.my_network

						lifecycle {
							destroy = false
						}
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) error {
						assert.Equal(t, mock429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestNetworkGroupAssignmentResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a network group assignment - errors
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						),
						1,
					)
					util.MockGetGroup("group_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						},
					),
						2,
					)

					// network group assignment create
					gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/network-group-assignments").
						Reply(400).
						JSON(map[string]interface{}{
							"httpStatusCode": 400,
							"errorCode":      "HPE_GL_ERROR_BAD_REQUEST",
							"message":        "Validation error - bad request",
							"debugId":        "12312-123123-123123-1231212",
						})
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						name = "name"
					}

					import {
						to = uxi_wired_network.my_network
						id = "network_uid"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wired_network.my_network.id
						group_id   = uxi_group.my_group.id
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// read not found error
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						),
						1,
					)
					util.MockGetGroup("group_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						},
					),
						2,
					)

					// network group assignment read
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						name = "name"
					}

					import {
						to = uxi_wired_network.my_network
						id = "network_uid"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wired_network.my_network.id
						group_id   = uxi_group.my_group.id
					}

					import {
						to = uxi_network_group_assignment.my_network_group_assignment
						id = "network_group_assignment_uid"
					}
				`,
				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Read 5xx error
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						),
						1,
					)
					util.MockGetGroup("group_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						},
					),
						2,
					)

					// network group assignment read
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/network-group-assignments").
						Reply(500).
						JSON(map[string]interface{}{
							"httpStatusCode": 500,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						name = "name"
					}

					import {
						to = uxi_wired_network.my_network
						id = "network_uid"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wired_network.my_network.id
						group_id   = uxi_group.my_group.id
					}

					import {
						to = uxi_network_group_assignment.my_network_group_assignment
						id = "network_group_assignment_uid"
					}
				`,

				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Actually creating a network group assignment - for next step
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						),
						1,
					)
					util.MockGetGroup("group_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						},
					),
						2,
					)

					// required for network group assignment create
					util.MockPostNetworkGroupAssignment(
						util.GenerateNetworkGroupAssignmentRequest(
							"network_group_assignment_uid",
							"",
						),
						util.GenerateNetworkGroupAssignmentResponse(
							"network_group_assignment_uid",
							"",
						),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						name = "name"
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
					resource.TestCheckResourceAttr(
						"uxi_network_group_assignment.my_network_group_assignment",
						"network_id",
						"network_uid",
					),
				),
			},
			// Delete network-group assignment and remove networks from state - errors
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						1,
					)
					util.MockGetGroup("group_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						},
					),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)

					// network group assignment create
					gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/network-group-assignments").
						Reply(403).
						JSON(map[string]interface{}{
							"httpStatusCode": 403,
							"errorCode":      "HPE_GL_ERROR_FORBIDDEN",
							"message":        "Forbidden - user has insufficient permissions to complete the request",
							"debugId":        "12312-123123-123123-1231212",
						})

				},
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_wired_network.my_network

						lifecycle {
							destroy = false
						}
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Forbidden - user has insufficient permissions to complete the request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Actually delete network-group assignment and remove networks from state
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"network_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWiredNetworkResponse("network_uid", ""),
							},
						),
						1,
					)
					util.MockGetGroup("group_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						},
					),
						1,
					)
					util.MockGetNetworkGroupAssignment(
						"network_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNetworkGroupAssignmentResponse(
									"network_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteNetworkGroupAssignment("network_group_assignment_uid", 1)
				},
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

	mockOAuth.Mock.Disable()
}
