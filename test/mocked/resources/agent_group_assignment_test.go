package resource_test

import (
	"net/http"
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
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_id", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_id", "", ""),
							},
						),
						1,
					)

					// required for agent group assignment create
					util.MockPostAgentGroupAssignment(
						util.GenerateAgentGroupAssignmentRequest("agent_group_assignment_id", ""),
						util.GenerateAgentGroupAssignmentResponse("agent_group_assignment_id", ""),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id",
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
						parent_group_id = "parent_id"
					}

					resource "uxi_agent" "my_agent" {
						name 			= "name"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_id"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id       = uxi_agent.my_agent.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						"agent_id",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						"group_id",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"id",
						"agent_group_assignment_id",
					),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id",
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
						"agent_id_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id_2", "_2"),
							},
						),
						2,
					)
					util.MockGetAgent(
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						2,
					)
					util.MockGetGroup(
						"group_id_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_id_2", "_2", "_2"),
							},
						),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_id", "", ""),
							},
						),
						2,
					)
					util.MockDeleteAgentGroupAssignment("agent_group_assignment_id", 1)

					// required for creating another group
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_id_2", "_2", "_2"),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_id_2", "_2", "_2"),
						),
						1,
					)

					// required for agent group assignment create
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id_2",
									"_2",
								),
							},
						),
						2,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id",
									"",
								),
							},
						),
						1,
					)
					util.MockPostAgentGroupAssignment(
						util.GenerateAgentGroupAssignmentRequest(
							"agent_group_assignment_id_2",
							"_2",
						),
						util.GenerateAgentGroupAssignmentResponse(
							"agent_group_assignment_id_2",
							"_2",
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_agent" "my_agent" {
						name 			= "name"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_id"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "name_2"
						parent_group_id = "parent_id_2"
					}

					resource "uxi_agent" "my_agent_2" {
						name 			= "name_2"
						notes 			= "notes_2"
						pcap_mode 		= "light_2"
					}

					import {
						to = uxi_agent.my_agent_2
						id = "agent_id_2"
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
						"agent_id_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						"group_id_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"id",
						"agent_group_assignment_id_2",
					),
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						2,
					)
					util.MockGetAgent(
						"agent_id_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id_2", "_2"),
							},
						),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_id", "", ""),
							},
						),
						2,
					)
					util.MockGetGroup(
						"group_id_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_id_2", "_2", "_2"),
							},
						),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id",
									"",
								),
							},
						),
						2,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id_2",
									"_2",
								),
							},
						),
						1,
					)
					util.MockDeleteGroup("group_id", 1)
					util.MockDeleteGroup("group_id_2", 1)
					util.MockDeleteAgent("agent_id", 1)
					util.MockDeleteAgent("agent_id_2", 1)
					util.MockDeleteAgentGroupAssignment("agent_group_assignment_id_2", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestAgentGroupAssignmentResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a agent group assignment
			{
				PreConfig: func() {
					// required for agent import
					util.MockGetAgent(
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						2,
					)

					// required for agent group assignment create
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/agent-group-assignments").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_id", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_id", "", ""),
							},
						),
						1,
					)

					// required for agent group assignment create
					util.MockPostAgentGroupAssignment(
						util.GenerateAgentGroupAssignmentRequest("agent_group_assignment_id", ""),
						util.GenerateAgentGroupAssignmentResponse("agent_group_assignment_id", ""),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id",
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
						parent_group_id = "parent_id"
					}

					resource "uxi_agent" "my_agent" {
						name 			= "name"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_id"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id       = uxi_agent.my_agent.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"id",
						"agent_group_assignment_id",
					),
					func(s *terraform.State) error {
						assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/agent-group-assignments").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id",
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
					assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
					return nil
				},
			},
			// Delete testing
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_id", "", ""),
							},
						),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id",
									"",
								),
							},
						),
						1,
					)
					util.MockDeleteGroup("group_id", 1)
					util.MockDeleteAgent("agent_id", 1)
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/agent-group-assignments/agent_group_assignment_id").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockDeleteAgentGroupAssignment("agent_group_assignment_id", 1)
				},
				Config: provider.ProviderConfig,
				Check: func(s *terraform.State) error {
					assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
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
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_id", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						),
						1,
					)
					util.MockGetGroup("group_id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						},
					),
						2,
					)

					// agent group assignment create
					gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/agent-group-assignments").
						Reply(http.StatusBadRequest).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusBadRequest,
							"errorCode":      "HPE_GL_ERROR_BAD_REQUEST",
							"message":        "Validation error - bad request",
							"debugId":        "12312-123123-123123-1231212",
						})
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_agent" "my_agent" {
						name      = "name"
						notes 	  = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_id"
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
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_id", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						),
						1,
					)
					util.MockGetGroup("group_id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						},
					),
						2,
					)

					// agent group assignment read
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_agent" "my_agent" {
						name 	  = "name"
						notes 	  = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_id"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = uxi_agent.my_agent.id
						group_id   = uxi_group.my_group.id
					}

					import {
						to = uxi_agent_group_assignment.my_agent_group_assignment
						id = "agent_group_assignment_id"
					}
				`,
				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Read 5xx error
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_id", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						),
						1,
					)
					util.MockGetGroup("group_id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						},
					),
						2,
					)

					// agent group assignment read
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/agent-group-assignments").
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_agent" "my_agent" {
						name 	  = "name"
						notes 	  = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "agent_id"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = uxi_agent.my_agent.id
						group_id   = uxi_group.my_group.id
					}

					import {
						to = uxi_agent_group_assignment.my_agent_group_assignment
						id = "agent_group_assignment_id"
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
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_id", "", ""),
						util.StructToMap(
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						),
						1,
					)
					util.MockGetGroup("group_id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						},
					),
						2,
					)

					// required for agent group assignment create
					util.MockPostAgentGroupAssignment(
						util.GenerateAgentGroupAssignmentRequest(
							"agent_group_assignment_id",
							"",
						),
						util.GenerateAgentGroupAssignmentResponse(
							"agent_group_assignment_id",
							"",
						),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id",
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
						parent_group_id = "parent_id"
					}

					resource "uxi_agent" "my_agent" {
						name 	  = "name"
						notes 	  = "notes"
						pcap_mode = "light"

					}

					import {
						to = uxi_agent.my_agent
						id = "agent_id"
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = uxi_agent.my_agent.id
						group_id = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						"agent_id",
					),
				),
			},
			// Delete agent-group assignment and remove agents from state - errors
			{
				PreConfig: func() {
					util.MockGetAgent(
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						1,
					)
					util.MockGetGroup("group_id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						},
					),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id",
									"",
								),
							},
						),
						1,
					)

					// agent group assignment create
					gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/agent-group-assignments").
						Reply(http.StatusForbidden).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusForbidden,
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
						"agent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentResponseModel("agent_id", ""),
							},
						),
						1,
					)
					util.MockGetGroup("group_id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("group_id", "", ""),
						},
					),
						1,
					)
					util.MockGetAgentGroupAssignment(
						"agent_group_assignment_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateAgentGroupAssignmentResponse(
									"agent_group_assignment_id",
									"",
								),
							},
						),
						1,
					)
					util.MockDeleteGroup("group_id", 1)
					util.MockDeleteAgentGroupAssignment("agent_group_assignment_id", 1)
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
