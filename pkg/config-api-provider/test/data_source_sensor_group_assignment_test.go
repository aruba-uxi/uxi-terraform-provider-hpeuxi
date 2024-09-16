package test

import (
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSensorGroupAssignmentDataSource(t *testing.T) {
	defer gock.Off()
	MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					MockGetSensorGroupAssignment(
						"uid",
						GenerateSensorGroupAssignmentPaginatedResponse([]map[string]interface{}{
							StructToMap(GenerateSensorGroupAssignmentResponse("uid", "")),
						}),
						3,
					)
				},
				Config: providerConfig + `
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
}
