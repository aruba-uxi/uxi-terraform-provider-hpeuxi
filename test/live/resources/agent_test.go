/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func TestAgentResource(t *testing.T) {
	// we provision an agent here so that we have something to delete later on
	agentID, err := util.ProvisionAgent{
		CustomerID:        config.CustomerID,
		ProvisionToken:    os.Getenv("HPEUXI_PROVISION_TOKEN"),
		DeviceSerial:      config.AgentCreateSerial,
		DeviceGatewayHost: config.DeviceGatewayHost,
	}.Provision()
	if err != nil {
		panic(err)
	}

	agent := util.GetAgent(agentID)
	updatedAgent := agent
	updatedNotes := "notes"
	updatedPcapMode := config_api_client.AGENTPCAPMODE_OFF
	updatedAgent.Name = "tf_provider_acceptance_test_agent_resource_updated_name"
	updatedAgent.Notes = *config_api_client.NewNullableString(&updatedNotes)
	updatedAgent.PcapMode = *config_api_client.NewNullableAgentPcapMode(&updatedPcapMode)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating an agent is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_agent" "my_agent" {
						name = "` + agent.Name + `"
					}`,

				ExpectError: regexp.MustCompile(
					`creating an agent is not supported; agents can only be imported`,
				),
			},
			// Importing an agent
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_agent" "my_agent" {
						name = "` + agent.Name + `"
					}

					import {
						to = hpeuxi_agent.my_agent
						id = "` + agentID + `"
					}`,
				Check: shared.CheckStateAgainstAgent(t, "hpeuxi_agent.my_agent", agent),
			},
			// ImportState
			{
				ResourceName:      "hpeuxi_agent.my_agent",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_agent" "my_agent" {
						name = "tf_provider_acceptance_test_agent_resource_updated_name"
						notes = "notes"
						pcap_mode = "off"
					}`,
				Check: shared.CheckStateAgainstAgent(t, "hpeuxi_agent.my_agent", updatedAgent),
			},
			// Delete
			{
				Config: provider.ProviderConfig,
			},
		},
		CheckDestroy: func(s *terraform.State) error {
			assert.Nil(t, util.GetAgent(agentID))

			return nil
		},
	})
}
