/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWiredNetworkResource(t *testing.T) {
	wiredNetwork := util.GetWiredNetwork(config.WiredNetworkID)

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
					resource "hpeuxi_wired_network" "my_wired_network" {
						name = "name"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)creating a wired_network is not supported; wired_networks can only be\s*imported`,
				),
			},
			// Importing a wired_network
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_wired_network" "wired_network_0" {
						name = "` + config.WiredNetworkName + `"
					}

					import {
						to = hpeuxi_wired_network.wired_network_0
						id = "` + config.WiredNetworkID + `"
					}`,

				Check: shared.CheckStateAgainstWiredNetwork(
					t,
					"hpeuxi_wired_network.wired_network_0",
					wiredNetwork,
				),
			},
			// ImportState testing
			{
				ResourceName:      "hpeuxi_wired_network.wired_network_0",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a wired_network is not allowed
			{
				Config: provider.ProviderConfig + `
				resource "hpeuxi_wired_network" "wired_network_0" {
					name = "` + config.WiredNetworkID + `-updated-name"
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
						from = hpeuxi_wired_network.wired_network_0

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}
