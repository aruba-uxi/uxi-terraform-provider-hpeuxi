/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package data_source_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func TestWirelessNetworkDataSource(t *testing.T) {
	wirelessNetwork := util.GetWirelessNetwork(config.WirelessNetworkID)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					data "hpeuxi_wireless_network" "my_wireless_network" {
						filter = {
							id = "` + config.WirelessNetworkID + `"
						}
					}
				`,
				Check: shared.CheckStateAgainstWirelessNetwork(
					t,
					"data.hpeuxi_wireless_network.my_wireless_network",
					wirelessNetwork,
				),
			},
		},
	})
}
