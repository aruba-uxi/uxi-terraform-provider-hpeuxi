package data_source_test

import (
	"fmt"
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

func TestGroupDataSource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			// Test Read, is_root not set
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/configuration/app/v1/groups").
						Reply(429).
						SetHeaders(map[string]string{
							"X-RateLimit-Limit":     "100",
							"X-RateLimit-Remaining": "0",
							"X-RateLimit-Reset":     "1",
						})
					util.MockGetGroup(
						"uid",
						util.GenerateGroupPaginatedResponse([]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseModel("uid", "", ""))}),
						2,
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
					PendingMockCheck{t: t, mocksToExclude: []gock.Mock{mockOAuth.Mock}}.AssertNoPendingMocks(),
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func removeElements(slice1, slice2 []gock.Mock) []gock.Mock {
	result := []gock.Mock{}
	for _, v1 := range slice1 {
		found := false
		for _, v2 := range slice2 {
			if v1 == v2 {
				found = true
				break
			}
		}
		if !found {
			result = append(result, v1)
		}
	}
	return result
}

type PendingMockCheck struct {
	t              st.Fatalf
	mocksToExclude []gock.Mock
}

func (p PendingMockCheck) AssertNoPendingMocks() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		stillPending := removeElements(gock.Pending(), p.mocksToExclude)

		if len(stillPending) > 0 {
			for _, mock := range stillPending {
				fmt.Printf(
					"Mock still pending: %s %s\tpending calls: %d, done: %t\n",
					mock.Request().Method,
					mock.Request().URLStruct,
					mock.Request().Counter,
					mock.Done(),
				)
			}
		}

		st.Assert(p.t, len(stillPending), 0)

		return nil
	}
}
