package test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAgentResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating an agent is not allowed
			{
				Config: providerConfig + `
					resource "uxi_agent" "my_agent" {
						name = "name"
						notes = "note"
						pcap_mode = "light"
					}`,

				ExpectError: regexp.MustCompile(`creating an agent is not supported; agents can only be imported`),
			},
			// Importing an agent
			{
				PreConfig: func() {
					resources.GetAgent = func() resources.AgentResponseModel {
						return resources.AgentResponseModel{
							UID:                "uid",
							Serial:             "serial",
							Name:               "imported_name",
							ModelNumber:        "model_number",
							WifiMacAddress:     "wifi_mac_address",
							EthernetMacAddress: "ethernet_mac_address",
							Notes:              "imported_notes",
							PCapMode:           "light",
						}
					}
				},
				Config: providerConfig + `
					resource "uxi_agent" "my_agent" {
						name = "imported_name"
						notes = "imported_notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "test_uid"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify first order item updated
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", "imported_name"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", "imported_notes"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "pcap_mode", "light"),
					// Verify first coffee item has Computed attributes updated.
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "id", "test_uid"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_agent.my_agent",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					resources.GetAgent = func() resources.AgentResponseModel {
						return resources.AgentResponseModel{
							UID:                "uid",
							Serial:             "serial",
							Name:               "updated_name",
							ModelNumber:        "model_number",
							WifiMacAddress:     "wifi_mac_address",
							EthernetMacAddress: "ethernet_mac_address",
							Notes:              "updated_notes",
							PCapMode:           "not_light",
						}
					}
				},
				Config: providerConfig + `
				resource "uxi_agent" "my_agent" {
					name = "updated_name"
					notes = "updated_notes"
					pcap_mode = "not_light"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", "updated_name"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", "updated_notes"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "pcap_mode", "not_light"),
				),
			},
			// Deleting an agent is not allowed
			{
				Config:      providerConfig + ``,
				ExpectError: regexp.MustCompile(`deleting an agent is not supported; agents can only removed from state`),
			},
			// Remove agent from state
			{
				Config: providerConfig + `
					removed {
						from = uxi_agent.my_agent

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}
