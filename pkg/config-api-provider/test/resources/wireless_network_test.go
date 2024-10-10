package resource_test

import (
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWirelessNetworkResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a wireless_network is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "uxi_wireless_network" "my_wireless_network" {
						name = "name"
					}`,

				ExpectError: regexp.MustCompile(`(?s)creating a wireless_network is not supported; wireless_networks can only be\s*imported`),
			},
			// Importing a wireless_network
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateWirelessNetworkResponse("uid", "")}),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_wireless_network" "my_wireless_network" {
						name = "name"
					}

					import {
						to = uxi_wireless_network.my_wireless_network
						id = "uid"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_wireless_network.my_wireless_network", "name", "name"),
					resource.TestCheckResourceAttr("uxi_wireless_network.my_wireless_network", "id", "uid"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateWirelessNetworkResponse("uid", "")}),
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
					util.MockGetWirelessNetwork(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateWirelessNetworkResponse("uid", "")}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_wireless_network" "my_wireless_network" {
					name = "updated_name"
				}`,
				ExpectError: regexp.MustCompile(`(?s)updating a wireless_network is not supported; wireless_networks can only be\s*updated through the dashboard`),
			},
			// Deleting a wireless_network is not allowed
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateWirelessNetworkResponse("uid", "")}),
						1,
					)
				},
				Config:      provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(`(?s)deleting a wireless_network is not supported; wireless_networks can only\s*removed from state`),
			},
			// Remove wireless_network from state
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateWirelessNetworkResponse("uid", "")}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_wireless_network.my_wireless_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestWirelessNetworkResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_wireless_network" "my_wireless_network" {
						name = "name"
					}

					import {
						to = uxi_wireless_network.my_wireless_network
						id = "uid"
					}`,
				ExpectError: regexp.MustCompile(`Could not find specified resource`),
			},
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/uxi/v1alpha1/wireless-networks").
						Reply(500).
						JSON(map[string]interface{}{
							"httpStatusCode": 500,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_wireless_network" "my_wireless_network" {
						name = "name"
					}

					import {
						to = uxi_wireless_network.my_wireless_network
						id = "uid"
					}`,
				ExpectError: regexp.MustCompile(`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
