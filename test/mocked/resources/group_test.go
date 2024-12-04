/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
)

func Test_CreateGroupResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("id", "", ""),
						util.GenerateGroupPostResponse("id", "", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 2)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						"parent_id",
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
				),
			},
			// Update that does not trigger a recreate
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// updated group
					util.MockPatchGroup(
						"id",
						util.GenerateGroupPatchRequest("_2"),
						util.GenerateGroupPatchResponse("id", "_2", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "_2", ""), 3)
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_id"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name_2"),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						"parent_id",
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
				),
			},
			// Update that does trigger a recreate
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// new group (replacement)
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("new_id", "", "_2"),
						util.GenerateGroupPostResponse("new_id", "", "_2"),
						1,
					)
					util.MockGetGroup(
						"new_id",
						util.GenerateGroupGetResponse("new_id", "", "_2"),
						1,
					)
					// delete old group (being replaced)
					util.MockDeleteGroup("id", 1)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id_2"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						"parent_id_2",
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "new_id"),
				),
			},
			// Delete
			{
				PreConfig: func() {
					util.MockGetGroup(
						"new_id",
						util.GenerateGroupGetResponse("new_id", "", "_2"),
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

func Test_ImportGroupResource_ShouldSucceed(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("id", "", ""),
						util.GenerateGroupPostResponse("id", "", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 2)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
			},
			// ImportState
			{
				PreConfig: func() {
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				ResourceName:      "hpeuxi_group.my_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete
			{
				PreConfig: func() {
					util.MockGetGroup(
						"id",
						util.GenerateGroupGetResponse("id", "", ""),
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

func Test_CreateGroupResource_WithRootParent(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a group attached to the root
			{
				PreConfig: func() {
					util.MockPostGroup(
						util.GenerateGroupAttachedToRootGroupPostRequest("id", ""),
						util.GenerateGroupAttachedToRootGroupPostResponse("id", ""),
						1,
					)
					// to indicate the group has the root group as a parent
					util.MockGetGroup(util.MockRootGroupID, util.GenerateRootGroupGetResponse(), 1)
					util.MockGetGroup(
						"id",
						util.GenerateGroupAttachedToRootGroupGetResponse("id", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name = "name"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name"),
					resource.TestCheckNoResourceAttr("hpeuxi_group.my_group", "parent_group_id"),
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup(
						"id",
						util.GenerateGroupAttachedToRootGroupGetResponse("id", ""),
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

func Test_ImportGroupResource_WithRootParent_ShouldFail(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Importing the root group does not work
			{
				PreConfig: func() {
					util.MockGetGroup(util.MockRootGroupID, util.GenerateRootGroupGetResponse(), 1)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_root_group" {
					name            = "name"
				}

				import {
					to = hpeuxi_group.my_root_group
					id = "` + util.MockRootGroupID + `"
				}`,
				ExpectError: regexp.MustCompile(`The root group cannot be used as a resource`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_UpdateGroupResource_WithoutRecreate_ShouldSucceed(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("id", "", ""),
						util.GenerateGroupPostResponse("id", "", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 2)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						"parent_id",
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
				),
			},
			// Update that does not trigger a recreate
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// updated group
					util.MockPatchGroup(
						"id",
						util.GenerateGroupPatchRequest("_2"),
						util.GenerateGroupPatchResponse("id", "_2", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "_2", ""), 3)
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_id"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name_2"),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						"parent_id",
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
				),
			},
			// Delete
			{
				PreConfig: func() {
					util.MockGetGroup(
						"id",
						util.GenerateGroupGetResponse("id", "", "_2"),
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

func Test_UpdateGroupResource_WithRecreate_ShouldSucceed(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("id", "", ""),
						util.GenerateGroupPostResponse("id", "", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 2)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
			},
			// Update that does trigger a recreate
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// new group (replacement)
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("new_id", "", "_2"),
						util.GenerateGroupPostResponse("new_id", "", "_2"),
						1,
					)
					util.MockGetGroup(
						"new_id",
						util.GenerateGroupGetResponse("new_id", "", "_2"),
						1,
					)
					// delete old group (being replaced)
					util.MockDeleteGroup("id", 1)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id_2"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						"parent_id_2",
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "new_id"),
				),
			},
			// Delete
			{
				PreConfig: func() {
					util.MockGetGroup(
						"new_id",
						util.GenerateGroupGetResponse("new_id", "", "_2"),
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

func TestGroupResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Post("/networking-uxi/v1alpha1/groups").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPostGroup(
						util.GenerateNonRootGroupPostRequest("id", "", ""),
						util.GenerateGroupPostResponse("id", "", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 2)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Read
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Post("/networking-uxi/v1alpha1/groups").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				ResourceName:      "hpeuxi_group.my_group",
				ImportState:       true,
				ImportStateVerify: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Update
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// new group
					mockTooManyRequests = gock.New(util.MockUXIURL).
						Patch("/networking-uxi/v1alpha1/groups/id").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockPatchGroup(
						"id",
						util.GenerateGroupPatchRequest("_2"),
						util.GenerateGroupPatchResponse("id", "_2", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "_2", ""), 3)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_id"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", "name_2"),
					func(s *terraform.State) error {
						assert.Equal(t, 0, mockTooManyRequests.Mock.Request().Counter)

						return nil
					},
				),
			},
			// Delete
			{
				PreConfig: func() {
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
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
			// read HTTP error
			{
				PreConfig: func() {
					gock.New(util.MockUXIURL).
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
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					import {
						to = hpeuxi_group.my_group
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
					util.MockGetGroup("id", util.EmptyGetListResponse, 1)
				},
				Config: provider.ProviderConfig + `
					resource "hpeuxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_id"
					}

					import {
						to = hpeuxi_group.my_group
						id = "id"
					}
				`,

				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Create HTTP error
			{
				PreConfig: func() {
					gock.New(util.MockUXIURL).
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
				resource "hpeuxi_group" "my_group" {
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
						util.GenerateNonRootGroupPostRequest("id", "", ""),
						util.GenerateGroupPostResponse("id", "", ""),
						1,
					)
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 2)
					// to indicate the group has a parent
					util.MockGetGroup(
						"parent_id",
						util.GenerateGroupGetResponse("parent_id", "", ""),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_group" {
					name            = "name"
					parent_group_id = "parent_id"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", "id"),
				),
			},
			// Update HTTP error
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// new group - with error
					gock.New(util.MockUXIURL).
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
					resource "hpeuxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_id"
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Unable to create group - a sibling group already has the specified name\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			// Delete HTTP error
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// delete group - with error
					gock.New(util.MockUXIURL).
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
					util.MockGetGroup("id", util.GenerateGroupGetResponse("id", "", ""), 1)
					// delete group
					util.MockDeleteGroup("id", 1)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}
