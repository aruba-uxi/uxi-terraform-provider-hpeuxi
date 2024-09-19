package data_source_test

import (
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestWiredNetworkDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"uid",
						util.GenerateWiredNetworkPaginatedResponse([]map[string]interface{}{util.GenerateWiredNetworkResponse("uid", "")}),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_wired_network" "my_wired_network" {
						filter = {
							wired_network_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_wired_network.my_wired_network", "id", "uid"),
					resource.TestCheckResourceAttr("data.uxi_wired_network.my_wired_network", "alias", "alias"),
					resource.TestCheckResourceAttr("data.uxi_wired_network.my_wired_network", "ip_version", "ip_version"),
					resource.TestCheckResourceAttr("data.uxi_wired_network.my_wired_network", "security", "security"),
					resource.TestCheckResourceAttr("data.uxi_wired_network.my_wired_network", "dns_lookup_domain", "dns_lookup_domain"),
					resource.TestCheckResourceAttr("data.uxi_wired_network.my_wired_network", "disable_edns", "false"),
					resource.TestCheckResourceAttr("data.uxi_wired_network.my_wired_network", "use_dns64", "false"),
					resource.TestCheckResourceAttr("data.uxi_wired_network.my_wired_network", "external_connectivity", "false"),
					resource.TestCheckResourceAttr("data.uxi_wired_network.my_wired_network", "vlan_id", "123"),
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
