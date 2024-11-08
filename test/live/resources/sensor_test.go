package resource_test

import (
	"context"

	"github.com/aruba-uxi/terraform-provider-configuration/test/live/config"
	"github.com/aruba-uxi/terraform-provider-configuration/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-configuration/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/nbio/st"

	"regexp"
	"testing"
)

type sensorProperties struct {
	id           string
	name         string
	notes        *string
	addressNotes *string
	pcapMode     *string
}

func TestSensorResource(t *testing.T) {
	originalSensor := getSensorProperties(config.SensorUid)
	updatedSensor := sensorProperties{
		id:           config.SensorUid,
		name:         "tf_provider_acceptance_test_update_name",
		notes:        originalSensor.notes,
		addressNotes: originalSensor.addressNotes,
		pcapMode:     originalSensor.pcapMode,
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a sensor is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "` + originalSensor.name + `"
					}`,

				ExpectError: regexp.MustCompile(
					`creating a sensor is not supported; sensors can only be imported`,
				),
			},
			// Importing a sensor
			{
				Config: provider.ProviderConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "` + originalSensor.name + `"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "` + config.SensorUid + `"
						}`,

				Check: resource.ComposeAggregateTestCheckFunc(),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_sensor.my_sensor",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update testing
			{
				Config: provider.ProviderConfig + `
				resource "uxi_sensor" "my_sensor" {
					name 		 = "` + updatedSensor.name + `"
					// address_note = "address_note_2"
					// notes 		 = "notes_2"
					// pcap_mode 	 = "light_2"
				}`,
				Check: checkStateAgainstSensor(t, updatedSensor),
			},
			// Update sensor back to original
			{
				Config: provider.ProviderConfig + `
				resource "uxi_sensor" "my_sensor" {
					name 		 = "` + originalSensor.name + `"
					// address_note = "address_note_2"
					// notes 		 = "notes_2"
					// pcap_mode 	 = "light_2"
				}`,
				Check: checkStateAgainstSensor(t, originalSensor),
			},
			// Deleting a sensor is not allowed
			{
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`deleting a sensor is not supported; sensors can only removed from state`,
				),
			},
			// Remove sensor from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})
}

func checkStateAgainstSensor(t st.Fatalf, sensor sensorProperties) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "id", config.SensorUid),
		resource.TestCheckResourceAttrWith(
			"uxi_sensor.my_sensor",
			"name",
			func(value string) error {
				st.Assert(t, value, sensor.name)
				return nil
			},
		),
		func() resource.TestCheckFunc {
			if sensor.addressNotes == nil {
				return resource.TestCheckNoResourceAttr(
					"uxi_sensor.my_sensor",
					"address_note",
				)
			} else {
				return resource.TestCheckResourceAttrWith(
					"uxi_sensor.my_sensor",
					"address_note",
					func(value string) error {
						st.Assert(t, value, sensor.addressNotes)
						return nil
					},
				)
			}
		}(),
		resource.TestCheckResourceAttrWith(
			"uxi_sensor.my_sensor",
			"notes",
			func(value string) error {
				st.Assert(t, value, sensor.notes)
				return nil
			},
		),
		func() resource.TestCheckFunc {
			if sensor.pcapMode == nil {
				return resource.TestCheckNoResourceAttr(
					"uxi_sensor.my_sensor",
					"pcap_mode",
				)
			} else {
				return resource.TestCheckResourceAttrWith(
					"uxi_sensor.my_sensor",
					"pcap_mode",
					func(value string) error {
						st.Assert(t, value, sensor.pcapMode)
						return nil
					},
				)
			}
		}(),
	)
}

func getSensorProperties(id string) sensorProperties {
	result, _, err := util.Client.ConfigurationAPI.
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
	return sensorProperties{
		id:           sensor.Id,
		name:         sensor.Name,
		notes:        sensor.Notes.Get(),
		addressNotes: sensor.AddressNote.Get(),
		pcapMode:     sensor.PcapMode.Get(),
	}
}
