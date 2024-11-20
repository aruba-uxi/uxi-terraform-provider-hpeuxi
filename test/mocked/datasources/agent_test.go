/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package data_source_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAgentDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test Read
			{
				PreConfig: func() {
					util.MockGetAgent("id", util.GenerateAgentResponse("id", ""), 3)
				},
				Config: provider.ProviderConfig + `
					data "uxi_agent" "my_agent" {
						filter = {
							agent_id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "id", "id"),
					resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "name", "name"),
					resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "serial", "serial"),
					resource.TestCheckResourceAttr(
						"data.uxi_agent.my_agent",
						"model_number",
						"model_number",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_agent.my_agent",
						"wifi_mac_address",
						"wifi_mac_address",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_agent.my_agent",
						"ethernet_mac_address",
						"ethernet_mac_address",
					),
					resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "notes", "notes"),
					resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "pcap_mode", "light"),
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestAgentDataSourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			// Test Read
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New(util.MockUxiUrl).
						Get(shared.AgentPath).
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetAgent("id", util.GenerateAgentResponse("id", ""), 3)
				},
				Config: provider.ProviderConfig + `
					data "uxi_agent" "my_agent" {
						filter = {
							agent_id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "id", "id"),
					func(s *terraform.State) error {
						assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestAgentDataSourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// HTTP error
			{
				PreConfig: func() {
					gock.New(util.MockUxiUrl).
						Get(shared.AgentPath).
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					data "uxi_agent" "my_agent" {
						filter = {
							agent_id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Not found error
			{
				PreConfig: func() {
					util.MockGetAgent("id", util.EmptyGetListResponse, 1)
				},
				Config: provider.ProviderConfig + `
					data "uxi_agent" "my_agent" {
						filter = {
							agent_id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Could not find specified data source`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
