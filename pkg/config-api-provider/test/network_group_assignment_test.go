package test

import (
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNetworkGroupAssignmentResource(t *testing.T) {

	// Test Wired Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a network group assignment
			{
				PreConfig: func() {
					// required for network import
					resources.GetWiredNetwork = func(uid string) resources.WiredNetworkResponseModel {
						return resources.WiredNetworkResponseModel{
							Uid:                  uid,
							Alias:                "alias",
							DatetimeCreated:      "datetime_created",
							DatetimeUpdated:      "datetime_updated",
							IpVersion:            "ip_version",
							Security:             "security",
							DnsLookupDomain:      "dns_lookup_domain",
							DisableEdns:          false,
							UseDns64:             false,
							ExternalConnectivity: false,
							VlanId:               123,
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

					// required for network group assignment create
					networkGroupAssignmentResponse := resources.NetworkGroupAssignmentResponseModel{
						UID:        "network_group_assignment_uid",
						GroupUID:   "group_uid",
						NetworkUID: "network_uid",
					}
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return networkGroupAssignmentResponse
					}
					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						return networkGroupAssignmentResponse
					}
				},

				Config: providerConfig + `
					resource "uxi_group" "my_group" {
						name       = "name"
						parent_uid = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						alias = "alias"
					}

					import {
						to = uxi_wired_network.my_network
						id = "network_uid"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wired_network.my_network.id
						group_id   = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "network_id", "network_uid"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "group_id", "group_uid"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "id", "network_group_assignment_uid"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_network_group_assignment.my_network_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					resources.GetWiredNetwork = func(uid string) resources.WiredNetworkResponseModel {
						if uid == "network_uid" {
							return resources.WiredNetworkResponseModel{
								Uid:                  uid,
								Alias:                "alias",
								DatetimeCreated:      "datetime_created",
								DatetimeUpdated:      "datetime_updated",
								IpVersion:            "ip_version",
								Security:             "security",
								DnsLookupDomain:      "dns_lookup_domain",
								DisableEdns:          false,
								UseDns64:             false,
								ExternalConnectivity: false,
								VlanId:               123,
							}
						} else {
							return resources.WiredNetworkResponseModel{
								Uid:                  uid,
								Alias:                "alias_2",
								DatetimeCreated:      "datetime_created_2",
								DatetimeUpdated:      "datetime_updated_2",
								IpVersion:            "ip_version_2",
								Security:             "security_2",
								DnsLookupDomain:      "dns_lookup_domain_2",
								DisableEdns:          false,
								UseDns64:             false,
								ExternalConnectivity: false,
								VlanId:               123,
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

					// required for network group assignment create
					networkGroupAssignmentOriginal := resources.NetworkGroupAssignmentResponseModel{
						UID:        "network_group_assignment_uid",
						GroupUID:   "group_uid",
						NetworkUID: "network_uid",
					}
					networkGroupAssignmentUpdated := resources.NetworkGroupAssignmentResponseModel{
						UID:        "network_group_assignment_uid_2",
						GroupUID:   "group_uid_2",
						NetworkUID: "network_uid_2",
					}

					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						if uid == "network_group_assignment_uid" {
							return networkGroupAssignmentOriginal
						} else {
							return networkGroupAssignmentUpdated
						}
					}
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return networkGroupAssignmentUpdated
					}
				},
				Config: providerConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name       = "name"
						parent_uid = "parent_uid"
					}

					resource "uxi_wired_network" "my_network" {
						alias = "alias"
					}

					import {
						to = uxi_wired_network.my_network
						id = "network_uid"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name       = "name_2"
						parent_uid = "parent_uid_2"
					}

					resource "uxi_wired_network" "my_network_2" {
						alias = "alias_2"
					}

					import {
						to = uxi_wired_network.my_network_2
						id = "network_uid_2"
					}

					// the assignment update, updated from network/group to network_2/group_2
					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = uxi_wired_network.my_network_2.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "network_id", "network_uid_2"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "group_id", "group_uid_2"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "id", "network_group_assignment_uid_2"),
				),
			},
			// Remove networks from state
			{
				Config: providerConfig + `
					removed {
						from = uxi_wired_network.my_network

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = uxi_wired_network.my_network_2

						lifecycle {
							destroy = false
						}
					}`,
			},
			// Delete testing automatically occurs in TestCase
		},
	})

	// Test Wireless Network Group Assignment
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a network group assignment
			{
				PreConfig: func() {
					// required for network import
					resources.GetWirelessNetwork = func(uid string) resources.WirelessNetworkResponseModel {
						return resources.WirelessNetworkResponseModel{
							Uid:                  uid,
							Ssid:                 "ssid",
							DatetimeCreated:      "datetime_created",
							DatetimeUpdated:      "datetime_updated",
							Alias:                "alias",
							IpVersion:            "ip_version",
							Security:             "security",
							Hidden:               false,
							BandLocking:          "band_locking",
							DnsLookupDomain:      "dns_lookup_domain",
							DisableEdns:          false,
							UseDns64:             false,
							ExternalConnectivity: false,
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

					// required for network group assignment create
					networkGroupAssignmentResponse := resources.NetworkGroupAssignmentResponseModel{
						UID:        "network_group_assignment_uid",
						GroupUID:   "group_uid",
						NetworkUID: "network_uid",
					}
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return networkGroupAssignmentResponse
					}
					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						return networkGroupAssignmentResponse
					}
				},

				Config: providerConfig + `
					resource "uxi_group" "my_group" {
						name       = "name"
						parent_uid = "parent_uid"
					}

					resource "uxi_wireless_network" "my_network" {
						alias = "alias"
					}

					import {
						to = uxi_wireless_network.my_network
						id = "network_uid"
					}

					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id = uxi_wireless_network.my_network.id
						group_id   = uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "network_id", "network_uid"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "group_id", "group_uid"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "id", "network_group_assignment_uid"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_network_group_assignment.my_network_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					resources.GetWirelessNetwork = func(uid string) resources.WirelessNetworkResponseModel {
						if uid == "network_uid" {
							return resources.WirelessNetworkResponseModel{
								Uid:                  uid,
								Ssid:                 "ssid",
								DatetimeCreated:      "datetime_created",
								DatetimeUpdated:      "datetime_updated",
								Alias:                "alias",
								IpVersion:            "ip_version",
								Security:             "security",
								Hidden:               false,
								BandLocking:          "band_locking",
								DnsLookupDomain:      "dns_lookup_domain",
								DisableEdns:          false,
								UseDns64:             false,
								ExternalConnectivity: false,
							}
						} else {
							return resources.WirelessNetworkResponseModel{
								Uid:                  uid,
								Ssid:                 "ssid_2",
								DatetimeCreated:      "datetime_created_2",
								DatetimeUpdated:      "datetime_updated_2",
								Alias:                "alias_2",
								IpVersion:            "ip_version_2",
								Security:             "security_2",
								Hidden:               false,
								BandLocking:          "band_locking_2",
								DnsLookupDomain:      "dns_lookup_domain_2",
								DisableEdns:          false,
								UseDns64:             false,
								ExternalConnectivity: false,
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

					// required for network group assignment create
					networkGroupAssignmentOriginal := resources.NetworkGroupAssignmentResponseModel{
						UID:        "network_group_assignment_uid",
						GroupUID:   "group_uid",
						NetworkUID: "network_uid",
					}
					networkGroupAssignmentUpdated := resources.NetworkGroupAssignmentResponseModel{
						UID:        "network_group_assignment_uid_2",
						GroupUID:   "group_uid_2",
						NetworkUID: "network_uid_2",
					}

					resources.GetNetworkGroupAssignment = func(uid string) resources.NetworkGroupAssignmentResponseModel {
						if uid == "network_group_assignment_uid" {
							return networkGroupAssignmentOriginal
						} else {
							return networkGroupAssignmentUpdated
						}
					}
					resources.CreateNetworkGroupAssignment = func(request resources.NetworkGroupAssignmentRequestModel) resources.NetworkGroupAssignmentResponseModel {
						return networkGroupAssignmentUpdated
					}
				},
				Config: providerConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name       = "name"
						parent_uid = "parent_uid"
					}

					resource "uxi_wireless_network" "my_network" {
						alias = "alias"
					}

					import {
						to = uxi_wireless_network.my_network
						id = "network_uid"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name       = "name_2"
						parent_uid = "parent_uid_2"
					}

					resource "uxi_wireless_network" "my_network_2" {
						alias = "alias_2"
					}

					import {
						to = uxi_wireless_network.my_network_2
						id = "network_uid_2"
					}

					// the assignment update, updated from network/group to network_2/group_2
					resource "uxi_network_group_assignment" "my_network_group_assignment" {
						network_id       = uxi_wireless_network.my_network_2.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "network_id", "network_uid_2"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "group_id", "group_uid_2"),
					resource.TestCheckResourceAttr("uxi_network_group_assignment.my_network_group_assignment", "id", "network_group_assignment_uid_2"),
				),
			},
			// Remove networks from state
			{
				Config: providerConfig + `
					removed {
						from = uxi_wireless_network.my_network

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = uxi_wireless_network.my_network_2

						lifecycle {
							destroy = false
						}
					}`,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
