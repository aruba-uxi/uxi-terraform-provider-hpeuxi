package test

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestWirelessNetworkDataSource(t *testing.T) {
	defer gock.Off()
	MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					MockGetWirelessNetwork(
						"uid",
						GenerateWirelessNetworkPaginatedResponse([]map[string]interface{}{GenerateWirelessNetworkResponse("uid", "")}),
						3,
					)
				},
				Config: providerConfig + `
					data "uxi_wireless_network" "my_wireless_network" {
						filter = {
							wireless_network_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_wireless_network.my_wireless_network", "id", "uid"),
				),
			},
		},
	})
}
