package data_source_test

import (
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
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
						util.GenerateSensorGroupAssignmentPaginatedResponse([]map[string]interface{}{
							util.StructToMap(util.GenerateSensorGroupAssignmentResponse("uid", "")),
						}),
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
