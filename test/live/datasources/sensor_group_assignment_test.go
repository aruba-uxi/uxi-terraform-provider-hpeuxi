package data_source_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestSensorGroupAssignmentDataSource(t *testing.T) {
	const groupName = "tf_provider_acceptance_test_sensor_group_assignment_datasource"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					// create the resource to be used as a datasource
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					data "uxi_sensor" "my_sensor" {
						filter = {
							sensor_id = "` + config.SensorUid + `"
						}
					}

					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id = data.uxi_sensor.my_sensor.id
						group_id   = uxi_group.my_group.id
					}

					// the actual datasource
					data "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						filter = {
							sensor_group_assignment_id = uxi_sensor_group_assignment.my_sensor_group_assignment.id
						}
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						resourceName := "uxi_sensor_group_assignment.my_sensor_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						return util.CheckStateAgainstSensorGroupAssignment(
							t,
							"data.uxi_sensor_group_assignment.my_sensor_group_assignment",
							util.GetSensorGroupAssignment(rs.Primary.ID),
						)(s)
					},
				),
			},
		},
	})
}
