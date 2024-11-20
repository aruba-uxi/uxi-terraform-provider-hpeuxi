/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"
)

func TestServiceTestGroupAssignmentResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a service test group assignment
			{
				PreConfig: func() {
					// required for service test import
					util.MockGetServiceTest(
						"service_test_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("service_test_id", ""),
							},
						),
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
					// required for service test group assignment create
					util.MockPostServiceTestGroupAssignment(
						util.GenerateServiceTestGroupAssignmentPostRequest(
							"service_test_group_assignment_id",
							"",
						),
						util.GenerateServiceTestGroupAssignmentPostResponse(
							"service_test_group_assignment_id",
							"",
						),
						1,
					)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_id", "",
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "service_test_id"
					}

					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						"service_test_id",
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						"group_id",
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						"service_test_group_assignment_id",
					),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_id", "",
						),
						1,
					)
				},
				ResourceName:      "uxi_service_test_group_assignment.my_service_test_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				PreConfig: func() {
					util.MockGetServiceTest(
						"service_test_id_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("service_test_id_2", "_2"),
							},
						),
						2,
					)
					util.MockGetServiceTest(
						"service_test_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("service_test_id", ""),
							},
						),
						2,
					)
					util.MockGetGroup(
						"group_id_2",
						util.GenerateGroupGetResponse("group_id_2", "_2", "_2"),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						3,
					)

					// required for deleting existing group
					util.MockDeleteServiceTestGroupAssignment("service_test_group_assignment", 1)

					// required for creating another group
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("group_id_2", "_2", "_2"),
						util.GenerateGroupPostResponse("group_id_2", "_2", "_2"),
						1,
					)

					// required for service test group assignment create
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id_2",
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_id_2", "_2",
						),
						2,
					)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_id", "",
						),
						1,
					)
					util.MockPostServiceTestGroupAssignment(
						util.GenerateServiceTestGroupAssignmentPostRequest(
							"service_test_group_assignment_id_2",
							"_2",
						),
						util.GenerateServiceTestGroupAssignmentPostResponse(
							"service_test_group_assignment_id_2",
							"_2",
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "service_test_id"
					}

					// the new resources we wanna update the assignment to
					resource "uxi_group" "my_group_2" {
						name            = "name_2"
						parent_group_id = "parent_id_2"
					}

					resource "uxi_service_test" "my_service_test_2" {
						name = "name_2"
					}

					import {
						to = uxi_service_test.my_service_test_2
						id = "service_test_id_2"
					}

					// the assignment update, updated from service_test/group to service_test_2/group_2
					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test_2.id
						group_id 		= uxi_group.my_group_2.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						"service_test_id_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						"group_id_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						"service_test_group_assignment_id_2",
					),
				),
			},
			// Delete service test group assignment and remove service tests from state
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)
					util.MockGetGroup(
						"group_id_2",
						util.GenerateGroupGetResponse("group_id_2", "_2", "_2"),
						1,
					)

					util.MockDeleteGroup("group_id", 1)
					util.MockDeleteGroup("group_id_2", 1)
					util.MockDeleteServiceTestGroupAssignment(
						"service_test_group_assignment_id_2",
						1,
					)
				},
				Config: provider.ProviderConfig + `
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
		},
	})

	mockOAuth.Mock.Disable()
}

func TestServiceTestGroupAssignmentResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					// required for service test import
					util.MockGetServiceTest(
						"service_test_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("service_test_id", ""),
							},
						),
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

					// required for service test group assignment create
					mockTooManyRequests = gock.New(util.MockUxiUrl).
						Post(shared.ServiceTestGroupAssignmentPath).
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPostServiceTestGroupAssignment(
						util.GenerateServiceTestGroupAssignmentPostRequest(
							"service_test_group_assignment_id",
							"",
						),
						util.GenerateServiceTestGroupAssignmentPostResponse(
							"service_test_group_assignment_id",
							"",
						),
						1,
					)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_id", "",
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "service_test_id"
					}

					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						"service_test_group_assignment_id",
					),
					func(s *terraform.State) error {
						assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Read
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/service-test-group-assignments").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_id", "",
						),
						1,
					)
				},
				ResourceName:      "uxi_service_test_group_assignment.my_service_test_group_assignment",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) error {
						assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
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
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_id", "",
						),
						1,
					)

					util.MockDeleteGroup("group_id", 1)
					util.MockDeleteServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						1,
					)
				},
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_service_test.my_service_test

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestServiceTestGroupAssignmentResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a service test group assignment - errors
			{
				PreConfig: func() {
					// required for service test import
					util.MockGetServiceTest(
						"service_test_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("service_test_id", ""),
							},
						),
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

					// required for service test group assignment create
					gock.New(util.MockUxiUrl).
						Post(shared.ServiceTestGroupAssignmentPath).
						Reply(http.StatusBadRequest).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusBadRequest,
							"errorCode":      "HPE_GL_ERROR_BAD_REQUEST",
							"message":        "Validation error - bad request",
							"debugId":        "12312-123123-123123-1231212",
						})
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "service_test_id"
					}

					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group.id
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Read not found
			{
				PreConfig: func() {
					util.MockGetServiceTest(
						"service_test_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("service_test_id", ""),
							},
						),
						1,
					)
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)

					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						util.EmptyGetListResponse,
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group.id
					}

					import {
						to = uxi_service_test_group_assignment.my_service_test_group_assignment
						id = "service_test_group_assignment_id"
					}`,
				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Actually creating a service test group assignment - needed for next step
			{
				PreConfig: func() {
					// required for service test import
					util.MockGetServiceTest(
						"service_test_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("service_test_id", ""),
							},
						),
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
					// required for service test group assignment create
					util.MockPostServiceTestGroupAssignment(
						util.GenerateServiceTestGroupAssignmentPostRequest(
							"service_test_group_assignment_id",
							"",
						),
						util.GenerateServiceTestGroupAssignmentPostResponse(
							"service_test_group_assignment_id",
							"",
						),
						1,
					)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_id", "",
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "service_test_id"
					}

					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group.id
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						"service_test_id",
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						"group_id",
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						"service_test_group_assignment_id",
					),
				),
			},
			// Read HTTP error
			{
				PreConfig: func() {
					// required for service test import
					util.MockGetServiceTest(
						"service_test_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("service_test_id", ""),
							},
						),
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

					// required for service test group assignment read
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/service-test-group-assignments").
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group.id
					}

					import {
						to = uxi_service_test_group_assignment.my_service_test_group_assignment
						id = "service_test_group_assignment_id"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Delete service test group assignment - errors
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_id",
							"",
						),
						1,
					)
					gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/service-test-group-assignments/service_test_group_assignment_id").
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
						from = uxi_service_test.my_service_test

						lifecycle {
							destroy = false
						}
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Actually delete service test group assignment and remove service tests from state
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_id",
						util.GenerateGroupGetResponse("group_id", "", ""),
						1,
					)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_id",
							"",
						),
						1,
					)

					util.MockDeleteGroup("group_id", 1)
					util.MockDeleteServiceTestGroupAssignment(
						"service_test_group_assignment_id",
						1,
					)
				},
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_service_test.my_service_test

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
