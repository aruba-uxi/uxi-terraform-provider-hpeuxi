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

func CheckDataSourceStateAgainstSensor(
	t st.Fatalf,
	entity string,
	sensor config_api_client.SensorItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", config.SensorId),
		resource.TestCheckResourceAttr(entity, "serial", sensor.Serial),
		resource.TestCheckResourceAttr(entity, "model_number", sensor.ModelNumber),
		resource.TestCheckResourceAttrWith(
			entity,
			"name",
			func(value string) error {
				st.Assert(t, value, sensor.Name)
				return nil
			},
		),
		TestOptionalValue(t, entity, "wifi_mac_address", sensor.WifiMacAddress.Get()),
		TestOptionalValue(t, entity, "ethernet_mac_address", sensor.EthernetMacAddress.Get()),
		TestOptionalValue(t, entity, "address_note", sensor.AddressNote.Get()),
		TestOptionalFloatValue(t, entity, "latitude", sensor.Latitude.Get()),
		TestOptionalFloatValue(t, entity, "longitude", sensor.Longitude.Get()),
		TestOptionalValue(t, entity, "notes", sensor.Notes.Get()),
		TestOptionalValue(t, entity, "pcap_mode", sensor.PcapMode.Get()),
	)
}

func CheckResourceStateAgainstSensor(
	t st.Fatalf,
	entity string,
	sensor config_api_client.SensorItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", config.SensorId),
		resource.TestCheckResourceAttrWith(
			entity,
			"name",
			func(value string) error {
				st.Assert(t, value, sensor.Name)
				return nil
			},
		),
		TestOptionalValue(t, entity, "address_note", sensor.AddressNote.Get()),
		TestOptionalValue(t, entity, "notes", sensor.Notes.Get()),
		TestOptionalValue(t, entity, "pcap_mode", sensor.PcapMode.Get()),
	)
}
