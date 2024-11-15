/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"

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
						name = "name"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)creating a wired_network is not supported; wired_networks can only be\s*imported`,
				),
			},
			// Importing a wired_network
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateWiredNetworkResponse("id", "")},
						),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_wired_network" "my_wired_network" {
						name = "name"
					}

					import {
						to = uxi_wired_network.my_wired_network
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_wired_network.my_wired_network",
						"name",
						"name",
					),
					resource.TestCheckResourceAttr(
						"uxi_wired_network.my_wired_network",
						"id",
						"id",
					),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateWiredNetworkResponse("id", "")},
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
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateWiredNetworkResponse("id", "")},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_wired_network" "my_wired_network" {
					name = "updated_name"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)updating a wired_network is not supported; wired_networks can only be updated\s*through the dashboard`,
				),
			},
			// Deleting a wired_network is not allowed
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateWiredNetworkResponse("id", "")},
						),
						2,
					)
				},
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`(?s)deleting a wired_network is not supported; wired_networks can only removed\s*from state`,
				),
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

func TestWiredNetworkResourceHttpErrorHandling(t *testing.T) {
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
					util.MockGetWiredNetwork(
						"id",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_wired_network" "my_wired_network" {
						name = "name"
					}

					import {
						to = uxi_wired_network.my_wired_network
						id = "id"
					}`,
				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/wired-networks").
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_wired_network" "my_wired_network" {
						name = "name"
					}

					import {
						to = uxi_wired_network.my_wired_network
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
