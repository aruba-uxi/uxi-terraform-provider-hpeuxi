/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
)

func TestSensorGroupAssignmentResource(t *testing.T) {
	const (
		groupName  = "tf_provider_acceptance_test_sensor_assignment_resource"
		group2Name = "tf_provider_acceptance_test_sensor_assignment_resource_two"
	)

	var (
		existingSensorProperties = util.GetSensor(config.SensorID)
		resourceIDBeforeRecreate string
		resourceIDAfterRecreate  string
	)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating
			{
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "hpeuxi_sensor" "my_sensor" {
						name 			= "` + existingSensorProperties.Name + `"
						` + util.ConditionalProperty("address_note", existingSensorProperties.AddressNote.Get()) + `
						` + util.ConditionalProperty("notes", existingSensorProperties.Notes.Get()) + `
						` + util.ConditionalProperty("pcap_mode", (*string)(existingSensorProperties.GetPcapMode().Ptr())) + `
					}

					import {
						to = hpeuxi_sensor.my_sensor
						id = "` + config.SensorID + `"
					}

					resource "hpeuxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id = hpeuxi_sensor.my_sensor.id
						group_id  = hpeuxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"sensor_id",
						config.SensorID,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"group_id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupName).Id)

							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "hpeuxi_sensor_group_assignment.my_sensor_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIDBeforeRecreate = rs.Primary.ID

						return util.CheckStateAgainstSensorGroupAssignment(
							t,
							"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
							util.GetSensorGroupAssignment(resourceIDBeforeRecreate),
						)(s)
					},
				),
			},
			// ImportState
			{
				ResourceName:      "hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: provider.ProviderConfig + `
					// the original resources
					resource "hpeuxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "hpeuxi_sensor" "my_sensor" {
						name 			= "` + existingSensorProperties.Name + `"
						` + util.ConditionalProperty("address_note", existingSensorProperties.AddressNote.Get()) + `
						` + util.ConditionalProperty("notes", existingSensorProperties.Notes.Get()) + `
						` + util.ConditionalProperty("pcap_mode", (*string)(existingSensorProperties.PcapMode.Get().Ptr())) + `
					}

					// the new resources we wanna update the assignment to
					resource "hpeuxi_group" "my_group_2" {
						name = "` + group2Name + `"
					}

					// the assignment update, updated from sensor/group to sensor/group_2
					resource "hpeuxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id = hpeuxi_sensor.my_sensor.id
						group_id  = hpeuxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check configured properties
					resource.TestCheckResourceAttr(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"sensor_id",
						config.SensorID,
					),
					resource.TestCheckResourceAttrWith(
						"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
						"group_id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(group2Name).Id)

							return nil
						},
					),
					// Check properties match what is on backend
					func(s *terraform.State) error {
						resourceName := "hpeuxi_sensor_group_assignment.my_sensor_group_assignment"
						rs := s.RootModule().Resources[resourceName]
						resourceIDAfterRecreate = rs.Primary.ID

						return util.CheckStateAgainstSensorGroupAssignment(
							t,
							"hpeuxi_sensor_group_assignment.my_sensor_group_assignment",
							util.GetSensorGroupAssignment(resourceIDAfterRecreate),
						)(s)
					},
				),
			},
			// Delete sensor-group assignments and remove sensors from state
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
		CheckDestroy: func(s *terraform.State) error {
			assert.Equal(t, util.GetGroupByName(groupName), nil)
			assert.Equal(t, util.GetGroupByName(group2Name), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIDBeforeRecreate), nil)
			assert.Equal(t, util.GetAgentGroupAssignment(resourceIDAfterRecreate), nil)

			return nil
		},
	})
}
