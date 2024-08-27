package test

import (
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
					response := resources.GroupResponseModel{
						UID:       "uid",
						Name:      "name",
						ParentUid: "parent_uid",
						Path:      "parent_uid.uid",
					}
					resources.CreateGroup = func(request resources.GroupCreateRequestModel) resources.GroupResponseModel {
						return response
					}
					resources.GetGroup = func() resources.GroupResponseModel {
						return response
					}
				},
				Config: providerConfig + `
				resource "uxi_group" "my_group" {
					name       = "name"
					parent_uid = "parent_uid"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "parent_uid", "parent_uid"),
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
					response := resources.GroupResponseModel{
						UID:       "uid",
						Name:      "updated_name",
						ParentUid: "parent_uid",
						Path:      "parent_uid.uid",
					}
					resources.UpdateGroup = func(request resources.GroupUpdateRequestModel) resources.GroupResponseModel {
						return response
					}
					resources.GetGroup = func() resources.GroupResponseModel {
						return response
					}
				},
				Config: providerConfig + `
					resource "uxi_group" "my_group" {
						name       = "updated_name"
						parent_uid = "parent_uid"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "updated_name"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "parent_uid", "parent_uid"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "uid"),
				),
				Destroy: false,
			},
			// Update that does trigger a recreate
			{
				PreConfig: func() {
					response := resources.GroupResponseModel{
						UID:       "new_uid",
						Name:      "name",
						ParentUid: "updated_parent_uid",
						Path:      "updated_parent_uid.uid",
					}
					resources.CreateGroup = func(request resources.GroupCreateRequestModel) resources.GroupResponseModel {
						return response
					}
					called := false
					resources.GetGroup = func() resources.GroupResponseModel {
						if !called {
							called = true
							return resources.GroupResponseModel{
								UID:       "uid",
								Name:      "name",
								ParentUid: "parent_uid",
								Path:      "parent_uid.uid",
							}
						} else {
							return resources.GroupResponseModel{
								UID:       "new_uid",
								Name:      "name",
								ParentUid: "updated_parent_uid",
								Path:      "updated_parent_uid.uid",
							}
						}
					}
				},
				Config: providerConfig + `
					resource "uxi_group" "my_group" {
						name       = "name"
						parent_uid = "updated_parent_uid"
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.my_group", "name", "name"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "parent_uid", "updated_parent_uid"),
					resource.TestCheckResourceAttr("uxi_group.my_group", "id", "new_uid"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
