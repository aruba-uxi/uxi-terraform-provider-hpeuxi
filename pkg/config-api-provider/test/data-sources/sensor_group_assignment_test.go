package data_source_test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/nbio/st"
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
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("uid", "")}),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						filter = {
							sensor_group_assignment_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_sensor_group_assignment.my_sensor_group_assignment", "id", "uid"),
					resource.TestCheckResourceAttr("data.uxi_sensor_group_assignment.my_sensor_group_assignment", "group_id", "group_uid"),
					resource.TestCheckResourceAttr("data.uxi_sensor_group_assignment.my_sensor_group_assignment", "sensor_id", "sensor_uid"),
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestSensorGroupAssignmentDataSource429Handling(t *testing.T) {
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
						Get("/uxi/v1alpha1/sensor-group-assignments").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetSensorGroupAssignment(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("uid", "")}),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						filter = {
							sensor_group_assignment_id = "uid"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.uxi_sensor_group_assignment.my_sensor_group_assignment", "id", "uid"),
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
func TestSensorGroupAssignmentDataSourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/uxi/v1alpha1/sensor-group-assignments").
						Reply(500).
						JSON(map[string]interface{}{
							"httpStatusCode": 500,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					data "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						filter = {
							sensor_group_assignment_id = "uid"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
