package util

import (
	"context"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
)

type SensorProperties struct {
	Id                 string
	Serial             string
	Name               string
	ModelNumber        string
	WifiMacAddress     *string
	EthernetMacAddress *string
	Latitude           *float32
	Longitude          *float32
	Notes              *string
	AddressNote        *string
	PcapMode           *string
}

func GetSensorProperties(id string) SensorProperties {
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
	sensor := result.Items[0]
	// Read these in, as they may not be always constant with the acceptance test
	// customer
	return SensorProperties{
		Id:                 sensor.Id,
		Serial:             sensor.Serial,
		Name:               sensor.Name,
		ModelNumber:        sensor.ModelNumber,
		WifiMacAddress:     sensor.WifiMacAddress.Get(),
		EthernetMacAddress: sensor.EthernetMacAddress.Get(),
		Latitude:           sensor.Latitude.Get(),
		Longitude:          sensor.Longitude.Get(),
		Notes:              sensor.Notes.Get(),
		AddressNote:        sensor.AddressNote.Get(),
		PcapMode:           sensor.PcapMode.Get(),
	}
}

func CheckStateAgainstSensor(t st.Fatalf, sensor SensorProperties) resource.TestCheckFunc {
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
			sensor.WifiMacAddress,
		),
		TestOptionalValue(
			t,
			"data.uxi_sensor.my_sensor",
			"ethernet_mac_address",
			sensor.EthernetMacAddress,
		),
		TestOptionalValue(t, "data.uxi_sensor.my_sensor", "address_note", sensor.AddressNote),
		TestOptionalFloatValue(t, "data.uxi_sensor.my_sensor", "latitude", sensor.Latitude),
		TestOptionalFloatValue(t, "data.uxi_sensor.my_sensor", "longitude", sensor.Longitude),
		TestOptionalValue(t, "data.uxi_sensor.my_sensor", "notes", sensor.Notes),
		TestOptionalValue(t, "data.uxi_sensor.my_sensor", "pcap_mode", sensor.PcapMode),
	)
}
