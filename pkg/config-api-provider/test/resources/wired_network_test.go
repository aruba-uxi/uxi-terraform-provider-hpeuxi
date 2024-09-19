package test

import (
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWiredNetworkResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

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
						alias = "alias"
					}`,

				ExpectError: regexp.MustCompile(`(?s)creating a wired_network is not supported; wired_networks can only be\s*imported`),
			},
			// Importing a wired_network
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"uid",
						util.GenerateWiredNetworkPaginatedResponse(
							[]map[string]interface{}{util.GenerateWiredNetworkResponse("uid", "")},
						),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_wired_network" "my_wired_network" {
						alias = "alias"
					}

					import {
						to = uxi_wired_network.my_wired_network
						id = "uid"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_wired_network.my_wired_network", "alias", "alias"),
					resource.TestCheckResourceAttr("uxi_wired_network.my_wired_network", "id", "uid"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"uid",
						util.GenerateWiredNetworkPaginatedResponse(
							[]map[string]interface{}{util.GenerateWiredNetworkResponse("uid", "")},
						),
						1,
					)
				},
				ResourceName:      "uxi_wired_network.my_wired_network",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a wired_network is not allowed
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"uid",
						util.GenerateWiredNetworkPaginatedResponse(
							[]map[string]interface{}{util.GenerateWiredNetworkResponse("uid", "")},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_wired_network" "my_wired_network" {
					alias = "updated_alias"
				}`,
				ExpectError: regexp.MustCompile(`(?s)updating a wired_network is not supported; wired_networks can only be updated\s*through the dashboard`),
			},
			// Deleting a wired_network is not allowed
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"uid",
						util.GenerateWiredNetworkPaginatedResponse(
							[]map[string]interface{}{util.GenerateWiredNetworkResponse("uid", "")},
						),
						2,
					)
				},
				Config:      provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(`(?s)deleting a wired_network is not supported; wired_networks can only removed\s*from state`),
			},
			// Remove wired_network from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_wired_network.my_wired_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
