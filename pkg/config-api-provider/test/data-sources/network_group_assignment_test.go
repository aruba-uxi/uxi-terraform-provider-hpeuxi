package data_source_test

import (
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/nbio/st"
)

func TestNetworkGroupAssignmentDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					util.MockGetNetworkGroupAssignment(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateNetworkGroupAssignmentGetResponse("uid", "")}),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_network_group_assignment" "my_network_group_assignment" {
						filter = {
							network_group_assignment_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_network_group_assignment.my_network_group_assignment", "id", "uid"),
					resource.TestCheckResourceAttr("data.uxi_network_group_assignment.my_network_group_assignment", "group_id", "group_uid"),
					resource.TestCheckResourceAttr("data.uxi_network_group_assignment.my_network_group_assignment", "network_id", "network_uid"),
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestNetworkGroupAssignmentSource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mock429 *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					mock429 = gock.New("https://test.api.capenetworks.com").
						Get("/uxi/v1alpha1/network-group-assignments").
						Reply(429).
						SetHeaders(map[string]string{
							"X-RateLimit-Limit":     "100",
							"X-RateLimit-Remaining": "0",
							"X-RateLimit-Reset":     "1",
						})
					util.MockGetNetworkGroupAssignment(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateNetworkGroupAssignmentGetResponse("uid", "")}),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_network_group_assignment" "my_network_group_assignment" {
						filter = {
							network_group_assignment_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_network_group_assignment.my_network_group_assignment", "id", "uid"),
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
