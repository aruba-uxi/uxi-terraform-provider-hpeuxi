package test

import (
	"testing"

	datasources "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/data-sources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestRootGroupDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					datasources.GetRootGroup = func() datasources.RootGroupResponseModel {
						return GenerateRootGroupResponseModel("mock_uid")
					}
				},
				Config: providerConfig + `
					data "uxi_root_group" "my_root_group" {}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_root_group.my_root_group", "id", "mock_uid"),
				),
			},
		},
	})
}
