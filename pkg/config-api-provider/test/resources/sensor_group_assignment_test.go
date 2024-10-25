package resource_test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/nbio/st"
)

func TestSensorGroupAssignmentResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a sensor group assignment
			{
				PreConfig: func() {
					// required for sensor import
					util.MockGetSensor("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("sensor_uid", "")}),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						1,
					)

					// required for sensor group assignment create
					util.MockPostSensorGroupAssignment(
						util.GenerateSensorGroupAssignmentRequest("sensor_group_assignment_uid", ""),
						util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", ""),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")}),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "sensor_uid"
					}

					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = uxi_sensor.my_sensor.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor_group_assignment.my_sensor_group_assignment", "sensor_id", "sensor_uid"),
					resource.TestCheckResourceAttr("uxi_sensor_group_assignment.my_sensor_group_assignment", "group_id", "group_uid"),
					resource.TestCheckResourceAttr("uxi_sensor_group_assignment.my_sensor_group_assignment", "id", "sensor_group_assignment_uid"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")}),
						1,
					)
				},
				ResourceName:      "uxi_sensor_group_assignment.my_sensor_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					util.MockGetSensor("sensor_uid_2", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("sensor_uid_2", "_2")}),
						2,
					)
					util.MockGetSensor("sensor_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("sensor_uid", "")}),
						3,
					)

					util.MockGetGroup(
						"group_uid_2",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")}),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						2,
					)

					// required for creating another group
					util.MockDeleteSensorGroupAssignment("sensor_group_assignment_uid", 1)
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid_2", "_2", "_2"),
						util.StructToMap(util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")),
						1,
					)

					// required for sensor group assignment create
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")}),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid_2",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid_2", "_2")}),
						1,
					)

					util.MockPostSensorGroupAssignment(
						util.GenerateSensorGroupAssignmentRequest("sensor_group_assignment_uid_2", "_2"),
						util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid_2", "_2"),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "sensor_uid"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "name_2"
						parent_group_id = "parent_uid_2"
					}

					resource "uxi_sensor" "my_sensor_2" {
						name 			= "name_2"
						address_note 	= "address_note_2"
						notes 			= "notes_2"
						pcap_mode 		= "light_2"
					}

					import {
						to = uxi_sensor.my_sensor_2
						id = "sensor_uid_2"
					}

					// the assignment update, updated from sensor/group to sensor_2/group_2
					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = uxi_sensor.my_sensor_2.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor_group_assignment.my_sensor_group_assignment", "sensor_id", "sensor_uid_2"),
					resource.TestCheckResourceAttr("uxi_sensor_group_assignment.my_sensor_group_assignment", "group_id", "group_uid_2"),
					resource.TestCheckResourceAttr("uxi_sensor_group_assignment.my_sensor_group_assignment", "id", "sensor_group_assignment_uid_2"),
				),
			},
			// Delete sensor-group assignments and remove sensors from state
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_uid_2",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")}),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")}),
						1,
					)

					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteGroup("group_uid_2", 1)
					util.MockDeleteSensorGroupAssignment("sensor_group_assignment_uid_2", 1)
				},
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = uxi_sensor.my_sensor_2

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestSensorGroupAssignmentResource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mock429 *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a sensor group assignment
			{
				PreConfig: func() {
					// required for sensor import
					util.MockGetSensor("sensor_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("sensor_uid", "")}),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						1,
					)

					// required for sensor group assignment create
					mock429 = gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/sensor-group-assignments").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPostSensorGroupAssignment(
						util.GenerateSensorGroupAssignmentRequest("sensor_group_assignment_uid", ""),
						util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", ""),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")}),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "sensor_uid"
					}

					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = uxi_sensor.my_sensor.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor_group_assignment.my_sensor_group_assignment", "id", "sensor_group_assignment_uid"),
					func(s *terraform.State) error {
						st.Assert(t, mock429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Delete sensor-group assignments and remove sensors from state
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")}),
						1,
					)

					util.MockDeleteGroup("group_uid", 1)
					mock429 = gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/sensor-group-assignments/sensor_group_assignment_uid").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockDeleteSensorGroupAssignment("sensor_group_assignment_uid", 1)
				},
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = uxi_sensor.my_sensor_2

						lifecycle {
							destroy = false
						}
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) error {
						st.Assert(t, mock429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestSensorGroupAssignmentResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a sensor group assignment - errors
			{
				PreConfig: func() {
					// required for sensor import
					util.MockGetSensor("sensor_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("sensor_uid", "")}),
						1,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						1,
					)

					// required for sensor group assignment create
					gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/sensor-group-assignments").
						Reply(400).
						JSON(map[string]interface{}{
							"httpStatusCode": 400,
							"errorCode":      "HPE_GL_ERROR_BAD_REQUEST",
							"message":        "Validation error - bad request",
							"debugId":        "12312-123123-123123-1231212",
						})
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "sensor_uid"
					}

					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = uxi_sensor.my_sensor.id
						group_id 		= uxi_group.my_group.id
					}`,
				ExpectError: regexp.MustCompile(`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`),
			},
			// read not found error
			{
				PreConfig: func() {
					// required for sensor import
					util.MockGetSensor("sensor_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("sensor_uid", "")}),
						1,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						1,
					)

					// sensor group assignment create
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "sensor_uid"
					}

					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = uxi_sensor.my_sensor.id
						group_id 		= uxi_group.my_group.id
					}

					import {
						to = uxi_sensor_group_assignment.my_sensor_group_assignment
						id = "sensor_group_assignment_uid"
					}
				`,
				ExpectError: regexp.MustCompile(`Could not find specified resource`),
			},
			// Read 5xx error
			{
				PreConfig: func() {
					// required for sensor import
					util.MockGetSensor("sensor_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("sensor_uid", "")}),
						1,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						1,
					)

					// required for sensor group assignment read
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/sensor-group-assignments").
						Reply(500).
						JSON(map[string]interface{}{
							"httpStatusCode": 500,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "sensor_uid"
					}

					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = uxi_sensor.my_sensor.id
						group_id 		= uxi_group.my_group.id
					}

					import {
						to = uxi_sensor_group_assignment.my_sensor_group_assignment
						id = "sensor_group_assignment_uid"
					}
				`,

				ExpectError: regexp.MustCompile(`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`),
			},
			// Actually Creating a sensor group assignment - needed for next step
			{
				PreConfig: func() {
					// required for sensor import
					util.MockGetSensor("sensor_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.GenerateSensorResponseModel("sensor_uid", "")}),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						1,
					)

					// required for sensor group assignment create
					util.MockPostSensorGroupAssignment(
						util.GenerateSensorGroupAssignmentRequest("sensor_group_assignment_uid", ""),
						util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", ""),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")}),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = uxi_sensor.my_sensor
						id = "sensor_uid"
					}

					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = uxi_sensor.my_sensor.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_sensor_group_assignment.my_sensor_group_assignment", "id", "sensor_group_assignment_uid"),
				),
			},
			// Delete sensor-group assignments and remove sensors from state - errors
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")}),
						1,
					)

					gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/sensor-group-assignments/sensor_group_assignment_uid").
						Reply(400).
						JSON(map[string]interface{}{
							"httpStatusCode": 400,
							"errorCode":      "HPE_GL_ERROR_BAD_REQUEST",
							"message":        "Validation error - bad request",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = uxi_sensor.my_sensor_2

						lifecycle {
							destroy = false
						}
					}`,
				ExpectError: regexp.MustCompile(`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`),
			},
			// Actually Delete sensor-group assignments and remove sensors from state
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateGroupResponseModel("group_uid", "", "")}),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")}),
						1,
					)
					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteSensorGroupAssignment("sensor_group_assignment_uid", 1)
				},
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = uxi_sensor.my_sensor_2

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
