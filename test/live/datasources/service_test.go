/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package data_source_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func TestServiceTestDataSource(t *testing.T) {
	serviceTest := util.GetServiceTest(config.ServiceTestId)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					data "uxi_service_test" "my_service_test" {
						filter = {
							id = "` + config.ServiceTestId + `"
						}
					}
				`,
				Check: shared.CheckStateAgainstServiceTest(
					t,
					"data.uxi_service_test.my_service_test",
					serviceTest,
				),
			},
		},
	})
}
