package util

import (
	"context"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
)

type AgentProperties struct {
	Id                 string
	Serial             string
	Name               string
	ModelNumber        *string
	WifiMacAddress     *string
	EthernetMacAddress *string
	Notes              *string
	PcapMode           *string
}

func GetAgentProperties(id string) AgentProperties {
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
	agent := result.Items[0]
	// Read these in, as they may not be always constant with the acceptance test customer
	return AgentProperties{
		Id:                 agent.Id,
		Serial:             agent.Serial,
		Name:               agent.Name,
		ModelNumber:        agent.ModelNumber.Get(),
		WifiMacAddress:     agent.WifiMacAddress.Get(),
		EthernetMacAddress: agent.EthernetMacAddress.Get(),
		Notes:              agent.Notes.Get(),
		PcapMode:           agent.PcapMode.Get(),
	}
}

func CheckStateAgainstAgent(t st.Fatalf, agent AgentProperties) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "id", config.AgentUid),
		resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "serial", agent.Serial),
		TestOptionalValue(t, "data.uxi_agent.my_agent", "model_number", agent.ModelNumber),
		resource.TestCheckResourceAttrWith(
			"data.uxi_agent.my_agent",
			"name",
			func(value string) error {
				st.Assert(t, value, agent.Name)
				return nil
			},
		),
		TestOptionalValue(t, "data.uxi_agent.my_agent", "wifi_mac_address", agent.WifiMacAddress),
		TestOptionalValue(
			t,
			"data.uxi_agent.my_agent",
			"ethernet_mac_address",
			agent.EthernetMacAddress,
		),
		TestOptionalValue(t, "data.uxi_agent.my_agent", "notes", agent.Notes),
		TestOptionalValue(t, "data.uxi_agent.my_agent", "pcap_mode", agent.PcapMode),
	)
}
