package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

type Fetcher interface {
	FetchData() ([]byte, error)
}

func TestGroupResource(t *testing.T) {
	createGroupConfig := `
		resource "uxi_group" "test_group" {
		  name       = "test_name"
		  parent_uid = "9999"
		}`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + createGroupConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_group.test_group", "name", "test_name"),
					resource.TestCheckResourceAttr("uxi_group.test_group", "parent_uid", "9999"),
					resource.TestCheckResourceAttrSet("uxi_group.test_group", "id"),
					resource.TestCheckResourceAttrSet("uxi_group.test_group", "last_updated"),
				),
			},
		},
	})
}
