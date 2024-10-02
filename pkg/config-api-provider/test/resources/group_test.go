package resource_test

import (
	"regexp"
	"testing"

	config_api_client "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
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
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponsePostModel("uid", "", "")), 1)
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseGetModel("uid", "", ""))}),
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
					resource.TestCheckResourceAttr("uxi_group.my_group", "parent_group_id", "parent_uid"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "uid"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseGetModel("uid", "", ""))}),
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
					resources.UpdateGroup = func(request resources.GroupUpdateRequestModel) config_api_client.GroupsPostResponse {
						return util.GenerateGroupResponsePostModel("uid", "_2", "")
					}
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseGetModel("uid", "_2", ""))}),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name_2"
						parent_group_id = "parent_uid"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name_2"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "parent_group_id", "parent_uid"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "uid"),
				),
				Destroy: false,
			},
			// Update that does trigger a recreate
			{
				PreConfig: func() {
					// existing group
					util.MockGetGroup("uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseGetModel("uid", "", ""))}),
						1,
					)
					// new group (replacement)
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponsePostModel("new_uid", "", "_2")), 1)
					util.MockGetGroup("new_uid", util.GeneratePaginatedResponse(
						[]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseGetModel("new_uid", "", "_2"))}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_group" "my_group" {
						name            = "name"
						parent_group_id = "parent_uid_2"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "parent_group_id", "parent_uid_2"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "new_uid"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})

	mockOAuth.Mock.Disable()
}

func TestRootGroupResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Importing the root group does not work
			{
				PreConfig: func() {
					resources.GetRootGroupUID = func() string { return "my_root_group_uid" }
				},
				Config: provider.ProviderConfig + `
				resource "uxi_group" "my_root_group" {
					name            = "name"
					parent_group_id = "some_random_string"
				}

				import {
					to = uxi_group.my_root_group
					id = "my_root_group_uid"
				}`,
				ExpectError: regexp.MustCompile(`the root node cannot be used as a resource`),
			},
		},
	})
}

func TestGroupResource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var create429 *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				PreConfig: func() {
					create429 = gock.New("https://test.api.capenetworks.com").
						Post("/uxi/v1alpha1/groups").
						Reply(429).
						SetHeaders(map[string]string{
							"X-RateLimit-Limit":     "100",
							"X-RateLimit-Remaining": "0",
							"X-RateLimit-Reset":     "1",
						})
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponsePostModel("uid", "", "")), 1)
					util.MockGetGroup(
						"uid",
						util.GeneratePaginatedResponse([]map[string]interface{}{util.StructToMap(util.GenerateGroupResponseGetModel("uid", "", ""))}),
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
						st.Assert(t, create429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// TODO: Test Updating 429s
			// TODO: Test Deleting 429s
		},
	})

	mockOAuth.Mock.Disable()
}
