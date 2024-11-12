package data_source_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestServiceTestDataSource(t *testing.T) {
	serviceTest := util.GetServiceTest(config.ServiceTestUid)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					data "uxi_service_test" "my_service_test" {
						filter = {
							service_test_id = "` + config.ServiceTestUid + `"
						}
					}
				`,
				Check: util.CheckStateAgainstServiceTest(t, serviceTest),
			},
		},
	})
}
