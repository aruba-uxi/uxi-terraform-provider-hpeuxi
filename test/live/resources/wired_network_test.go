package resource_test

import (
	"github.com/aruba-uxi/terraform-provider-configuration/test/live/config"
	"github.com/aruba-uxi/terraform-provider-configuration/test/live/provider"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWiredNetworkResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a wired_network is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "uxi_wired_network" "my_wired_network" {
						name = "name"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)creating a wired_network is not supported; wired_networks can only be\s*imported`,
				),
			},
			// Importing a wired_network
			{
				Config: provider.ProviderConfig + `
					resource "uxi_wired_network" "wired_network_0" {
						name = "` + config.WiredNetworkName + `"
					}

					import {
						to = uxi_wired_network.wired_network_0
						id = "` + config.WiredNetworkUid + `"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_wired_network.wired_network_0",
						"name",
						config.WiredNetworkName,
					),
					resource.TestCheckResourceAttr(
						"uxi_wired_network.wired_network_0",
						"id",
						config.WiredNetworkUid,
					),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_wired_network.wired_network_0",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a wired_network is not allowed
			{
				Config: provider.ProviderConfig + `
				resource "uxi_wired_network" "wired_network_0" {
					name = "` + config.WiredNetworkUid + `-updated-name"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)updating a wired_network is not supported; wired_networks can only be updated\s*through the dashboard`,
				),
			},
			// Deleting a wired_network is not allowed
			{
				Config: provider.ProviderConfig,
				ExpectError: regexp.MustCompile(
					`(?s)deleting a wired_network is not supported; wired_networks can only removed\s*from state`,
				),
			},
			// Remove wired_network from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_wired_network.wired_network_0

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}
