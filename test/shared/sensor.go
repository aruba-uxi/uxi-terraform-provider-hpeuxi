/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package shared

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func CheckStateAgainstSensor(
	t *testing.T,
	entity string,
	sensor config_api_client.SensorItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", sensor.Id),
		resource.TestCheckResourceAttr(entity, "serial", sensor.Serial),
		resource.TestCheckResourceAttr(entity, "model_number", sensor.ModelNumber),
		resource.TestCheckResourceAttrWith(
			entity,
			"name",
			func(value string) error {
				assert.Equal(t, value, sensor.Name)
				return nil
			},
		),
		TestOptionalValue(t, entity, "wifi_mac_address", sensor.WifiMacAddress.Get()),
		TestOptionalValue(
			t,
			entity,
			"ethernet_mac_address",
			sensor.EthernetMacAddress.Get(),
		),
		TestOptionalValue(t, entity, "address_note", sensor.AddressNote.Get()),
		TestOptionalFloatValue(t, entity, "latitude", sensor.Latitude.Get()),
		TestOptionalFloatValue(t, entity, "longitude", sensor.Longitude.Get()),
		TestOptionalValue(t, entity, "notes", sensor.Notes.Get()),
		TestOptionalValue(t, entity, "pcap_mode", sensor.PcapMode.Get()),
	)
}
