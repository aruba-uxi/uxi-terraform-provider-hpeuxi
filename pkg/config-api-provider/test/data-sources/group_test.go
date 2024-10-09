package data_source_test

import (
	config_api_client "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/nbio/st"
)

func TestGroupDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test Read
			{
				PreConfig: func() {
					util.MockGetGroup(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseGetModel("uid", "", ""))}),
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
			// TODO: Test retrieving the root group
			{
				PreConfig: func() {
					util.MockGetGroup(
						"my_root_group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.StructToMap(config_api_client.GroupsGetItem{
							Id:     "my_root_group_uid",
							Name:   "root",
							Parent: *config_api_client.NewNullableParent(nil),
							Path:   "my_root_group_uid",
						})}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_group" "my_group" {
						filter = {
							group_id = "my_root_group_uid"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`the root group cannot be used as a data source`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestGroupDataSource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mock429 *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			// Test Read
			{
				PreConfig: func() {
					mock429 = gock.New("https://test.api.capenetworks.com").
						Get("/uxi/v1alpha1/groups").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetGroup(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseGetModel("uid", "", ""))}),
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
					func(s *terraform.State) error {
						st.Assert(t, mock429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
