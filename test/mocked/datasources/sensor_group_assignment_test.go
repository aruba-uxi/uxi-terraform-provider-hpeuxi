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

func TestSensorGroupAssignmentDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					util.MockGetSensorGroupAssignment(
						"id",
						util.GenerateSensorGroupAssignmentResponse("id", ""),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						filter = {
							sensor_group_assignment_id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.uxi_sensor_group_assignment.my_sensor_group_assignment",
						"id",
						"id",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_sensor_group_assignment.my_sensor_group_assignment",
						"group_id",
						"group_id",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_sensor_group_assignment.my_sensor_group_assignment",
						"sensor_id",
						"sensor_id",
					),
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestSensorGroupAssignmentDataSourceTooManyRequestsHandling(t *testing.T) {
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
						Get(shared.SensorGroupAssignmentPath).
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetSensorGroupAssignment(
						"id",
						util.GenerateSensorGroupAssignmentResponse("id", ""),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						filter = {
							sensor_group_assignment_id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.uxi_sensor_group_assignment.my_sensor_group_assignment",
						"id",
						"id",
					),
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

func TestSensorGroupAssignmentDataSourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					gock.New(util.MockUXIURL).
						Get(shared.SensorGroupAssignmentPath).
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						filter = {
							sensor_group_assignment_id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			{
				PreConfig: func() {
					util.MockGetSensorGroupAssignment("id", util.EmptyGetListResponse, 1)
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						filter = {
							sensor_group_assignment_id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Could not find specified data source`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
