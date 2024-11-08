package resource_test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-configuration/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-configuration/test/mocked/util"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/nbio/st"
)

func TestServiceTestGroupAssignmentResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a serviceTest group assignment
			{
				PreConfig: func() {
					// required for serviceTest import
					util.MockGetServiceTest(
						"service_test_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponseModel("service_test_uid", ""),
							},
						),
						2,
					)
					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						1,
					)
					// required for serviceTest group assignment create
					util.MockPostServiceTestGroupAssignment(
						util.GenerateServiceTestGroupAssignmentRequest(
							"service_test_group_assignment_uid",
							"",
						),
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_uid",
							"",
						),
						1,
					)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{
							util.GenerateServiceTestGroupAssignmentResponse(
								"service_test_group_assignment_uid", "",
							)},
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "name"
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
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						"service_test_uid",
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						"group_uid",
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						"service_test_group_assignment_uid",
					),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{
							util.GenerateServiceTestGroupAssignmentResponse(
								"service_test_group_assignment_uid", "",
							)},
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
						"service_test_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponseModel("service_test_uid_2", "_2"),
							},
						),
						2,
					)
					util.MockGetServiceTest(
						"service_test_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponseModel("service_test_uid", ""),
							},
						),
						2,
					)
					util.MockGetGroup(
						"group_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
							},
						),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						3,
					)

					// required for deleting existing group
					util.MockDeleteServiceTestGroupAssignment("service_test_group_assignment", 1)

					// required for creating another group
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid_2", "_2", "_2"),
						util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
						1,
					)

					// required for serviceTest group assignment create
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_uid_2",
						util.GeneratePaginatedResponse([]map[string]interface{}{
							util.GenerateServiceTestGroupAssignmentResponse(
								"service_test_group_assignment_uid_2", "_2",
							)},
						),
						2,
					)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{
							util.GenerateServiceTestGroupAssignmentResponse(
								"service_test_group_assignment_uid", "",
							)},
						),
						1,
					)
					util.MockPostServiceTestGroupAssignment(
						util.GenerateServiceTestGroupAssignmentRequest(
							"service_test_group_assignment_uid_2",
							"_2",
						),
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_uid_2",
							"_2",
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					// the original resources
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "name"
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
						name = "name_2"
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
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						"service_test_uid_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						"group_uid_2",
					),
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						"service_test_group_assignment_uid_2",
					),
				),
			},
			// Remove serviceTests from state
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						1,
					)
					util.MockGetGroup(
						"group_uid_2",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid_2", "_2", "_2"),
							},
						),
						1,
					)

					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteGroup("group_uid_2", 1)
					util.MockDeleteServiceTestGroupAssignment(
						"service_test_group_assignment_uid_2",
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

func TestServiceTestGroupAssignmentResource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mock429 *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a serviceTest group assignment
			{
				PreConfig: func() {
					// required for serviceTest import
					util.MockGetServiceTest(
						"service_test_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponseModel("service_test_uid", ""),
							},
						),
						2,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						1,
					)

					// required for serviceTest group assignment create
					mock429 = gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/service-test-group-assignments").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)

					util.MockPostServiceTestGroupAssignment(
						util.GenerateServiceTestGroupAssignmentRequest(
							"service_test_group_assignment_uid",
							"",
						),
						util.GenerateServiceTestGroupAssignmentResponse(
							"service_test_group_assignment_uid",
							"",
						),
						1,
					)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{
							util.GenerateServiceTestGroupAssignmentResponse(
								"service_test_group_assignment_uid", "",
							)},
						),
						1,
					)
				},

				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					resource "uxi_service_test" "my_service_test" {
						name = "name"
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
					resource.TestCheckResourceAttr(
						"uxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						"service_test_group_assignment_uid",
					),
					func(s *terraform.State) error {
						st.Assert(t, mock429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Remove serviceTests from state
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						1,
					)
					util.MockGetServiceTestGroupAssignment(
						"service_test_group_assignment_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{
							util.GenerateServiceTestGroupAssignmentResponse(
								"service_test_group_assignment_uid", "",
							)},
						),
						1,
					)

					util.MockDeleteGroup("group_uid", 1)
					util.MockDeleteServiceTestGroupAssignment(
						"service_test_group_assignment_uid",
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
			// Creating a serviceTest group assignment - errors
			{
				PreConfig: func() {
					// required for serviceTest import
					util.MockGetServiceTest(
						"service_test_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponseModel("service_test_uid", ""),
							},
						),
						1,
					)

					// required for group create
					util.MockPostGroup(
						util.GenerateGroupRequestModel("group_uid", "", ""),
						util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
						1,
					)
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						1,
					)

					// required for serviceTest group assignment create
					gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/service-test-group-assignments").
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

					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "service_test_uid"
					}

					resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						service_test_id = uxi_service_test.my_service_test.id
						group_id 		= uxi_group.my_group.id
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Remove serviceTests from state
			{
				PreConfig: func() {
					util.MockGetGroup(
						"group_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("group_uid", "", ""),
							},
						),
						1,
					)

					util.MockDeleteGroup("group_uid", 1)
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
