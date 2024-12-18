/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resource_test

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/stretchr/testify/assert"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func setupGroupCreateMocks(groupID string, parentID string, name string) {
	var postRequest config_api_client.GroupPostRequest
	if parentID == util.MockRootGroupID {
		postRequest = createGroupPostRequest(name, nil)
	} else {
		postRequest = createGroupPostRequest(name, &parentID)
	}

	groupPath := parentID + "." + groupID
	postResponse := createGroupPostResponse(groupID, groupPath, name, parentID)
	util.MockPostGroup(postRequest, postResponse, 1)

	getResponse := createGroupGetResponse(groupID, parentID, groupPath, name)
	util.MockGetGroup(groupID, getResponse, 1)
}

func setupGroupUpdateMocks(groupID string, parentID string, name string) {
	path := parentID + "." + groupID
	if parentID == util.MockRootGroupID {
		util.MockGetGroup(util.MockRootGroupID, util.GenerateRootGroupGetResponse(), 1)
	} else {
		util.MockGetGroup(
			parentID,
			createGroupGetResponse(
				parentID,
				"fake_parent_id",
				"fake_parent_id."+parentID,
				"fake_parent",
			),
			1,
		)
	}

	util.MockPatchGroup(
		groupID,
		createGroupPatchRequest(name),
		createGroupPatchResponse(groupID, name, path, parentID),
		1,
	)

	util.MockGetGroup(
		groupID,
		createGroupGetResponse(groupID, parentID, path, name),
		1,
	)
}

func setupGroupDeleteMocks(groupID string, parentID string, name string) {
	groupPath := parentID + "." + groupID
	getResponse := createGroupGetResponse(groupID, parentID, groupPath, name)
	util.MockGetGroup(
		groupID,
		getResponse,
		1,
	)
	util.MockDeleteGroup(groupID, 1)
}

func setupGroupImportMocks(groupID string, parentID string, name string) {
	var getParentResponse config_api_client.GroupsGetResponse
	if parentID == util.MockRootGroupID {
		getParentResponse = util.GenerateRootGroupGetResponse()
	} else {
		getParentResponse = createGroupGetResponse(
			parentID,
			"fake_parent_id",
			"fake_parent_id."+parentID,
			"fake parent",
		)
	}

	util.MockGetGroup(parentID, getParentResponse, 1)
	groupPath := parentID + "." + groupID
	getResponse := createGroupGetResponse(groupID, parentID, groupPath, name)
	util.MockGetGroup(groupID, getResponse, 1)
}

func createGroupPostRequest(name string, parentID *string) config_api_client.GroupPostRequest {
	postRequest := config_api_client.NewGroupPostRequest(name)
	if parentID != nil {
		realParent := *parentID
		postRequest.SetParentId(realParent)
	}

	return *postRequest
}

func createGroupPostResponse(
	groupID string,
	groupPath string,
	name string,
	parentID string,
) config_api_client.GroupPostResponse {
	return *config_api_client.NewGroupPostResponse(
		groupID,
		name,
		groupPath,
		*config_api_client.NewGroupPostParent(parentID),
		shared.GroupType,
	)
}

func createGroupGetResponse(
	groupID string,
	parentID string,
	groupPath string,
	name string,
) config_api_client.GroupsGetResponse {
	getResponseItems := []config_api_client.GroupsGetItem{*config_api_client.NewGroupsGetItem(
		groupID,
		name,
		*config_api_client.NewNullableGroupsGetParent(config_api_client.NewGroupsGetParent(parentID)),
		groupPath,
		shared.GroupType,
	)}
	getResponse := config_api_client.NewGroupsGetResponse(
		getResponseItems,
		1,
		*config_api_client.NewNullableString(nil),
	)

	return *getResponse
}

func createGroupPatchRequest(name string) config_api_client.GroupPatchRequest {
	request := config_api_client.NewGroupPatchRequest()
	request.SetName(name)

	return *request
}

func createGroupPatchResponse(
	groupID string,
	name string,
	parentID string,
	path string,
) config_api_client.GroupPatchResponse {
	response := config_api_client.NewGroupPatchResponse(
		groupID,
		name,
		path,
		*config_api_client.NewGroupPatchParent(parentID),
		shared.GroupType,
	)

	return *response
}

func Test_CreateGroupResource_ShouldSucceed(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()
	testParentID := "create_parent"
	testGroupID := "create_id"
	testName := "test_name"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					setupGroupCreateMocks(testGroupID, testParentID, testName)
				},
				Config: fmt.Sprintf(`%s
				resource "hpeuxi_group" "my_group" {
					name            = "%s"
					parent_group_id = "%s"
				}`, provider.ProviderConfig, testName, testParentID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"hpeuxi_group.my_group",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectUnknownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("id"),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("parent_group_id"),
							knownvalue.StringExact("create_parent"),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("name"),
							knownvalue.StringExact(testName),
						),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", testName),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						"create_parent",
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", testGroupID),
				),
			},
			// Delete
			{
				PreConfig: func() {
					setupGroupDeleteMocks(testGroupID, testParentID, testName)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_ImportGroupResource_ShouldSucceed(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()
	testParentID := "import_parent"
	testName := "import_name"
	testID := "import_id"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					setupGroupCreateMocks(testID, testParentID, testName)
				},
				Config: fmt.Sprintf(`%s
				resource "hpeuxi_group" "my_group" {
					name            = "%s"
					parent_group_id = "%s"
				}`, provider.ProviderConfig, testName, testParentID),
			},
			// ImportState
			{
				PreConfig: func() {
					setupGroupImportMocks(testID, testParentID, testName)
				},
				ResourceName:      "hpeuxi_group.my_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete
			{
				PreConfig: func() {
					setupGroupDeleteMocks(testID, testParentID, testName)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_CreateGroupResource_WithRootParent(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()
	testGroupID := "root_child"
	testName := "child of root"
	rootParentID := util.MockRootGroupID

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Creating a group attached to the root
			{
				PreConfig: func() {
					setupGroupCreateMocks(testGroupID, rootParentID, testName)
				},
				Config: fmt.Sprintf(`%s
				resource "hpeuxi_group" "my_group" {
					name            = "%s"
				}`, provider.ProviderConfig, testName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"hpeuxi_group.my_group",
							plancheck.ResourceActionCreate,
						),
						plancheck.ExpectUnknownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("id"),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("parent_group_id"),
							knownvalue.Null(),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("name"),
							knownvalue.StringExact(testName),
						),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", testGroupID),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", testName),
					resource.TestCheckNoResourceAttr("hpeuxi_group.my_group", "parent_group_id"),
				),
			},
			// Delete testing
			{
				PreConfig: func() {
					// existing group
					setupGroupDeleteMocks(testGroupID, rootParentID, testName)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_ImportGroupResource_WithRoot_ShouldFail(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Importing the root group does not work
			{
				PreConfig: func() {
					setupGroupImportMocks(util.MockRootGroupID, util.MockRootGroupID, "root")
				},
				Config: provider.ProviderConfig + `
				resource "hpeuxi_group" "my_root_group" {
					name = "root"
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
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()
	testParentID := "parent_id"
	testGroupID := "no_recreate"

	oldTestName := "create_name"
	newTestName := "update_name"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					setupGroupCreateMocks(testGroupID, testParentID, oldTestName)
				},
				Config: fmt.Sprintf(`%s
				resource "hpeuxi_group" "my_group" {
					name            = "%s"
					parent_group_id = "%s"
				}`, provider.ProviderConfig, oldTestName, testParentID),
			},
			// Update that does not trigger a recreate
			{
				PreConfig: func() {
					setupGroupImportMocks(testGroupID, testParentID, oldTestName)

					// updated group
					setupGroupUpdateMocks(testGroupID, testParentID, newTestName)
				},
				Config: fmt.Sprintf(`%s
				resource "hpeuxi_group" "my_group" {
					name            = "%s"
					parent_group_id = "%s"
				}`, provider.ProviderConfig, newTestName, testParentID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"hpeuxi_group.my_group",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("name"),
							knownvalue.StringExact(newTestName),
						),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", newTestName),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						testParentID,
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", testGroupID),
				),
			},
			// Delete
			{
				PreConfig: func() {
					setupGroupDeleteMocks(testGroupID, testParentID, newTestName)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_UpdateGroupResource_WithoutParent_ShouldSucceed(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()
	testGroupID := "id"
	oldTestName := "create_name"
	newTestName := "update_name"
	rootParentID := util.MockRootGroupID

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					setupGroupCreateMocks(testGroupID, rootParentID, oldTestName)
				},
				Config: fmt.Sprintf(`%s
				resource "hpeuxi_group" "my_group" {
					name            = "%s"
				}`, provider.ProviderConfig, oldTestName),
			},
			// Update that does not trigger a recreate
			{
				PreConfig: func() {
					setupGroupImportMocks(testGroupID, rootParentID, oldTestName)
					setupGroupUpdateMocks(testGroupID, util.MockRootGroupID, newTestName)
				},
				Config: fmt.Sprintf(`%s
				resource "hpeuxi_group" "my_group" {
					name            = "%s"
				}`, provider.ProviderConfig, newTestName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"hpeuxi_group.my_group",
							plancheck.ResourceActionUpdate,
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("name"),
							knownvalue.StringExact(newTestName),
						),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", newTestName),
					resource.TestCheckNoResourceAttr("hpeuxi_group.my_group", "parent_group_id"),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", testGroupID),
				),
			},
			// Delete
			{
				PreConfig: func() {
					setupGroupDeleteMocks(testGroupID, rootParentID, newTestName)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func Test_UpdateGroupResource_WithRecreate_ShouldSucceed(t *testing.T) {
	defer gock.OffAll()
	mockOAuth := util.MockOAuth()

	oldParentID := "parent_id"
	oldGroupID := "id"

	newParentID := "parent_id_2"
	newGroupID := "new_id"

	testName := "constant_name"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				PreConfig: func() {
					setupGroupCreateMocks(oldGroupID, oldParentID, testName)
				},
				Config: fmt.Sprintf(`%s
				resource "hpeuxi_group" "my_group" {
					name            = "%s"
					parent_group_id = "%s"
				}`, provider.ProviderConfig, testName, oldParentID),
			},
			// Update that does trigger a recreate
			{
				PreConfig: func() {
					setupGroupImportMocks(oldGroupID, oldParentID, testName)
					setupGroupDeleteMocks(oldGroupID, oldParentID, testName)

					setupGroupCreateMocks(newGroupID, newParentID, testName)
				},
				Config: fmt.Sprintf(`%s
				resource "hpeuxi_group" "my_group" {
					name            = "%s"
					parent_group_id = "%s"
				}`, provider.ProviderConfig, testName, newParentID),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(
							"hpeuxi_group.my_group",
							plancheck.ResourceActionDestroyBeforeCreate,
						),
						plancheck.ExpectUnknownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("id"),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("name"),
							knownvalue.StringExact(testName),
						),
						plancheck.ExpectKnownValue(
							"hpeuxi_group.my_group",
							tfjsonpath.New("parent_group_id"),
							knownvalue.StringExact(newParentID),
						),
					},
				},
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "name", testName),
					resource.TestCheckResourceAttr(
						"hpeuxi_group.my_group",
						"parent_group_id",
						newParentID,
					),
					resource.TestCheckResourceAttr("hpeuxi_group.my_group", "id", newGroupID),
				),
			},
			// Delete
			{
				PreConfig: func() {
					setupGroupDeleteMocks(newGroupID, newParentID, testName)
				},
				Config: provider.ProviderConfig,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestGroupResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.OffAll()
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
	defer gock.OffAll()
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
