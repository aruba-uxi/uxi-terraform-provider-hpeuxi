package data_source_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSensorDataSource(t *testing.T) {
	sensor := util.GetSensorProperties(config.SensorUid)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					data "uxi_sensor" "my_sensor" {
						filter = {
							sensor_id = "` + config.SensorUid + `"
						}
					}
				`,
				Check: util.CheckDataSourceStateAgainstSensor(t, "data.uxi_sensor.my_sensor", sensor),
			},
		},
	})
}
