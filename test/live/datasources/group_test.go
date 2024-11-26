/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package data_source_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
)

func TestGroupDataSource(t *testing.T) {
	groupName := "tf_provider_acceptance_test_group_datasource"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					data "uxi_group" "my_group" {
						filter = {
							id = "` + config.GroupIdRoot + `"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`The root group cannot be used as a data source`),
			},
			{
				Config: provider.ProviderConfig + `
					// create the resource to use subsequently as datasource
					resource "uxi_group" "my_group_resource" {
						name = "` + groupName + `"
					}

					data "uxi_group" "my_group" {
						filter = {
							id = uxi_group.my_group_resource.id
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrWith(
						"data.uxi_group.my_group",
						"id",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupName).Id)

							return nil
						},
					),
					resource.TestCheckResourceAttr("data.uxi_group.my_group", "name", groupName),
					resource.TestCheckResourceAttrWith(
						"data.uxi_group.my_group",
						"path",
						func(value string) error {
							assert.Equal(t, value, util.GetGroupByName(groupName).Path)

							return nil
						},
					),
					resource.TestCheckResourceAttr(
						"data.uxi_group.my_group",
						"parent_group_id",
						config.GroupIdRoot,
					),
				),
			},
		},
	})
}
