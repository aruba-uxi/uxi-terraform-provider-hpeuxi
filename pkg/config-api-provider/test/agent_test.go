package test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAgentResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
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
					resources.GetAgent = func(uid string) resources.AgentResponseModel {
						return resources.AgentResponseModel{
							UID:                uid,
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
						id = "uid"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", "imported_name"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", "imported_notes"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "pcap_mode", "light"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "id", "uid"),
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
					resources.GetAgent = func(uid string) resources.AgentResponseModel {
						return resources.AgentResponseModel{
							UID:                uid,
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
			// Delete testing automatically occurs in TestCase
		},
	})
}
