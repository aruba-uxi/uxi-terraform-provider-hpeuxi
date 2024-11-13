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

func TestAgentGroupAssignmentResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a agent group assignment
			{
				PreConfig: func() {
					// required for agent import
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
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

					// required for agent group assignment create
					util.MockPostAgentGroupAssignment(
						util.GenerateAgentGroupAssignmentRequest("agent_group_assignment_uid", ""),
						util.GenerateAgentGroupAssignmentResponse("agent_group_assignment_uid", ""),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
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

					resource "uxi_agent" "my_agent" {
						name 			= "name"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_uid"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id       = uxi_agent.my_agent.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						"agent_uid",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						"group_uid",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"id",
						"agent_group_assignment_uid",
					),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
				},
				ResourceName:      "uxi_agent_group_assignment.my_agent_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update testing
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid_2", "_2"),
							},
						),
						2,
					)
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
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
						2,
					)
					util.MockDeleteAgentGroupAssignment("agent_group_assignment_uid", 1)

					// required for creating another group
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid_2", "_2", "_2"),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
						),
						1,
					)

					// required for agent group assignment create
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid_2",
									"_2",
								),
							},
						),
						2,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
					util.MockPostAgentGroupAssignment(
						util.GenerateAgentGroupAssignmentRequest(
							"agent_group_assignment_uid_2",
							"_2",
						),
						util.GenerateAgentGroupAssignmentResponse(
							"agent_group_assignment_uid_2",
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

					resource "uxi_agent" "my_agent" {
						name 			= "name"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_uid"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "name_2"
						parent_group_id = "parent_uid_2"
					}

					resource "uxi_agent" "my_agent_2" {
						name 			= "name_2"
						notes 			= "notes_2"
						pcap_mode 		= "light_2"
					}

					import {
						to = uxi_agent.my_agent_2
						id = "agent_uid_2"
					}

					// the assignment update, updated from agent/group to agent_2/group_2
					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id       = uxi_agent.my_agent_2.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						"agent_uid_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						"group_uid_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"id",
						"agent_group_assignment_uid_2",
					),
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
							},
						),
						2,
					)
					util.MockGetAgent(
						"agent_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid_2", "_2"),
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
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
									"",
								),
							},
						),
						2,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid_2",
									"_2",
								),
							},
						),
						1,
					)
					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteGroup("group_uid_2", 1)
					util.MockDeleteAgent("agent_uid", 1)
					util.MockDeleteAgent("agent_uid_2", 1)
					util.MockDeleteAgentGroupAssignment("agent_group_assignment_uid_2", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestAgentGroupAssignmentResource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mock429 *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a agent group assignment
			{
				PreConfig: func() {
					// required for agent import
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
							},
						),
						2,
					)

					// required for agent group assignment create
					mock429 = gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/agent-group-assignments").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
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

					// required for agent group assignment create
					util.MockPostAgentGroupAssignment(
						util.GenerateAgentGroupAssignmentRequest("agent_group_assignment_uid", ""),
						util.GenerateAgentGroupAssignmentResponse("agent_group_assignment_uid", ""),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
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

					resource "uxi_agent" "my_agent" {
						name 			= "name"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_uid"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id       = uxi_agent.my_agent.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"id",
						"agent_group_assignment_uid",
					),
					func(s *terraform.State) error {
						assert.Equal(t, mock429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					mock429 = gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/agent-group-assignments").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
				},
				ResourceName:      "uxi_agent_group_assignment.my_agent_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
				Check: func(s *terraform.State) error {
					assert.Equal(t, mock429.Mock.Request().Counter, 0)
					return nil
				},
			},
			// Delete testing
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
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
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteAgent("agent_uid", 1)
					mock429 = gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/agent-group-assignments/agent_group_assignment_uid").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockDeleteAgentGroupAssignment("agent_group_assignment_uid", 1)
				},
				Config: provider.ProviderConfig,
				Check: func(s *terraform.State) error {
					assert.Equal(t, mock429.Mock.Request().Counter, 0)
					return nil
				},
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestAgentGroupAssignmentResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating an agent group assignment - errors
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
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

					// agent group assignment create
					gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/agent-group-assignments").
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

					resource "uxi_agent" "my_agent" {
						name      = "name"
						notes 	  = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_uid"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = uxi_agent.my_agent.id
						group_id   = uxi_group.my_group.id
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// read not found error
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
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

					// agent group assignment read
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_agent" "my_agent" {
						name 	  = "name"
						notes 	  = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_uid"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = uxi_agent.my_agent.id
						group_id   = uxi_group.my_group.id
					}

					import {
						to = uxi_agent_group_assignment.my_agent_group_assignment
						id = "agent_group_assignment_uid"
					}
				`,
				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Read 5xx error
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
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

					// agent group assignment read
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/agent-group-assignments").
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

					resource "uxi_agent" "my_agent" {
						name 	  = "name"
						notes 	  = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_uid"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = uxi_agent.my_agent.id
						group_id   = uxi_group.my_group.id
					}

					import {
						to = uxi_agent_group_assignment.my_agent_group_assignment
						id = "agent_group_assignment_uid"
					}
				`,

				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Actually creating an agent group assignment - for next step
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
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

					// required for agent group assignment create
					util.MockPostAgentGroupAssignment(
						util.GenerateAgentGroupAssignmentRequest(
							"agent_group_assignment_uid",
							"",
						),
						util.GenerateAgentGroupAssignmentResponse(
							"agent_group_assignment_uid",
							"",
						),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
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

					resource "uxi_agent" "my_agent" {
						name 	  = "name"
						notes 	  = "notes"
						pcap_mode = "light"

					}

					import {
						to = uxi_agent.my_agent
						id = "agent_uid"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = uxi_agent.my_agent.id
						group_id = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						"agent_uid",
					),
				),
			},
			// Delete agent-group assignment and remove agents from state - errors
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
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
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)

					// agent group assignment create
					gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/agent-group-assignments").
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
						from = uxi_agent.my_agent

						lifecycle {
							destroy = false
						}
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Forbidden - user has insufficient permissions to complete the request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Actually delete agent-group assignment and remove agents from state
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_uid", ""),
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
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_uid",
									"",
								),
							},
						),
						1,
					)
					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteAgentGroupAssignment("agent_group_assignment_uid", 1)
				},
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_agent.my_agent

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
