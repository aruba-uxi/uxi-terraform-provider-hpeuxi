package test

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestWiredNetworkDataSource(t *testing.T) {
	defer gock.Off()
	MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					MockGetWiredNetwork(
						"uid",
						GenerateWiredNetworkPaginatedResponse([]map[string]interface{}{GenerateWiredNetworkResponse("uid", "")}),
						3,
					)
				},
				Config: providerConfig + `
					data "uxi_wired_network" "my_wired_network" {
						filter = {
							wired_network_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_wired_network.my_wired_network", "id", "uid"),
				),
			},
		},
	})
}
