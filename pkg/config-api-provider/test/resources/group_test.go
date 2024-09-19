package resource_test

import (
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/provider"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/test/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"regexp"
	"testing"
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
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponseModel("uid", "", "")), 1)
					util.MockGetGroup("uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("uid", "", "")),
						}),
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
					util.MockGetGroup("uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("uid", "", "")),
						}),
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
					resources.UpdateGroup = func(request resources.GroupUpdateRequestModel) resources.GroupResponseModel {
						return util.GenerateGroupResponseModel("uid", "_2", "")
					}
					util.MockGetGroup("uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("uid", "_2", "")),
						}),
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
					util.MockGetGroup("uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("uid", "", "")),
						}),
						1,
					)
					// new group (replacement)
					util.MockPostGroup(util.StructToMap(util.GenerateGroupResponseModel("new_uid", "", "_2")), 1)
					util.MockGetGroup("new_uid", util.GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							util.StructToMap(util.GenerateGroupResponseModel("new_uid", "", "_2")),
						}),
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
