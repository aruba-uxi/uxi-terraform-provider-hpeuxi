package data_source_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestWiredNetworkDataSource(t *testing.T) {
	wired_network := util.GetWiredNetwork(config.WiredNetworkId)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					data "uxi_wired_network" "my_wired_network" {
						filter = {
							wired_network_id = "` + config.WiredNetworkId + `"
						}
					}
				`,
				Check: util.CheckStateAgainstWiredNetwork(t, wired_network),
			},
		},
	})
}
