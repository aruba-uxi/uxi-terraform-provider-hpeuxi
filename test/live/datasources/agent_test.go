package data_source_test

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAgentDataSource(t *testing.T) {
	agent := util.GetAgentProperties(config.AgentUid)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: provider.ProviderConfig + `
					data "uxi_agent" "my_agent" {
						filter = {
							agent_id = "` + config.AgentUid + `"
						}
					}
				`,
				Check: util.CheckStateAgainstAgent(t, agent),
			},
		},
	})
}
