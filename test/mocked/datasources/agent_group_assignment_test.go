/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package data_source_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func TestAgentGroupAssignmentDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					util.MockGetAgentGroupAssignment(
						"id",
						util.GenerateAgentGroupAssignmentsGetResponse("id", ""),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_agent_group_assignment" "my_agent_group_assignment" {
						filter = {
							id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.hpeuxi_agent_group_assignment.my_agent_group_assignment",
						"id",
						"id",
					),
					resource.TestCheckResourceAttr(
						"data.hpeuxi_agent_group_assignment.my_agent_group_assignment",
						"group_id",
						"group_id",
					),
					resource.TestCheckResourceAttr(
						"data.hpeuxi_agent_group_assignment.my_agent_group_assignment",
						"agent_id",
						"agent_id",
					),
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestAgentGroupAssignmentDataSourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Get(shared.AgentGroupAssignmentPath).
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetAgentGroupAssignment(
						"id",
						util.GenerateAgentGroupAssignmentsGetResponse("id", ""),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_agent_group_assignment" "my_agent_group_assignment" {
						filter = {
							id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.hpeuxi_agent_group_assignment.my_agent_group_assignment",
						"id",
						"id",
					),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestAgentGroupAssignmentDataSourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					gock.New(util.MockUXIURL).
						Get(shared.AgentGroupAssignmentPath).
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_agent_group_assignment" "my_agent_group_assignment" {
						filter = {
							id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			{
				PreConfig: func() {
					util.MockGetAgentGroupAssignment("id", util.EmptyGetListResponse, 1)
				},
				Config: provider.ProviderConfig + `
					data "hpeuxi_agent_group_assignment" "my_agent_group_assignment" {
						filter = {
							id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Could not find specified data source`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
