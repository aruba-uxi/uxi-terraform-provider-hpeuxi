/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package data_source_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSensorDataSource(t *testing.T) {
	sensor := util.GetSensor(config.SensorId)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					data "uxi_sensor" "my_sensor" {
						filter = {
							sensor_id = "` + config.SensorId + `"
						}
					}
				`,
				Check: shared.CheckStateAgainstSensor(t, "data.uxi_sensor.my_sensor", sensor),
			},
		},
	})
}
