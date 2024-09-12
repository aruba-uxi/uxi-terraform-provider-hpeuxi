package test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

type Fetcher interface {
	FetchData() ([]byte, error)
}

func TestGroupResource(t *testing.T) {
	defer gock.Off()
	MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				PreConfig: func() {
					MockPostGroup(StructToMap(GenerateGroupResponseModel("uid", "", "")), 1)
					MockGetGroup("uid", GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							StructToMap(GenerateGroupResponseModel("uid", "", "")),
						}),
						1,
					)
				},
				Config: providerConfig + `
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
					MockGetGroup("uid", GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							StructToMap(GenerateGroupResponseModel("uid", "", "")),
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
						return GenerateGroupResponseModel("uid", "_2", "")
					}
					MockGetGroup("uid", GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							StructToMap(GenerateGroupResponseModel("uid", "_2", "")),
						}),
						2,
					)
				},
				Config: providerConfig + `
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
					MockGetGroup("uid", GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							StructToMap(GenerateGroupResponseModel("uid", "", "")),
						}),
						1,
					)
					// new group (replacement)
					MockPostGroup(StructToMap(GenerateGroupResponseModel("new_uid", "", "_2")), 1)
					MockGetGroup("new_uid", GenerateGroupPaginatedResponse(
						[]map[string]interface{}{
							StructToMap(GenerateGroupResponseModel("new_uid", "", "_2")),
						}),
						1,
					)
				},
				Config: providerConfig + `
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
}

func TestRootGroupResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Importing the root group does not work
			{
				PreConfig: func() {
					resources.GetRootGroupUID = func() string { return "my_root_group_uid" }
				},
				Config: providerConfig + `
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
