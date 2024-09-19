package data_source_test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGroupDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test no filters set
			{
				Config: provider.ProviderConfig + `
					data "uxi_group" "my_group" {
						filter = {}
					}
				`,
				ExpectError: regexp.MustCompile(`either filter.group_id must be set or 'filter.is_root = true' is required`),
			},
			// Test too many filters set
			{
				Config: provider.ProviderConfig + `
					data "uxi_group" "my_group" {
						filter = {
							is_root  = true
							group_id = "uid"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`group_id and 'is_root = true' cannot both be set`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_root_group.my_root_group", "id", "mock_uid"),
				),
			},
			// Test Read, is_root not set
			{
				PreConfig: func() {
					util.MockGetGroup(
						"uid",
						util.GenerateGroupPaginatedResponse([]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseModel("uid", "", ""))}),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_group" "my_group" {
						filter = {
							group_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_group.my_group", "id", "uid"),
				),
			},
			// Test Read, is_root is false
			{
				PreConfig: func() {
					util.MockGetGroup(
						"uid",
						util.GenerateGroupPaginatedResponse([]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseModel("uid", "", ""))}),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_group" "my_group" {
						filter = {
							is_root  = false
							group_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_group.my_group", "id", "uid"),
				),
			},
			// TODO: Test retrieving the root group
		},
	})

	mockOAuth.Mock.Disable()
}
