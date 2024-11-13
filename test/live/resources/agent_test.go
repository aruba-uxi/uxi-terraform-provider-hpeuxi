package resource_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAgentResource(t *testing.T) {
	const (
		agentName        = "tf_provider_acceptance_test_agent_resource"
		agentNameUpdated = agentName + "_updated"
	)

	// we provision an agent here so that we have something to delete later on
	agentUid, err := util.ProvisionAgent{
		CustomerUid:       config.CustomerUid,
		ProvisionToken:    os.Getenv("UXI_PROVISION_TOKEN"),
		DeviceSerial:      config.AgentCreateSerial,
		DeviceGatewayHost: config.DeviceGatewayHost,
	}.Provision()
	if err != nil {
		panic(err)
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating an agent is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name = "` + agentName + `"
					}`,

				ExpectError: regexp.MustCompile(
					`creating an agent is not supported; agents can only be imported`,
				),
			},
			// Importing an agent
			{
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name  	  = "` + agentName + `"
						notes 	  = ""
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "` + agentUid + `"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "id", agentUid),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", agentName),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", ""),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "pcap_mode", "light"),
				),
			},
			// ImportState
			{
				ResourceName:      "uxi_agent.my_agent",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: provider.ProviderConfig + `
				resource "uxi_agent" "my_agent" {
					name = "` + agentNameUpdated + `"
					notes = "notes"
					pcap_mode = "off"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "id", agentUid),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", agentNameUpdated),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", "notes"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "pcap_mode", "off"),
				),
			},
			// Delete
			{
				Config:  provider.ProviderConfig,
				Destroy: true,
			},
		},
		CheckDestroy: func(s *terraform.State) error {
			assert.Equal(t, util.GetAgent(agentUid), nil)
			return nil
		},
	})
}
