/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/stretchr/testify/assert"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func TestSensorResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	sensor := util.GenerateSensorsGetResponse("id", "").Items[0]
	updatedSensor := util.GenerateSensorsGetResponse("id", "_2").Items[0]

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
					resource "hpeuxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "note"
						pcap_mode = "light"
					}`,

				ExpectError: regexp.MustCompile(
					`creating a sensor is not supported; sensors can only be imported`,
				),
			},
			// Importing a sensor
			{
				PreConfig: func() {
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("id", ""), 2)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "id"
					}`,

				Check: shared.CheckStateAgainstSensor(t, "hpeuxi_sensor.my_sensor", sensor),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("id", ""), 1)
				},
				ResourceName:      "hpeuxi_sensor.my_sensor",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				PreConfig: func() {
					// existing sensor
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("id", ""), 1)
					util.MockUpdateSensor(
						"id",
						util.GenerateSensorPatchRequest("_2"),
						util.GenerateSensorPatchResponse("id", "_2"),
						1,
					)
					// updated sensor
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("id", "_2"), 1)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_sensor" "my_sensor" {
					name = "name_2"
					address_note = "address_note_2"
					notes = "notes_2"
					pcap_mode = "light"
				}`,
				Check: shared.CheckStateAgainstSensor(t, "hpeuxi_sensor.my_sensor", updatedSensor),
			},
			// Deleting a sensor is not allowed
			{
				PreConfig: func() {
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("id", "_2"), 1)
				},
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`deleting a sensor is not supported; sensors can only removed from state`,
				),
			},
			// Remove sensor from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_SensorResource_WithInvalidPcapMode_ShouldFail(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// sensor with invalid pcap_mode
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_sensor" "my_sensor" {
						name 		 = "name"
						address_note = "address_note"
						notes 		 = "notes"
						pcap_mode 	 = "invalid_pcap_mode"
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Attribute pcap_mode value must be one of: \["light" "full" "off"\], got:\s*"invalid_pcap_mode"`,
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestSensorResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	sensor := util.GenerateSensorsGetResponse("id", "").Items[0]
	updatedSensor := util.GenerateSensorsGetResponse("id", "_2").Items[0]
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Importing a sensor
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Get(shared.SensorPath).
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("id", ""), 2)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					shared.CheckStateAgainstSensor(t, "hpeuxi_sensor.my_sensor", sensor),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Update and Read testing
			{
				PreConfig: func() {
					// existing sensor
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("id", ""), 1)
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Patch("/networking-uxi/v1alpha1/sensors/id").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockUpdateSensor(
						"id",
						util.GenerateSensorPatchRequest("_2"),
						util.GenerateSensorPatchResponse("id", "_2"),
						1,
					)
					// updated sensor
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("id", "_2"), 1)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_sensor" "my_sensor" {
					name = "name_2"
					address_note = "address_note_2"
					notes = "notes_2"
					pcap_mode = "light"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					shared.CheckStateAgainstSensor(t, "hpeuxi_sensor.my_sensor", updatedSensor),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Remove sensor from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestSensorResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	sensor := util.GenerateSensorsGetResponse("id", "").Items[0]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Read HTTP error
			{
				PreConfig: func() {
					gock.New(util.MockUXIURL).
						Get(shared.SensorPath).
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "id"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Read not found
			{
				PreConfig: func() {
					util.MockGetSensor("id", util.EmptyGetListResponse, 1)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "id"
					}`,

				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Actually import a sensor for subsequent testing
			{
				PreConfig: func() {
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("id", ""), 2)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_sensor" "my_sensor" {
						name = "name"
						address_note = "address_note"
						notes = "notes"
						pcap_mode = "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "id"
					}`,

				Check: shared.CheckStateAgainstSensor(t, "hpeuxi_sensor.my_sensor", sensor),
			},
			// Update HTTP error
			{
				PreConfig: func() {
					// existing sensor
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("id", ""), 1)
					// patch sensor - with error
					gock.New(util.MockUXIURL).
						Patch("/networking-uxi/v1alpha1/sensors/id").
						Reply(http.StatusUnprocessableEntity).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusUnprocessableEntity,
							"errorCode":      "HPE_GL_UXI_INVALID_PCAP_MODE_ERROR",
							"message":        "Unable to update sensor - pcap_mode must be one the following ['light', 'full', 'off'].",
							"debugId":        "debugId",
							"type":           "hpe.greenlake.uxi.invalid_pcap_mode",
						})
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_sensor" "my_sensor" {
					name = "name_2"
					address_note = "address_note_2"
					notes = "notes_2"
					pcap_mode = "light"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)Unable to update sensor - pcap_mode must be one the following \['light',\s*'full', 'off'\].\s*DebugID: debugId`,
				),
			},
			// Remove sensor from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
