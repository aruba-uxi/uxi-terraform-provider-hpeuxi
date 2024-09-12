package test

import (
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWirelessNetworkResource(t *testing.T) {
	defer gock.Off()
	MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a wireless_network is not allowed
			{
				Config: providerConfig + `
					resource "uxi_wireless_network" "my_wireless_network" {
						alias = "alias"
					}`,

				ExpectError: regexp.MustCompile(`(?s)creating a wireless_network is not supported; wireless_networks can only be\s*imported`),
			},
			// Importing a wireless_network
			{
				PreConfig: func() {
					MockGetWirelessNetwork(
						"uid",
						GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{GenerateWirelessNetworkResponse("uid", "")}),
						2,
					)
				},
				Config: providerConfig + `
					resource "uxi_wireless_network" "my_wireless_network" {
						alias = "alias"
					}

					import {
						to = uxi_wireless_network.my_wireless_network
						id = "uid"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_wireless_network.my_wireless_network", "alias", "alias"),
					resource.TestCheckResourceAttr("uxi_wireless_network.my_wireless_network", "id", "uid"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					MockGetWirelessNetwork(
						"uid",
						GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{GenerateWirelessNetworkResponse("uid", "")}),
						1,
					)
				},
				ResourceName:      "uxi_wireless_network.my_wireless_network",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a wireless_network is not allowed
			{
				PreConfig: func() {
					MockGetWirelessNetwork(
						"uid",
						GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{GenerateWirelessNetworkResponse("uid", "")}),
						1,
					)
				},
				Config: providerConfig + `
				resource "uxi_wireless_network" "my_wireless_network" {
					alias = "updated_alias"
				}`,
				ExpectError: regexp.MustCompile(`(?s)updating a wireless_network is not supported; wireless_networks can only be\s*updated through the dashboard`),
			},
			// Deleting a wireless_network is not allowed
			{
				PreConfig: func() {
					MockGetWirelessNetwork(
						"uid",
						GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{GenerateWirelessNetworkResponse("uid", "")}),
						1,
					)
				},
				Config:      providerConfig + ``,
				ExpectError: regexp.MustCompile(`(?s)deleting a wireless_network is not supported; wireless_networks can only\s*removed from state`),
			},
			// Remove wireless_network from state
			{
				PreConfig: func() {
					MockGetWirelessNetwork(
						"uid",
						GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{GenerateWirelessNetworkResponse("uid", "")}),
						1,
					)
				},
				Config: providerConfig + `
					removed {
						from = uxi_wireless_network.my_wireless_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}
