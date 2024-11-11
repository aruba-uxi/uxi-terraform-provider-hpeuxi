package util

import (
	"context"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
)

func GetAgentProperties(id string) config_api_client.AgentItem {
	result, _, err := Client.ConfigurationAPI.
		AgentsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("agent with id `" + id + "` could not be found")
	}
	return result.Items[0]
}

func CheckStateAgainstAgent(t st.Fatalf, agent config_api_client.AgentItem) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "id", config.AgentUid),
		resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "serial", agent.Serial),
		TestOptionalValue(t, "data.uxi_agent.my_agent", "model_number", agent.ModelNumber.Get()),
		resource.TestCheckResourceAttrWith(
			"data.uxi_agent.my_agent",
			"name",
			func(value string) error {
				st.Assert(t, value, agent.Name)
				return nil
			},
		),
		TestOptionalValue(
			t,
			"data.uxi_agent.my_agent",
			"wifi_mac_address",
			agent.WifiMacAddress.Get(),
		),
		TestOptionalValue(
			t,
			"data.uxi_agent.my_agent",
			"ethernet_mac_address",
			agent.EthernetMacAddress.Get(),
		),
		TestOptionalValue(t, "data.uxi_agent.my_agent", "notes", agent.Notes.Get()),
		TestOptionalValue(t, "data.uxi_agent.my_agent", "pcap_mode", agent.PcapMode.Get()),
	)
}
