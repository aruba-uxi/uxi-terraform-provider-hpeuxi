package resource_test

import (
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestSensorResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

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
						return util.GenerateSensorResponseModel(uid, "")
					}
				},
				Config: provider.ProviderConfig + `
					resource "uxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "uid"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "name", "name"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "address_note", "address_note"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "notes", "notes"),
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
						return util.GenerateSensorResponseModel(uid, "_2")
					}
				},
				Config: provider.ProviderConfig + `
				resource "uxi_sensor" "my_sensor" {
					name = "name_2"
					address_note = "address_note_2"
					notes = "notes_2"
					pcap_mode = "light_2"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "name", "name_2"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "address_note", "address_note_2"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "notes", "notes_2"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "pcap_mode", "light_2"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "id", "uid"),
				),
			},
			// Deleting a sensor is not allowed
			{
				Config:      provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(`deleting a sensor is not supported; sensors can only removed from state`),
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

	mockOAuth.Mock.Disable()
}
