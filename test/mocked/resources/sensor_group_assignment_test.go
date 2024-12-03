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
	"github.com/stretchr/testify/assert"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
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
					util.MockGetSensor("id", util.GenerateSensorsGetResponse("sensor_id", ""), 2)

					// required for group create
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("group_id", "", ""),
						util.GenerateGroupPostResponse("group_id", "", ""),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)

					// required for sensor group assignment create
					util.MockPostSensorGroupAssignment(
						util.GenerateSensorGroupAssignmentPostRequest(
							"sensor_group_assignment_id",
							"",
						),
						util.GenerateSensorGroupAssignmentPostResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "hpeuxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "sensor_id"
					}

					resource "hpeuxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = hpeuxi_sensor.my_sensor.id
						group_id 		= hpeuxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"sensor_id",
						"sensor_id",
					),
					resource.TestCheckResourceAttr(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"group_id",
						"group_id",
					),
					resource.TestCheckResourceAttr(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"id",
						"sensor_group_assignment_id",
					),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)
				},
				ResourceName:      "hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					util.MockGetSensor(
						"sensor_id_2",
						util.GenerateSensorsGetResponse("sensor_id_2", "_2"),
						2,
					)
					util.MockGetSensor(
						"sensor_id",
						util.GenerateSensorsGetResponse("sensor_id", ""),
						3,
					)

					util.MockGetGroup(
						"group_id_2",
						util.GenerateGroupGetResponse("group_id_2", "_2", "_2"),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						2,
					)

					// required for creating another group
					util.MockDeleteSensorGroupAssignment("sensor_group_assignment_id", 1)
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("group_id_2", "_2", "_2"),
						util.GenerateGroupPostResponse("group_id_2", "_2", "_2"),
						1,
					)

					// required for sensor group assignment create
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id_2",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id_2",
							"_2",
						),
						1,
					)

					util.MockPostSensorGroupAssignment(
						util.GenerateSensorGroupAssignmentPostRequest(
							"sensor_group_assignment_id_2",
							"_2",
						),
						util.GenerateSensorGroupAssignmentPostResponse(
							"sensor_group_assignment_id_2",
							"_2",
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					// the original resources
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "hpeuxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "sensor_id"
					}

					// the new resources we wanna update the assignment to
					resource "hpeuxi_group" "my_group_2" {
						name            = "name_2"
						parent_group_id = "parent_id_2"
					}

					resource "hpeuxi_sensor" "my_sensor_2" {
						name 			= "name_2"
						address_note 	= "address_note_2"
						notes 			= "notes_2"
						pcap_mode 		= "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor_2
						id = "sensor_id_2"
					}

					// the assignment update, updated from sensor/group to sensor_2/group_2
					resource "hpeuxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = hpeuxi_sensor.my_sensor_2.id
						group_id 		= hpeuxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"sensor_id",
						"sensor_id_2",
					),
					resource.TestCheckResourceAttr(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"group_id",
						"group_id_2",
					),
					resource.TestCheckResourceAttr(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"id",
						"sensor_group_assignment_id_2",
					),
				),
			},
			// Delete sensor-group assignments and remove sensors from state
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_id_2",
						util.GenerateGroupGetResponse("group_id_2", "_2", "_2"),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)

					util.MockDeleteGroup("group_id", 1)
					util.MockDeleteGroup("group_id_2", 1)
					util.MockDeleteSensorGroupAssignment("sensor_group_assignment_id_2", 1)
				},
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = hpeuxi_sensor.my_sensor_2

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestSensorGroupAssignmentResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating
			{
				PreConfig: func() {
					// required for sensor import
					util.MockGetSensor(
						"sensor_id",
						util.GenerateSensorsGetResponse("sensor_id", ""),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("group_id", "", ""),
						util.GenerateGroupPostResponse("group_id", "", ""),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)

					// required for sensor group assignment create
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Post(shared.SensorGroupAssignmentPath).
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPostSensorGroupAssignment(
						util.GenerateSensorGroupAssignmentPostRequest(
							"sensor_group_assignment_id",
							"",
						),
						util.GenerateSensorGroupAssignmentPostResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "hpeuxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "sensor_id"
					}

					resource "hpeuxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = hpeuxi_sensor.my_sensor.id
						group_id 		= hpeuxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"id",
						"sensor_group_assignment_id",
					),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Read
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/sensor-group-assignments").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)
				},
				ResourceName:      "hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Delete
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)

					util.MockDeleteGroup("group_id", 1)
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Delete("/networking-uxi/v1alpha1/sensor-group-assignments/sensor_group_assignment_id").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockDeleteSensorGroupAssignment("sensor_group_assignment_id", 1)
				},
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = hpeuxi_sensor.my_sensor_2

						lifecycle {
							destroy = false
						}
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

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
					util.MockGetSensor(
						"sensor_id",
						util.GenerateSensorsGetResponse("sensor_id", ""),
						1,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("group_id", "", ""),
						util.GenerateGroupPostResponse("group_id", "", ""),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)

					// required for sensor group assignment create
					gock.New(util.MockUXIURL).
						Post(shared.SensorGroupAssignmentPath).
						Reply(http.StatusBadRequest).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusBadRequest,
							"errorCode":      "HPE_GL_ERROR_BAD_REQUEST",
							"message":        "Validation error - bad request",
							"debugId":        "12312-123123-123123-1231212",
						})
				},

				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "hpeuxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "sensor_id"
					}

					resource "hpeuxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = hpeuxi_sensor.my_sensor.id
						group_id 		= hpeuxi_group.my_group.id
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// read not found error
			{
				PreConfig: func() {
					// required for sensor import
					util.MockGetSensor(
						"sensor_id",
						util.GenerateSensorsGetResponse("sensor_id", ""),
						1,
					)

					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)

					// sensor group assignment import
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.EmptyGetListResponse,
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "hpeuxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "sensor_id"
					}

					resource "hpeuxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = hpeuxi_sensor.my_sensor.id
						group_id 		= hpeuxi_group.my_group.id
					}

					import {
						to = hpeuxi_sensor_group_assignment.my_sensor_group_assignment
						id = "sensor_group_assignment_id"
					}
				`,
				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Read HTTP error
			{
				PreConfig: func() {
					// required for sensor import
					util.MockGetSensor(
						"sensor_id",
						util.GenerateSensorsGetResponse("sensor_id", ""),
						1,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("group_id", "", ""),
						util.GenerateGroupPostResponse("group_id", "", ""),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)

					// required for sensor group assignment read
					gock.New(util.MockUXIURL).
						Get(shared.SensorGroupAssignmentPath).
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "hpeuxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "sensor_id"
					}

					resource "hpeuxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = hpeuxi_sensor.my_sensor.id
						group_id 		= hpeuxi_group.my_group.id
					}

					import {
						to = hpeuxi_sensor_group_assignment.my_sensor_group_assignment
						id = "sensor_group_assignment_id"
					}
				`,

				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Actually creating a sensor group assignment - needed for next step
			{
				PreConfig: func() {
					// required for sensor import
					util.MockGetSensor(
						"sensor_id",
						util.GenerateSensorsGetResponse("sensor_id", ""),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("group_id", "", ""),
						util.GenerateGroupPostResponse("group_id", "", ""),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)

					// required for sensor group assignment create
					util.MockPostSensorGroupAssignment(
						util.GenerateSensorGroupAssignmentPostRequest(
							"sensor_group_assignment_id",
							"",
						),
						util.GenerateSensorGroupAssignmentPostResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "hpeuxi_sensor" "my_sensor" {
						name 			= "name"
						address_note 	= "address_note"
						notes 			= "notes"
						pcap_mode 		= "light"
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "sensor_id"
					}

					resource "hpeuxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id       = hpeuxi_sensor.my_sensor.id
						group_id 		= hpeuxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"id",
						"sensor_group_assignment_id",
					),
				),
			},
			// Delete sensor-group assignments and remove sensors from state - errors
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)

					gock.New(util.MockUXIURL).
						Delete("/networking-uxi/v1alpha1/sensor-group-assignments/sensor_group_assignment_id").
						Reply(http.StatusBadRequest).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusBadRequest,
							"errorCode":      "HPE_GL_ERROR_BAD_REQUEST",
							"message":        "Validation error - bad request",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					removed {
						from = hpeuxi_sensor.my_sensor

						lifecycle {
							destroy = false
						}
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Actually delete sensor-group assignments and remove sensors from state
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_id",
						util.GenerateSensorGroupAssignmentsGetResponse(
							"sensor_group_assignment_id",
							"",
						),
						1,
					)
					util.MockDeleteGroup("group_id", 1)
					util.MockDeleteSensorGroupAssignment("sensor_group_assignment_id", 1)
				},
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
