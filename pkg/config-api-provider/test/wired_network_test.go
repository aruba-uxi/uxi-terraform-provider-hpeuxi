package test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestWiredNetworkResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a wired_network is not allowed
			{
				Config: providerConfig + `
					resource "uxi_wired_network" "my_wired_network" {
						alias = "alias"
					}`,

				ExpectError: regexp.MustCompile(`(?s)creating a wired_network is not supported; wired_networks can only be\s*imported`),
			},
			// Importing a wired_network
			{
				PreConfig: func() {
					resources.GetWiredNetwork = func() resources.WiredNetworkResponseModel {
						return resources.WiredNetworkResponseModel{
							Uid:                  "uid",
							Alias:                "alias",
							DatetimeCreated:      "datetime_created",
							DatetimeUpdated:      "datetime_updated",
							IpVersion:            "ip_version",
							Security:             "security",
							DnsLookupDomain:      "dns_lookup_domain",
							DisableEdns:          false,
							UseDns64:             false,
							ExternalConnectivity: false,
							VlanId:               123,
						}
					}
				},
				Config: providerConfig + `
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
				ResourceName:      "uxi_wired_network.my_wired_network",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a wired_network is not allowed
			{
				Config: providerConfig + `
				resource "uxi_wired_network" "my_wired_network" {
					alias = "updated_alias"
				}`,
				ExpectError: regexp.MustCompile(`(?s)updating a wired_network is not supported; wired_networks can only be updated\s*through the dashboard`),
			},
			// Deleting a wired_network is not allowed
			{
				Config:      providerConfig + ``,
				ExpectError: regexp.MustCompile(`(?s)deleting a wired_network is not supported; wired_networks can only removed\s*from state`),
			},
			// Remove wired_network from state
			{
				Config: providerConfig + `
					removed {
						from = uxi_wired_network.my_wired_network

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}
