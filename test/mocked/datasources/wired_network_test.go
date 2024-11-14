package data_source_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/nbio/st"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateWiredNetworkResponse("uid", "")},
						),
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
					resource.TestCheckResourceAttr(
						"data.uxi_wired_network.my_wired_network",
						"id",
						"uid",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_wired_network.my_wired_network",
						"name",
						"name",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_wired_network.my_wired_network",
						"ip_version",
						"ip_version",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_wired_network.my_wired_network",
						"security",
						"security",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_wired_network.my_wired_network",
						"dns_lookup_domain",
						"dns_lookup_domain",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_wired_network.my_wired_network",
						"disable_edns",
						"false",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_wired_network.my_wired_network",
						"use_dns64",
						"false",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_wired_network.my_wired_network",
						"external_connectivity",
						"false",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_wired_network.my_wired_network",
						"vlan_id",
						"123",
					),
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestWiredNetworkDataSource429Handling(t *testing.T) {
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
						Get("/networking-uxi/v1alpha1/wired-networks").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetWiredNetwork(
						"uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateWiredNetworkResponse("uid", "")},
						),
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
					resource.TestCheckResourceAttr(
						"data.uxi_wired_network.my_wired_network",
						"id",
						"uid",
					),
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

func TestWiredNetworkAssignmentDataSourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/wired-networks").
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					data "uxi_wired_network" "my_wired_network" {
						filter = {
							wired_network_id = "uid"
						}
					}
				`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			{
				PreConfig: func() {
					util.MockGetWiredNetwork(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_wired_network" "my_wired_network" {
						filter = {
							wired_network_id = "uid"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Could not find specified data source`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
