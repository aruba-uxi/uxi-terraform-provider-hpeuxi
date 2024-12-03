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

func TestAgentGroupAssignmentResource(t *testing.T) {
	const (
		groupName  = "tf_provider_acceptance_test_agent_group_assignment_resource"
		group2Name = "tf_provider_acceptance_test_agent_group_assignment_resource_two"
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

					data "hpeuxi_agent" "my_agent" {
						filter = {
							id = "` + config.AgentID + `"
						}
					}

					resource "hpeuxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = data.hpeuxi_agent.my_agent.id
						group_id = hpeuxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check properties are what we configured
					resource.TestCheckResourceAttr(
						"hpeuxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						config.AgentID,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						func(value string) error {
							assert.Equal(t, util.GetGroupByName(groupName).Id, value)

							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "hpeuxi_agent_group_assignment.my_agent_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIDBeforeRecreate = rs.Primary.ID

						return util.CheckStateAgainstAgentGroupAssignment(
							t,
							"hpeuxi_agent_group_assignment.my_agent_group_assignment",
							*util.GetAgentGroupAssignment(resourceIDBeforeRecreate),
						)(s)
					},
				),
			},
			// ImportState
			{
				ResourceName:      "hpeuxi_agent_group_assignment.my_agent_group_assignment",
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

					data "hpeuxi_agent" "my_agent" {
						filter = {
							id = "` + config.AgentID + `"
						}
					}

					// the new resources we wanna update the assignment to
					resource "hpeuxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from agent/group to agent/group_2
					resource "hpeuxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = data.hpeuxi_agent.my_agent.id
						group_id = hpeuxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check properties are what we configured
					resource.TestCheckResourceAttr(
						"hpeuxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						config.AgentID,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						func(value string) error {
							assert.Equal(t, util.GetGroupByName(group2Name).Id, value)

							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "hpeuxi_agent_group_assignment.my_agent_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIDAfterRecreate = rs.Primary.ID

						return util.CheckStateAgainstAgentGroupAssignment(
							t,
							"hpeuxi_agent_group_assignment.my_agent_group_assignment",
							*util.GetAgentGroupAssignment(resourceIDAfterRecreate),
						)(s)
					},
				),
			},
			// Delete
			{
				Config: provider.ProviderConfig,
			},
		},
		CheckDestroy: func(s *terraform.State) error {
			assert.Nil(t, util.GetGroupByName(groupName))
			assert.Nil(t, util.GetGroupByName(group2Name))
			assert.Nil(t, util.GetAgentGroupAssignment(resourceIDBeforeRecreate))
			assert.Nil(t, util.GetAgentGroupAssignment(resourceIDAfterRecreate))

			return nil
		},
	})
}
