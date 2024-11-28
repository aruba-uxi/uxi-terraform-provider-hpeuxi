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

func TestWiredNetworkDataSource(t *testing.T) {
	wiredNetwork := util.GetWiredNetwork(config.WiredNetworkID)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					data "uxi_wired_network" "my_wired_network" {
						filter = {
							id = "` + config.WiredNetworkID + `"
						}
					}
				`,
				Check: shared.CheckStateAgainstWiredNetwork(
					t,
					"data.uxi_wired_network.my_wired_network",
					wiredNetwork,
				),
			},
		},
	})
}
