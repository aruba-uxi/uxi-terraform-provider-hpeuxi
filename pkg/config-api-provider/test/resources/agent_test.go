package test

import (
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"regexp"
	"testing"
)

func TestAgentResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating an agent is not allowed
			{
				Config: provider.ProviderConfig + `
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
						return util.GenerateAgentResponseModel(uid, "")
					}
				},
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name = "name"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "uid"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", "name"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", "notes"),
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
						return util.GenerateAgentResponseModel(uid, "_2")
					}
				},
				Config: provider.ProviderConfig + `
				resource "uxi_agent" "my_agent" {
					name = "name_2"
					notes = "notes_2"
					pcap_mode = "light_2"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", "name_2"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", "notes_2"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "pcap_mode", "light_2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})

	mockOAuth.Mock.Disable()
}
