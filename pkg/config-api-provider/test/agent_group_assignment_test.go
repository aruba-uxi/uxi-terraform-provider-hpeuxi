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
					// required for agent import
					resources.GetAgent = func(uid string) resources.AgentResponseModel {
						return resources.AgentResponseModel{
							UID:                uid,
							Serial:             "serial",
							Name:               "name",
							ModelNumber:        "model_number",
							WifiMacAddress:     "wifi_mac_address",
							EthernetMacAddress: "ethernet_mac_address",
							Notes:              "notes",
							PCapMode:           "light",
						}
					}

					// required for group create
					groupResponse := resources.GroupResponseModel{
						UID:       "group_uid",
						Name:      "name",
						ParentUid: "parent_uid",
						Path:      "parent_uid.group_uid",
					}
					resources.CreateGroup = func(request resources.GroupCreateRequestModel) resources.GroupResponseModel {
						return groupResponse
					}
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						return groupResponse
					}

					// required for agent group assignment create
					agentGroupAssignmentResponse := resources.AgentGroupAssignmentResponseModel{
						UID:      "agent_group_assignment_uid",
						GroupUID: "group_uid",
						AgentUID: "agent_uid",
					}
					resources.CreateAgentGroupAssignment = func(request resources.AgentGroupAssignmentRequestModel) resources.AgentGroupAssignmentResponseModel {
						return agentGroupAssignmentResponse
					}
					resources.GetAgentGroupAssignment = func(uid string) resources.AgentGroupAssignmentResponseModel {
						return agentGroupAssignmentResponse
					}
				},

				Config: providerConfig + `
					resource "uxi_group" "my_group" {
						name       = "name"
						parent_uid = "parent_uid"
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
					resources.GetAgent = func(uid string) resources.AgentResponseModel {
						if uid == "agent_uid" {
							return resources.AgentResponseModel{
								UID:                "agent_uid",
								Serial:             "serial",
								Name:               "name",
								ModelNumber:        "model_number",
								WifiMacAddress:     "wifi_mac_address",
								EthernetMacAddress: "ethernet_mac_address",
								Notes:              "notes",
								PCapMode:           "light",
							}
						} else {
							return resources.AgentResponseModel{
								UID:                "agent_uid_2",
								Serial:             "serial_2",
								Name:               "name_2",
								ModelNumber:        "model_number_2",
								WifiMacAddress:     "wifi_mac_address_2",
								EthernetMacAddress: "ethernet_mac_address_2",
								Notes:              "notes_2",
								PCapMode:           "light",
							}
						}
					}

					// required for creating another group
					newGroupResponse := resources.GroupResponseModel{
						UID:       "group_uid_2",
						Name:      "name_2",
						ParentUid: "parent_uid_2",
						Path:      "parent_uid_2.group_uid_2",
					}
					resources.CreateGroup = func(request resources.GroupCreateRequestModel) resources.GroupResponseModel {
						return newGroupResponse
					}
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						if uid == "group_uid" {
							return resources.GroupResponseModel{
								UID:       uid,
								Name:      "name",
								ParentUid: "parent_uid",
								Path:      "parent_uid.group_uid",
							}
						} else {
							return newGroupResponse
						}
					}

					// required for agent group assignment create
					agentGroupAssignmentOriginal := resources.AgentGroupAssignmentResponseModel{
						UID:      "agent_group_assignment_uid",
						GroupUID: "group_uid",
						AgentUID: "agent_uid",
					}
					agentGroupAssignmentUpdated := resources.AgentGroupAssignmentResponseModel{
						UID:      "agent_group_assignment_uid_2",
						GroupUID: "group_uid_2",
						AgentUID: "agent_uid_2",
					}

					resources.GetAgentGroupAssignment = func(uid string) resources.AgentGroupAssignmentResponseModel {
						if uid == "agent_group_assignment_uid" {
							return agentGroupAssignmentOriginal
						} else {
							return agentGroupAssignmentUpdated
						}
					}
					resources.CreateAgentGroupAssignment = func(request resources.AgentGroupAssignmentRequestModel) resources.AgentGroupAssignmentResponseModel {
						return agentGroupAssignmentUpdated
					}
				},
				Config: providerConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name       = "name"
						parent_uid = "parent_uid"
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
						name       = "name_2"
						parent_uid = "parent_uid_2"
					}

					resource "uxi_agent" "my_agent_2" {
						name 			= "name_2"
						notes 			= "notes_2"
						pcap_mode 		= "light"
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
			// Remove agents from state
			{
				Config: providerConfig + `
					removed {
						from = uxi_agent.my_agent

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = uxi_agent.my_agent_2

						lifecycle {
							destroy = false
						}
					}`,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
