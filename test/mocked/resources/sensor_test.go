package resource_test

import (
	"github.com/aruba-uxi/terraform-provider-configuration/internal/provider/resources"
	"github.com/aruba-uxi/terraform-provider-configuration/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-configuration/test/mocked/util"
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
	"github.com/nbio/st"
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

				ExpectError: regexp.MustCompile(
					`creating a sensor is not supported; sensors can only be imported`,
				),
			},
			// Importing a sensor
			{
				PreConfig: func() {
					util.MockGetSensor("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("uid", "")}),
						2,
					)
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
					resource.TestCheckResourceAttr(
						"uxi_sensor.my_sensor",
						"address_note",
						"address_note",
					),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "notes", "notes"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "pcap_mode", "light"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "id", "uid"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetSensor("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("uid", "")}),
						1,
					)
				},
				ResourceName:      "uxi_sensor.my_sensor",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					// existing sensor
					util.MockGetSensor("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("uid", "")}),
						1,
					)
					resources.UpdateSensor = func(request resources.SensorUpdateRequestModel) resources.SensorResponseModel {
						return util.GenerateMockedSensorResponseModel("uid", "_2")
					}
					// updated sensor
					util.MockGetSensor("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("uid", "_2")}),
						1,
					)
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
					resource.TestCheckResourceAttr(
						"uxi_sensor.my_sensor",
						"address_note",
						"address_note_2",
					),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "notes", "notes_2"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "pcap_mode", "light_2"),
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "id", "uid"),
				),
			},
			// Deleting a sensor is not allowed
			{
				PreConfig: func() {
					util.MockGetSensor("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("uid", "_2")}),
						1,
					)
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

func TestSensorResource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var request429 *gock.Response

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
					request429 = gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/sensors").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetSensor("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("uid", "")}),
						2,
					)
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
					resource.TestCheckResourceAttr("uxi_sensor.my_sensor", "id", "uid"),
					func(s *terraform.State) error {
						st.Assert(t, request429.Mock.Request().Counter, 0)
						return nil
					},
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

	mockOAuth.Mock.Disable()

}
func TestSensorResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Read 5xx error
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/sensors").
						Reply(500).
						JSON(map[string]interface{}{
							"httpStatusCode": 500,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
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

				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Read not found
			{
				PreConfig: func() {
					util.MockGetSensor(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
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

				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
