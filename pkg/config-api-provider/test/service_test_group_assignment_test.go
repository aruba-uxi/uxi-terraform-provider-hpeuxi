package test

import (
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestServiceTestGroupAssignmentResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a serviceTest group assignment
			{
				PreConfig: func() {
					MockOAuth()
					// required for serviceTest import
					resources.GetServiceTest = func(uid string) resources.ServiceTestResponseModel {
						return GenerateServiceTestResponseModel(uid, "")
					}

					// required for group create
					MockPostGroup(StructToMap(GenerateGroupResponseModel("group_uid", "", "")))
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						return GenerateGroupResponseModel("group_uid", "", "")
					}

					// required for serviceTest group assignment create
					serviceTestGroupAssignmentResponse := GenerateServiceTestGroupAssignmentResponse("service_test_group_assignment_uid", "")
					resources.CreateServiceTestGroupAssignment = func(request resources.ServiceTestGroupAssignmentRequestModel) resources.ServiceTestGroupAssignmentResponseModel {
						return serviceTestGroupAssignmentResponse
					}
					resources.GetServiceTestGroupAssignment = func(uid string) resources.ServiceTestGroupAssignmentResponseModel {
						return serviceTestGroupAssignmentResponse
					}
				},

				Config: providerConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_service_test" "my_service_test" {
						title = "title"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "service_test_uid"
					}

					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_service_test_group_assignment.my_service_test_group_assignment", "service_test_id", "service_test_uid"),
					resource.TestCheckResourceAttr("uxi_service_test_group_assignment.my_service_test_group_assignment", "group_id", "group_uid"),
					resource.TestCheckResourceAttr("uxi_service_test_group_assignment.my_service_test_group_assignment", "id", "service_test_group_assignment_uid"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "uxi_service_test_group_assignment.my_service_test_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					MockOAuth()
					resources.GetServiceTest = func(uid string) resources.ServiceTestResponseModel {
						if uid == "service_test_uid" {
							return GenerateServiceTestResponseModel("service_test_uid", "")
						} else {
							return GenerateServiceTestResponseModel("service_test_uid", "_2")
						}
					}

					// required for creating another group
					MockPostGroup(StructToMap(GenerateGroupResponseModel("group_uid_2", "_2", "_2")))
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						if uid == "group_uid" {
							return GenerateGroupResponseModel(uid, "", "")
						} else {
							return GenerateGroupResponseModel("group_uid_2", "_2", "_2")
						}
					}

					// required for serviceTest group assignment create
					resources.GetServiceTestGroupAssignment = func(uid string) resources.ServiceTestGroupAssignmentResponseModel {
						if uid == "service_test_group_assignment_uid" {
							return GenerateServiceTestGroupAssignmentResponse("service_test_group_assignment_uid", "")
						} else {
							return GenerateServiceTestGroupAssignmentResponse("service_test_group_assignment_uid_2", "_2")
						}
					}
					resources.CreateServiceTestGroupAssignment = func(request resources.ServiceTestGroupAssignmentRequestModel) resources.ServiceTestGroupAssignmentResponseModel {
						return GenerateServiceTestGroupAssignmentResponse("service_test_group_assignment_uid_2", "_2")
					}
				},
				Config: providerConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_service_test" "my_service_test" {
						title = "title"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "service_test_uid"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "name_2"
						parent_group_id = "parent_uid_2"
					}

					resource "uxi_service_test" "my_service_test_2" {
						title = "title_2"
					}

					import {
						to = uxi_service_test.my_service_test_2
						id = "service_test_uid_2"
					}

					// the assignment update, updated from service_test/group to service_test_2/group_2
					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test_2.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_service_test_group_assignment.my_service_test_group_assignment", "service_test_id", "service_test_uid_2"),
					resource.TestCheckResourceAttr("uxi_service_test_group_assignment.my_service_test_group_assignment", "group_id", "group_uid_2"),
					resource.TestCheckResourceAttr("uxi_service_test_group_assignment.my_service_test_group_assignment", "id", "service_test_group_assignment_uid_2"),
				),
			},
			// Remove serviceTests from state
			{
				Config: providerConfig + `
					removed {
						from = uxi_service_test.my_service_test

						lifecycle {
							destroy = false
						}
					}

					removed {
						from = uxi_service_test.my_service_test_2

						lifecycle {
							destroy = false
						}
					}`,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
