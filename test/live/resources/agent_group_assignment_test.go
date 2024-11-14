package resource_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/nbio/st"
)

func TestAgentGroupAssignmentResource(t *testing.T) {
	const groupName = "tf_provider_acceptance_test_agent_group_assignment_resource"
	const group2Name = "tf_provider_acceptance_test_agent_group_assignment_resource_two"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a agent group assignment
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					data "uxi_agent" "my_agent" {
						filter = {
							agent_id = "` + config.AgentPermanentId + `"
						}
					}

					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = data.uxi_agent.my_agent.id
						group_id = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check properties are what we configured
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						config.AgentPermanentId,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						func(value string) error {
							st.Assert(t, util.GetGroupByName(groupName).Id, value)
							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "uxi_agent_group_assignment.my_agent_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						return util.CheckStateAgainstAgentGroupAssignment(
							t,
							"uxi_agent_group_assignment.my_agent_group_assignment",
							util.GetAgentGroupAssignment(rs.Primary.ID),
						)(s)
					},
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_agent_group_assignment.my_agent_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update testing
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					data "uxi_agent" "my_agent" {
						filter = {
							agent_id = "` + config.AgentPermanentId + `"
						}
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "` + group2Name + `"
					}

					// the assignment update, updated from agent/group to agent/group_2
					resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
						agent_id = data.uxi_agent.my_agent.id
						group_id = uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check properties are what we configured
					resource.TestCheckResourceAttr(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						config.AgentPermanentId,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						func(value string) error {
							st.Assert(t, util.GetGroupByName(group2Name).Id, value)
							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "uxi_agent_group_assignment.my_agent_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						return util.CheckStateAgainstAgentGroupAssignment(
							t,
							"uxi_agent_group_assignment.my_agent_group_assignment",
							util.GetAgentGroupAssignment(rs.Primary.ID),
						)(s)
					},
				),
			},
			// Delete testing happens automatically
		},
	})
}
