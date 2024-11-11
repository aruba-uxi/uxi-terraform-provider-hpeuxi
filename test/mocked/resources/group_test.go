package resource_test

import (
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
						util.GenerateGroupRequestModel("uid", "", ""),
						util.StructToMap(util.GenerateNonRootGroupResponseModel("uid", "", "")),
						1,
					)
					util.MockGetGroup("uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("uid", "", ""),
							},
						),
						2,
					)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_uid", "", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_uid"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr(
						"uxi_group.my_group",
						"parent_group_id",
						"parent_uid",
					),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "uid"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetGroup(
						"uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("uid", "", ""),
							},
						),
						1,
					)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_uid", "", ""),
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
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("uid", "", ""),
						},
					),
						1,
					)
					// updated group
					util.MockUpdateGroup(
						"uid",
						map[string]interface{}{"name": "name_2"},
						util.GenerateNonRootGroupResponseModel("uid", "_2", ""),
						1,
					)
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("uid", "_2", ""),
						},
					),
						3,
					)
					util.MockGetGroup(
						"parent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_uid", "", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_uid"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name_2"),
					resource.TestCheckResourceAttr(
						"uxi_group.my_group",
						"parent_group_id",
						"parent_uid",
					),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "uid"),
				),
			},
			// Update that does trigger a recreate
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("uid", "", ""),
						},
					),
						1,
					)
					// new group (replacement)
					util.MockPostGroup(
						util.GenerateGroupRequestModel("new_uid", "", "_2"),
						util.GenerateNonRootGroupResponseModel("new_uid", "", "_2"),
						1,
					)
					util.MockGetGroup("new_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("new_uid", "", "_2"),
						},
					),
						1,
					)
					// delete old group (being replaced)
					util.MockDeleteGroup("uid", 1)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid_2"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr(
						"uxi_group.my_group",
						"parent_group_id",
						"parent_uid_2",
					),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "new_uid"),
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					util.MockGetGroup("new_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("new_uid", "", "_2"),
						},
					),
						1,
					)
					util.MockDeleteGroup("new_uid", 1)
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
						"my_root_group_uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{
							{
								"id":     "my_root_group_uid",
								"name":   "root",
								"parent": *config_api_client.NewNullableParent(nil),
								"path":   "my_root_group_uid",
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
					id = "my_root_group_uid"
				}`,
				ExpectError: regexp.MustCompile(`The root group cannot be used as a resource`),
			},
			// Creating a group attached to the root
			{
				PreConfig: func() {
					util.MockPostGroup(
						map[string]interface{}{"name": "name"},
						map[string]interface{}{
							"id":     "uid",
							"name":   "name",
							"parent": map[string]interface{}{"id": "root"},
							"path":   "uid",
							"type":   "networking-uxi/group",
						},
						1,
					)
					util.MockGetGroup(
						"uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								{
									"id":     "uid",
									"name":   "name",
									"parent": map[string]interface{}{"id": "root"},
									"path":   "uid",
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
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "uid"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name"),
					resource.TestCheckNoResourceAttr("uxi_group.my_group", "parent_group_id"),
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup(
						"uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								{
									"id":     "uid",
									"name":   "name",
									"parent": map[string]interface{}{"id": "root"},
									"path":   "uid",
									"type":   "networking-uxi/group",
								},
							},
						),
						1,
					)
					util.MockDeleteGroup("uid", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestGroupResource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var request429 *gock.Response
	var update429 *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				PreConfig: func() {
					request429 = gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/groups").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPostGroup(
						util.GenerateGroupRequestModel("uid", "", ""),
						util.GenerateNonRootGroupResponseModel("uid", "", ""),
						1,
					)
					util.MockGetGroup(
						"uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("uid", "", ""),
							},
						),
						2,
					)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_uid", "", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_uid"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "uid"),
					func(s *terraform.State) error {
						st.Assert(t, request429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Update that does not trigger a recreate
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("uid", "", ""),
						},
					),
						1,
					)
					// new group
					update429 = gock.New("https://test.api.capenetworks.com").
						Patch("/networking-uxi/v1alpha1/groups/uid").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockUpdateGroup(
						"uid",
						map[string]interface{}{"name": "name_2"},
						util.GenerateNonRootGroupResponseModel("uid", "_2", ""),
						1,
					)
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("uid", "_2", ""),
						},
					),
						3,
					)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_uid", "", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_uid"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name_2"),
					func(s *terraform.State) error {
						st.Assert(t, update429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("uid", "", ""),
						},
					),
						1,
					)
					util.MockDeleteGroup("uid", 1)
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
						Reply(500).
						JSON(map[string]interface{}{
							"httpStatusCode": 500,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					import {
						to = uxi_group.my_group
						id = "uid"
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
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid"
					}

					import {
						to = uxi_group.my_group
						id = "uid"
					}
				`,

				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Create 4xx
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/groups").
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
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)Validation error - bad request\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Create group in prep for next step
			{
				PreConfig: func() {
					util.MockPostGroup(
						util.GenerateGroupRequestModel("uid", "", ""),
						util.GenerateNonRootGroupResponseModel("uid", "", ""),
						1,
					)
					util.MockGetGroup(
						"uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("uid", "", ""),
							},
						),
						2,
					)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_uid",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateNonRootGroupResponseModel("parent_uid", "", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_uid"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "uid"),
				),
			},
			// Update 4xx
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("uid", "", ""),
						},
					),
						1,
					)
					// new group - with error
					gock.New("https://test.api.capenetworks.com").
						Patch("/networking-uxi/v1alpha1/groups/uid").
						Reply(422).
						JSON(map[string]interface{}{
							"httpStatusCode": 422,
							"errorCode":      "HPE_GL_UXI_DUPLICATE_SIBLING_GROUP_NAME",
							"message":        "Unable to create group - a sibling group already has the specified name",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_uid"
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Unable to create group - a sibling group already has the specified name\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Delete 4xx
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("uid", "", ""),
						},
					),
						1,
					)
					// delete group - with error
					gock.New("https://test.api.capenetworks.com").
						Delete("/networking-uxi/v1alpha1/groups/uid").
						Reply(422).
						JSON(map[string]interface{}{
							"httpStatusCode": 422,
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
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{
							util.GenerateNonRootGroupResponseModel("uid", "", ""),
						},
					),
						1,
					)
					// delete group
					util.MockDeleteGroup("uid", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
