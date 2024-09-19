package test

import (
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
					resources.GetSensor = func(uid string) resources.SensorResponseModel {
						return util.GenerateSensorResponseModel(uid, "")
					}

					// required for group create
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")), 1)
					util.MockGetGroup("group_uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						}),
						1,
					)

					// required for sensor group assignment create
					resources.CreateSensorGroupAssignment = func(request resources.SensorGroupAssignmentRequestModel) resources.SensorGroupAssignmentResponseModel {
						return util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")
					}
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GenerateSensorGroupAssignmentPaginatedResponse([]map[string]interface{}{
							util.StructToMap(util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")),
						}),
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
						util.GenerateSensorGroupAssignmentPaginatedResponse([]map[string]interface{}{
							util.StructToMap(util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")),
						}),
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
					resources.GetSensor = func(uid string) resources.SensorResponseModel {
						if uid == "sensor_uid" {
							return util.GenerateSensorResponseModel("sensor_uid", "")
						} else {
							return util.GenerateSensorResponseModel("sensor_uid", "_2")
						}
					}
					util.MockGetGroup("group_uid_2", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")),
						}),
						1,
					)
					util.MockGetGroup("group_uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						}),
						2,
					)

					// required for creating another group
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")), 1)

					// required for sensor group assignment create
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GenerateSensorGroupAssignmentPaginatedResponse([]map[string]interface{}{
							util.StructToMap(util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")),
						}),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid_2",
						util.GenerateSensorGroupAssignmentPaginatedResponse([]map[string]interface{}{
							util.StructToMap(util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid_2", "_2")),
						}),
						1,
					)

					resources.CreateSensorGroupAssignment = func(request resources.SensorGroupAssignmentRequestModel) resources.SensorGroupAssignmentResponseModel {
						return util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid_2", "_2")
					}
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
			// Remove sensors from state
			{
				PreConfig: func() {
					util.MockGetGroup("group_uid_2", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid_2", "_2", "_2")),
						}),
						1,
					)
					util.MockGetGroup("group_uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("group_uid", "", "")),
						}),
						1,
					)
					util.MockGetSensorGroupAssignment(
						"sensor_group_assignment_uid",
						util.GenerateSensorGroupAssignmentPaginatedResponse([]map[string]interface{}{
							util.StructToMap(util.GenerateSensorGroupAssignmentResponse("sensor_group_assignment_uid", "")),
						}),
						1,
					)
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
			// Delete testing automatically occurs in TestCase
		},
	})

	mockOAuth.Mock.Disable()
}
