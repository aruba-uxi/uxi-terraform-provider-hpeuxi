package test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

type Fetcher interface {
	FetchData() ([]byte, error)
}

func TestGroupResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				PreConfig: func() {
					MockOAuth()
					MockPostGroup(StructToMap(GenerateGroupResponseModel("uid", "", "")), 1)

					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						return GenerateGroupResponseModel(uid, "", "")
					}
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
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						return GenerateGroupResponseModel(uid, "_2", "")
					}
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
					MockOAuth()
					MockPostGroup(StructToMap(GenerateGroupResponseModel("new_uid", "", "_2")), 1)
					resources.GetGroup = func(uid string) resources.GroupResponseModel {
						if uid == "uid" {
							return GenerateGroupResponseModel(uid, "", "")
						} else {
							return GenerateGroupResponseModel(uid, "", "_2")
						}
					}
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
