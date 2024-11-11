package resource_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
)

func TestSensorGroupAssignmentResource(t *testing.T) {
	const groupName = "tf_provider_acceptance_test_sensor_association_test"
	const group2Name = "tf_provider_acceptance_test_sensor_association_test_two"
	existingSensorProperties := util.GetSensorProperties(config.SensorUid)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a sensor group assignment
			{
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "` + existingSensorProperties.Name + `"
						` + util.ConditionalProperty("address_note", existingSensorProperties.AddressNote) + `
						` + util.ConditionalProperty("notes", existingSensorProperties.Notes) + `
						` + util.ConditionalProperty("pcap_mode", existingSensorProperties.PcapMode) + `
					}

					import {
						to = uxi_sensor.my_sensor
						id = "` + config.SensorUid + `"
					}

					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id = uxi_sensor.my_sensor.id
						group_id  = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_sensor_group_assignment.my_sensor_group_assignment",
						"sensor_id",
						config.SensorUid,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_sensor_group_assignment.my_sensor_group_assignment",
						"group_id",
						func(value string) error {
							st.Assert(t, value, util.GetGroupByName(groupName).Id)
							return nil
						},
					),
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
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name = "` + groupName + `"
					}

					resource "uxi_sensor" "my_sensor" {
						name 			= "` + existingSensorProperties.Name + `"
						` + util.ConditionalProperty("address_note", existingSensorProperties.AddressNote) + `
						` + util.ConditionalProperty("notes", existingSensorProperties.Notes) + `
						` + util.ConditionalProperty("pcap_mode", existingSensorProperties.PcapMode) + `
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name = "` + group2Name + `"
					}

					// the assignment update, updated from sensor/group to sensor/group_2
					resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
						sensor_id = uxi_sensor.my_sensor.id
						group_id  = uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_sensor_group_assignment.my_sensor_group_assignment",
						"sensor_id",
						config.SensorUid,
					),
					resource.TestCheckResourceAttrWith(
						"uxi_sensor_group_assignment.my_sensor_group_assignment",
						"group_id",
						func(value string) error {
							st.Assert(t, value, util.GetGroupByName(group2Name).Id)
							return nil
						},
					),
				),
			},
			// Delete sensor-group assignments and remove sensors from state
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
