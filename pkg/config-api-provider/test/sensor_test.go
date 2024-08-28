package test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestSensorResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a sensor is not allowed
			{
				Config: providerConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "note"
						pcap_mode = "light"
					}`,

				ExpectError: regexp.MustCompile(`creating a sensor is not supported; sensors can only be imported`),
			},
			// Importing a sensor
			{
				PreConfig: func() {
					resources.GetSensor = func(uid string) resources.SensorResponseModel {
						return resources.SensorResponseModel{
							UID:                uid,
							Serial:             "serial",
							Name:               "imported_name",
							ModelNumber:        "model_number",
							WifiMacAddress:     "wifi_mac_address",
							EthernetMacAddress: "ethernet_mac_address",
							AddressNote:        "imported_address_note",
							Longitude:          "imported_longitude",
							Latitude:           "imported_latitude",
							Notes:              "imported_notes",
							PCapMode:           "light",
						}
					}
				},
				Config: providerConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "imported_name"
						address_note = "imported_address_note"
						notes = "imported_notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "uid"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "name", "imported_name"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "address_note", "imported_address_note"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "notes", "imported_notes"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "pcap_mode", "light"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "id", "uid"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_sensor.my_sensor",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					resources.GetSensor = func(uid string) resources.SensorResponseModel {
						return resources.SensorResponseModel{
							UID:                uid,
							Serial:             "serial",
							Name:               "updated_name",
							ModelNumber:        "model_number",
							WifiMacAddress:     "wifi_mac_address",
							EthernetMacAddress: "ethernet_mac_address",
							AddressNote:        "updated_address_note",
							Longitude:          "updated_longitude",
							Latitude:           "updated_latitude",
							Notes:              "updated_notes",
							PCapMode:           "not_light",
						}
					}
				},
				Config: providerConfig + `
				resource "uxi_sensor" "my_sensor" {
					name = "updated_name"
					address_note = "updated_address_note"
					notes = "updated_notes"
					pcap_mode = "not_light"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "name", "updated_name"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "address_note", "updated_address_note"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "notes", "updated_notes"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "pcap_mode", "not_light"),
				),
			},
			// Deleting a sensor is not allowed
			{
				Config:      providerConfig + ``,
				ExpectError: regexp.MustCompile(`deleting a sensor is not supported; sensors can only removed from state`),
			},
			// Remove sensor from state
			{
				Config: providerConfig + `
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
