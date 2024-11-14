package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/nbio/st"
)

type Fetcher interface {
	FetchData() ([]byte, error)
}

func TestGroupResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				PreConfig: func() {
					util.MockPostGroup(
						util.GenerateGroupRequestModel("id", "", ""),
						util.StructToMap(util.GenerateNonRootGroupResponseModel("id", "", "")),
						1,
					)
					util.MockGetGroup("id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("id", "", ""),
							},
						),
						2,
					)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_id", "", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr(
						"uxi_group.my_group",
						"parent_group_id",
						"parent_id",
					),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "id"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetGroup(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("id", "", ""),
							},
						),
						1,
					)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_id", "", ""),
							},
						),
						1,
					)
				},
				ResourceName:      "uxi_group.my_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update that does not trigger a recreate
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("id", "", ""),
						},
					),
						1,
					)
					// updated group
					util.MockUpdateGroup(
						"id",
						map[string]interface{}{"name": "name_2"},
						util.GenerateNonRootGroupResponseModel("id", "_2", ""),
						1,
					)
					util.MockGetGroup("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("id", "_2", ""),
						},
					),
						3,
					)
					util.MockGetGroup(
						"parent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_id", "", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_id"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name_2"),
					resource.TestCheckResourceAttr(
						"uxi_group.my_group",
						"parent_group_id",
						"parent_id",
					),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "id"),
				),
			},
			// Update that does trigger a recreate
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("id", "", ""),
						},
					),
						1,
					)
					// new group (replacement)
					util.MockPostGroup(
						util.GenerateGroupRequestModel("new_id", "", "_2"),
						util.GenerateNonRootGroupResponseModel("new_id", "", "_2"),
						1,
					)
					util.MockGetGroup("new_id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("new_id", "", "_2"),
						},
					),
						1,
					)
					// delete old group (being replaced)
					util.MockDeleteGroup("id", 1)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id_2"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr(
						"uxi_group.my_group",
						"parent_group_id",
						"parent_id_2",
					),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "new_id"),
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					util.MockGetGroup("new_id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("new_id", "", "_2"),
						},
					),
						1,
					)
					util.MockDeleteGroup("new_id", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestRootGroupResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Importing the root group does not work
			{
				PreConfig: func() {
					util.MockGetGroup(
						"my_root_group_id",
						util.GeneratePaginatedResponse([]map[string]interface{}{
							{
								"id":     "my_root_group_id",
								"name":   "root",
								"parent": *config_api_client.NewNullableParent(nil),
								"path":   "my_root_group_id",
								"type":   "networking-uxi/group",
							},
						}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_group" "my_root_group" {
					name            = "name"
				}

				import {
					to = uxi_group.my_root_group
					id = "my_root_group_id"
				}`,
				ExpectError: regexp.MustCompile(`The root group cannot be used as a resource`),
			},
			// Creating a group attached to the root
			{
				PreConfig: func() {
					util.MockPostGroup(
						map[string]interface{}{"name": "name"},
						map[string]interface{}{
							"id":     "id",
							"name":   "name",
							"parent": map[string]interface{}{"id": "root"},
							"path":   "id",
							"type":   "networking-uxi/group",
						},
						1,
					)
					util.MockGetGroup(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								{
									"id":     "id",
									"name":   "name",
									"parent": map[string]interface{}{"id": "root"},
									"path":   "id",
									"type":   "networking-uxi/group",
								},
							},
						),
						1,
					)
					// to indicate the group has the root group as a parent
					util.MockGetGroup(
						"root",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								{
									"id":   "root",
									"name": "root",
									"path": "root",
									"type": "networking-uxi/group",
								},
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_group" "my_group" {
					name = "name"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "id"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name"),
					resource.TestCheckNoResourceAttr("uxi_group.my_group", "parent_group_id"),
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								{
									"id":     "id",
									"name":   "name",
									"parent": map[string]interface{}{"id": "root"},
									"path":   "id",
									"type":   "networking-uxi/group",
								},
							},
						),
						1,
					)
					util.MockDeleteGroup("id", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestGroupResourcemockTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response
	var updateTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/groups").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPostGroup(
						util.GenerateGroupRequestModel("id", "", ""),
						util.GenerateNonRootGroupResponseModel("id", "", ""),
						1,
					)
					util.MockGetGroup(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("id", "", ""),
							},
						),
						2,
					)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_id", "", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "id"),
					func(s *terraform.State) error {
						st.Assert(t, mockTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Update that does not trigger a recreate
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("id", "", ""),
						},
					),
						1,
					)
					// new group
					update429 = gock.New("https://test.api.capenetworks.com").
						Patch("/networking-uxi/v1alpha1/groups/id").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockUpdateGroup(
						"id",
						map[string]interface{}{"name": "name_2"},
						util.GenerateNonRootGroupResponseModel("id", "_2", ""),
						1,
					)
					util.MockGetGroup("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("id", "_2", ""),
						},
					),
						3,
					)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_id", "", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_id"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name_2"),
					func(s *terraform.State) error {
						st.Assert(t, updateTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					util.MockGetGroup("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("id", "", ""),
						},
					),
						1,
					)
					util.MockDeleteGroup("id", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestGroupResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// read 5xx error
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/groups").
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

					import {
						to = uxi_group.my_group
						id = "id"
					}
				`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Read not found
			{
				PreConfig: func() {
					util.MockGetGroup(
						"id",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					import {
						to = uxi_group.my_group
						id = "id"
					}
				`,

				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Create 4xx
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/groups").
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
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Create group in prep for next step
			{
				PreConfig: func() {
					util.MockPostGroup(
						util.GenerateGroupRequestModel("id", "", ""),
						util.GenerateNonRootGroupResponseModel("id", "", ""),
						1,
					)
					util.MockGetGroup(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("id", "", ""),
							},
						),
						2,
					)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_id", "", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "id"),
				),
			},
			// Update 4xx
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("id", "", ""),
						},
					),
						1,
					)
					// new group - with error
					gock.New("https://test.api.capenetworks.com").
						Patch("/networking-uxi/v1alpha1/groups/id").
						Reply(http.StatusUnprocessableEntity).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusUnprocessableEntity,
							"errorCode":      "HPE_GL_UXI_DUPLICATE_SIBLING_GROUP_NAME",
							"message":        "Unable to create group - a sibling group already has the specified name",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_id"
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Unable to create group - a sibling group already has the specified name\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Delete 4xx
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("id", "", ""),
						},
					),
						1,
					)
					// delete group - with error
					gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/groups/id").
						Reply(http.StatusUnprocessableEntity).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusUnprocessableEntity,
							"errorCode":      "HPE_GL_UXI_GROUP_CANNOT_BE_DELETED",
							"message":        "Unable to delete group",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig,
				ExpectError: regexp.MustCompile(
					`(?s)Unable to delete group\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Actually delete group for cleanup reasons
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("id", "", ""),
						},
					),
						1,
					)
					// delete group
					util.MockDeleteGroup("id", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
