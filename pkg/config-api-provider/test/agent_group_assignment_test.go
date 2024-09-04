package test

import (
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAgentGroupAssignmentResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a agent group assignment
			{
				PreConfig: func() {
					MockOAuth()
					// required for agent import
					resources.GetAgent = func(uid string) resources.AgentResponseModel {
						return GenerateAgentResponseModel(uid, "")
					}

					// required for group create
					MockPostGroup(StructToMap(GenerateGroupResponseModel("group_uid", "", "")))
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						return GenerateGroupResponseModel("group_uid", "", "")
					}

					// required for agent group assignment create
					agentGroupAssignmentResponse := GenerateAgentGroupAssignmentResponse("agent_group_assignment_uid", "")
					resources.CreateAgentGroupAssignment = func(request resources.AgentGroupAssignmentRequestModel) resources.AgentGroupAssignmentResponseModel {
						return agentGroupAssignmentResponse
					}
					resources.GetAgentGroupAssignment = func(uid string) resources.AgentGroupAssignmentResponseModel {
						return agentGroupAssignmentResponse
					}
				},

				Config: providerConfig + `
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
					resource.TestCheckResourceAttr("uxi_agent_group_assignment.my_agent_group_assignment", "agent_id", "agent_uid"),
					resource.TestCheckResourceAttr("uxi_agent_group_assignment.my_agent_group_assignment", "group_id", "group_uid"),
					resource.TestCheckResourceAttr("uxi_agent_group_assignment.my_agent_group_assignment", "id", "agent_group_assignment_uid"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_agent_group_assignment.my_agent_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					MockOAuth()
					resources.GetAgent = func(uid string) resources.AgentResponseModel {
						if uid == "agent_uid" {
							return GenerateAgentResponseModel(uid, "")
						} else {
							return GenerateAgentResponseModel(uid, "_2")
						}
					}

					// required for creating another group
					MockPostGroup(StructToMap(GenerateGroupResponseModel("group_uid_2", "_2", "_2")))
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						if uid == "group_uid" {
							return GenerateGroupResponseModel(uid, "", "")
						} else {
							return GenerateGroupResponseModel(uid, "_2", "_2")
						}
					}

					// required for agent group assignment create
					resources.GetAgentGroupAssignment = func(uid string) resources.AgentGroupAssignmentResponseModel {
						if uid == "agent_group_assignment_uid" {
							return GenerateAgentGroupAssignmentResponse(uid, "")
						} else {
							return GenerateAgentGroupAssignmentResponse(uid, "_2")
						}
					}
					resources.CreateAgentGroupAssignment = func(request resources.AgentGroupAssignmentRequestModel) resources.AgentGroupAssignmentResponseModel {
						return GenerateAgentGroupAssignmentResponse("agent_group_assignment_uid_2", "_2")
					}
				},
				Config: providerConfig + `
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
					resource.TestCheckResourceAttr("uxi_agent_group_assignment.my_agent_group_assignment", "agent_id", "agent_uid_2"),
					resource.TestCheckResourceAttr("uxi_agent_group_assignment.my_agent_group_assignment", "group_id", "group_uid_2"),
					resource.TestCheckResourceAttr("uxi_agent_group_assignment.my_agent_group_assignment", "id", "agent_group_assignment_uid_2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
