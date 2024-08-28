package test

import (
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSensorGroupAssignmentResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a sensor group assignment
			{
				PreConfig: func() {
					// required for sensor import
					resources.GetSensor = func(uid string) resources.SensorResponseModel {
						return resources.SensorResponseModel{
							UID:                uid,
							Serial:             "serial",
							Name:               "name",
							ModelNumber:        "model_number",
							WifiMacAddress:     "wifi_mac_address",
							EthernetMacAddress: "ethernet_mac_address",
							AddressNote:        "address_note",
							Longitude:          "longitude",
							Latitude:           "latitude",
							Notes:              "notes",
							PCapMode:           "light",
						}
					}

					// required for group create
					groupResponse := resources.GroupResponseModel{
						UID:       "group_uid",
						Name:      "name",
						ParentUid: "parent_uid",
						Path:      "parent_uid.group_uid",
					}
					resources.CreateGroup = func(request resources.GroupCreateRequestModel) resources.GroupResponseModel {
						return groupResponse
					}
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						return groupResponse
					}

					// required for sensor group assignment create
					sensorGroupAssignmentResponse := resources.SensorGroupAssignmentResponseModel{
						UID:       "sensor_group_assignment_uid",
						GroupUID:  "group_uid",
						SensorUID: "sensor_uid",
					}
					resources.CreateSensorGroupAssignment = func(request resources.SensorGroupAssignmentRequestModel) resources.SensorGroupAssignmentResponseModel {
						return sensorGroupAssignmentResponse
					}
					resources.GetSensorGroupAssignment = func(uid string) resources.SensorGroupAssignmentResponseModel {
						return sensorGroupAssignmentResponse
					}
				},

				Config: providerConfig + `
					resource "uxi_group" "my_group" {
						name       = "name"
						parent_uid = "parent_uid"
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
				ResourceName:      "uxi_sensor_group_assignment.my_sensor_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					resources.GetSensor = func(uid string) resources.SensorResponseModel {
						if uid == "sensor_uid" {
							return resources.SensorResponseModel{
								UID:                "sensor_uid",
								Serial:             "serial",
								Name:               "name",
								ModelNumber:        "model_number",
								WifiMacAddress:     "wifi_mac_address",
								EthernetMacAddress: "ethernet_mac_address",
								AddressNote:        "address_note",
								Longitude:          "longitude",
								Latitude:           "latitude",
								Notes:              "notes",
								PCapMode:           "light",
							}
						} else {
							return resources.SensorResponseModel{
								UID:                "sensor_uid_2",
								Serial:             "serial_2",
								Name:               "name_2",
								ModelNumber:        "model_number_2",
								WifiMacAddress:     "wifi_mac_address_2",
								EthernetMacAddress: "ethernet_mac_address_2",
								AddressNote:        "address_note_2",
								Longitude:          "longitude_2",
								Latitude:           "latitude_2",
								Notes:              "notes_2",
								PCapMode:           "light",
							}
						}
					}

					// required for creating another group
					newGroupResponse := resources.GroupResponseModel{
						UID:       "group_uid_2",
						Name:      "name_2",
						ParentUid: "parent_uid_2",
						Path:      "parent_uid_2.group_uid_2",
					}
					resources.CreateGroup = func(request resources.GroupCreateRequestModel) resources.GroupResponseModel {
						return newGroupResponse
					}
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						if uid == "group_uid" {
							return resources.GroupResponseModel{
								UID:       uid,
								Name:      "name",
								ParentUid: "parent_uid",
								Path:      "parent_uid.group_uid",
							}
						} else {
							return newGroupResponse
						}
					}

					// required for sensor group assignment create
					sensorGroupAssignmentOriginal := resources.SensorGroupAssignmentResponseModel{
						UID:       "sensor_group_assignment_uid",
						GroupUID:  "group_uid",
						SensorUID: "sensor_uid",
					}
					sensorGroupAssignmentUpdated := resources.SensorGroupAssignmentResponseModel{
						UID:       "sensor_group_assignment_uid_2",
						GroupUID:  "group_uid_2",
						SensorUID: "sensor_uid_2",
					}

					resources.GetSensorGroupAssignment = func(uid string) resources.SensorGroupAssignmentResponseModel {
						if uid == "sensor_group_assignment_uid" {
							return sensorGroupAssignmentOriginal
						} else {
							return sensorGroupAssignmentUpdated
						}
					}
					resources.CreateSensorGroupAssignment = func(request resources.SensorGroupAssignmentRequestModel) resources.SensorGroupAssignmentResponseModel {
						return sensorGroupAssignmentUpdated
					}
				},
				Config: providerConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name       = "name"
						parent_uid = "parent_uid"
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
						name       = "name_2"
						parent_uid = "parent_uid_2"
					}

					resource "uxi_sensor" "my_sensor_2" {
						name 			= "name_2"
						address_note 	= "address_note_2"
						notes 			= "notes_2"
						pcap_mode 		= "light"
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
				Config: providerConfig + `
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
}
