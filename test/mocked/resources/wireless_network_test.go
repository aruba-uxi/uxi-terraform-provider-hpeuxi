package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/stretchr/testify/assert"
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

				ExpectError: regexp.MustCompile(
					`(?s)creating a wireless_network is not supported; wireless_networks can only be\s*imported`,
				),
			},
			// Importing a wireless_network
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("id", ""),
							},
						),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_wireless_network" "my_wireless_network" {
						name = "name"
					}

					import {
						to = uxi_wireless_network.my_wireless_network
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_wireless_network.my_wireless_network",
						"name",
						"name",
					),
					resource.TestCheckResourceAttr(
						"uxi_wireless_network.my_wireless_network",
						"id",
						"id",
					),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("id", ""),
							},
						),
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
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("id", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_wireless_network" "my_wireless_network" {
					name = "updated_name"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)updating a wireless_network is not supported; wireless_networks can only be\s*updated through the dashboard`,
				),
			},
			// Deleting a wireless_network is not allowed
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("id", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`(?s)deleting a wireless_network is not supported; wireless_networks can only\s*removed from state`,
				),
			},
			// Remove wireless_network from state
			{
				PreConfig: func() {
					util.MockGetWirelessNetwork(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("id", ""),
							},
						),
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

func TestWirelessNetworkResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Importing a service_test
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New(util.MockUxiUrl).
						Get(shared.WirelessNetworkPath).
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetWirelessNetwork(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateWirelessNetworkResponse("id", ""),
							},
						),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_wireless_network" "my_wireless_network" {
						name = "name"
					}

					import {
						to = uxi_wireless_network.my_wireless_network
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_wireless_network.my_wireless_network",
						"id",
						"id",
					),
					func(s *terraform.State) error {
						assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Remove service_test from state
			{
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
						"id",
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
						id = "id"
					}`,
				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			{
				PreConfig: func() {
					gock.New(util.MockUxiUrl).
						Get(shared.WirelessNetworkPath).
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
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
						id = "id"
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
