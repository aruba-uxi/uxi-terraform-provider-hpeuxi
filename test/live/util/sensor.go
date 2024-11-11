package util

import (
	"context"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
)

func GetSensorProperties(id string) config_api_client.SensorItem {
	result, _, err := Client.ConfigurationAPI.
		SensorsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("sensor with id `" + id + "` could not be found")
	}
	return result.Items[0]
}

func CheckStateAgainstSensor(
	t st.Fatalf,
	sensor config_api_client.SensorItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.uxi_sensor.my_sensor", "id", config.SensorUid),
		resource.TestCheckResourceAttr("data.uxi_sensor.my_sensor", "serial", sensor.Serial),
		resource.TestCheckResourceAttr(
			"data.uxi_sensor.my_sensor",
			"model_number",
			sensor.ModelNumber,
		),
		resource.TestCheckResourceAttrWith(
			"data.uxi_sensor.my_sensor",
			"name",
			func(value string) error {
				st.Assert(t, value, sensor.Name)
				return nil
			},
		),
		TestOptionalValue(
			t,
			"data.uxi_sensor.my_sensor",
			"wifi_mac_address",
			sensor.WifiMacAddress.Get(),
		),
		TestOptionalValue(
			t,
			"data.uxi_sensor.my_sensor",
			"ethernet_mac_address",
			sensor.EthernetMacAddress.Get(),
		),
		TestOptionalValue(t, "data.uxi_sensor.my_sensor", "address_note", sensor.AddressNote.Get()),
		TestOptionalFloatValue(t, "data.uxi_sensor.my_sensor", "latitude", sensor.Latitude.Get()),
		TestOptionalFloatValue(t, "data.uxi_sensor.my_sensor", "longitude", sensor.Longitude.Get()),
		TestOptionalValue(t, "data.uxi_sensor.my_sensor", "notes", sensor.Notes.Get()),
		TestOptionalValue(t, "data.uxi_sensor.my_sensor", "pcap_mode", sensor.PcapMode.Get()),
	)
}
