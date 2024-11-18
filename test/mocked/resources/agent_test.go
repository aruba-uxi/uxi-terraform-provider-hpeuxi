package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/stretchr/testify/assert"
)

func TestAgentResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating an agent is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name = "name"
						notes = "note"
						pcap_mode = "light"
					}`,

				ExpectError: regexp.MustCompile(
					`creating an agent is not supported; agents can only be imported`,
				),
			},
			// Importing an agent
			{
				PreConfig: func() {
					util.MockGetAgent(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateAgentResponse("id", "")},
						),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name = "name"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", "name"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", "notes"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "pcap_mode", "light"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "id", "id"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetAgent(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateAgentResponse("id", "")},
						),
						1,
					)
				},
				ResourceName:      "uxi_agent.my_agent",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					// original
					util.MockGetAgent(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateAgentResponse("id", "")},
						),
						1,
					)
					util.MockUpdateAgent(
						"id",
						util.GenerateAgentUpdateRequest("_2"),
						util.GenerateAgentResponse("id", "_2"),
						1,
					)
					// updated
					util.MockGetAgent(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateAgentResponse("id", "_2")},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_agent" "my_agent" {
					name = "name_2"
					notes = "notes_2"
					pcap_mode = "light"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", "name_2"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "notes", "notes_2"),
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "pcap_mode", "light"),
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					util.MockGetAgent("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateAgentResponse("id", "")}),
						1,
					)
					util.MockDeleteAgent("id", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestAgentResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Importing a agent
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/agents").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetAgent("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateAgentResponse("id", "")}),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name = "name"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "id", "id"),
					func(s *terraform.State) error {
						assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Update testing
			{
				PreConfig: func() {
					// original
					util.MockGetAgent(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateAgentResponse("id", "")},
						),
						1,
					)
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Patch("/networking-uxi/v1alpha1/agents").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockUpdateAgent(
						"id",
						util.GenerateAgentUpdateRequest("_2"),
						util.GenerateAgentResponse("id", "_2"),
						1,
					)
					// updated
					util.MockGetAgent(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateAgentResponse("id", "_2")},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_agent" "my_agent" {
					name = "name_2"
					notes = "notes_2"
					pcap_mode = "light"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "name", "name_2"),
					func(s *terraform.State) error {
						assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					util.MockGetAgent("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateAgentResponse("id", "")}),
						1,
					)
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/agents/id").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockDeleteAgent("id", 1)
				},
				Config: provider.ProviderConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
func TestAgentResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Read 5xx error
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/agents").
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name = "name"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "id"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Read not found
			{
				PreConfig: func() {
					util.MockGetAgent(
						"id",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name = "name"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "id"
					}`,

				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Actually importing an agent for testing purposes
			{
				PreConfig: func() {
					util.MockGetAgent(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateAgentResponse("id", "")},
						),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_agent" "my_agent" {
						name = "name"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_agent.my_agent
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_agent.my_agent", "id", "id"),
				),
			},
			// update 4xx
			{
				PreConfig: func() {
					// original
					util.MockGetAgent(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{util.GenerateAgentResponse("id", "")},
						),
						1,
					)
					// patch agent - with error
					gock.New("https://test.api.capenetworks.com").
						Patch("/networking-uxi/v1alpha1/agents/id").
						Reply(http.StatusUnprocessableEntity).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusUnprocessableEntity,
							"errorCode":      "HPE_GL_UXI_INVALID_PCAP_MODE_ERROR",
							"message":        "Unable to update agent - pcap_mode must be one the following ['light', 'full', 'off'].",
							"debugId":        "12312-123123-123123-1231212",
							"type":           "hpe.greenlake.uxi.invalid_pcap_mode",
						})
				},
				Config: provider.ProviderConfig + `
				resource "uxi_agent" "my_agent" {
					name = "name_2"
					notes = "notes_2"
					pcap_mode = "light"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)Unable to update agent - pcap_mode must be one the following \['light',\s*'full', 'off'\].\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Delete 4xx
			{
				PreConfig: func() {
					// existing agent
					util.MockGetAgent("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateAgentResponse("id", "")}),
						1,
					)
					// delete agent - with error
					gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/agents/id").
						Reply(http.StatusUnprocessableEntity).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusUnprocessableEntity,
							"errorCode":      "HPE_GL_NETWORKING_UXI_HARDWARE_SENSOR_DELETION_FORBIDDEN",
							"message":        "Cant delete sensor - hardware sensor deletion is forbidden",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig,
				ExpectError: regexp.MustCompile(
					`(?s)Cant delete sensor - hardware sensor deletion is forbidden\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Actually delete group for cleanup reasons
			{
				PreConfig: func() {
					// existing group
					util.MockGetAgent("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateAgentResponse("id", "")}),
						1,
					)
					// delete group
					util.MockDeleteAgent("id", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
