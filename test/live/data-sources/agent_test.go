package data_source_test

import (
	"context"
	"testing"

	"github.com/aruba-uxi/terraform-provider-configuration/test/live/config"
	"github.com/aruba-uxi/terraform-provider-configuration/test/live/provider"
	"github.com/aruba-uxi/terraform-provider-configuration/test/live/util"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/nbio/st"
)

type agentProperties struct {
	id                   string
	serial               string
	name                 string
	model_number         *string
	wifi_mac_address     *string
	ethernet_mac_address *string
	notes                *string
	pcapMode             *string
}

func TestAgentDataSource(t *testing.T) {
	agent := getAgentProperties(config.AgentUid)
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
				Check: checkStateAgainstAgent(t, agent),
			},
		},
	})
}

func checkStateAgainstAgent(t st.Fatalf, agent agentProperties) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "id", config.AgentUid),
		resource.TestCheckResourceAttr("data.uxi_agent.my_agent", "serial", agent.serial),
		func() resource.TestCheckFunc {
			if agent.wifi_mac_address == nil {
				return resource.TestCheckNoResourceAttr(
					"data.uxi_agent.my_agent",
					"model_number",
				)
			} else {
				return resource.TestCheckResourceAttrWith(
					"data.uxi_agent.my_agent",
					"model_number",
					func(value string) error {
						st.Assert(t, value, agent.model_number)
						return nil
					},
				)
			}
		}(),
		resource.TestCheckResourceAttrWith(
			"data.uxi_agent.my_agent",
			"name",
			func(value string) error {
				st.Assert(t, value, agent.name)
				return nil
			},
		),
		func() resource.TestCheckFunc {
			if agent.wifi_mac_address == nil {
				return resource.TestCheckNoResourceAttr(
					"data.uxi_agent.my_agent",
					"wifi_mac_address",
				)
			} else {
				return resource.TestCheckResourceAttrWith(
					"data.uxi_agent.my_agent",
					"wifi_mac_address",
					func(value string) error {
						st.Assert(t, value, agent.wifi_mac_address)
						return nil
					},
				)
			}
		}(),
		func() resource.TestCheckFunc {
			if agent.wifi_mac_address == nil {
				return resource.TestCheckNoResourceAttr(
					"data.uxi_agent.my_agent",
					"ethernet_mac_address",
				)
			} else {
				return resource.TestCheckResourceAttrWith(
					"data.uxi_agent.my_agent",
					"ethernet_mac_address",
					func(value string) error {
						st.Assert(t, value, agent.ethernet_mac_address)
						return nil
					},
				)
			}
		}(),
		func() resource.TestCheckFunc {
			if agent.wifi_mac_address == nil {
				return resource.TestCheckNoResourceAttr(
					"data.uxi_agent.my_agent",
					"notes",
				)
			} else {
				return resource.TestCheckResourceAttrWith(
					"data.uxi_agent.my_agent",
					"notes",
					func(value string) error {
						st.Assert(t, value, agent.notes)
						return nil
					},
				)
			}
		}(),
		func() resource.TestCheckFunc {
			if agent.pcapMode == nil {
				return resource.TestCheckNoResourceAttr(
					"data.uxi_agent.my_agent",
					"pcap_mode",
				)
			} else {
				return resource.TestCheckResourceAttrWith(
					"data.uxi_agent.my_agent",
					"pcap_mode",
					func(value string) error {
						st.Assert(t, value, agent.pcapMode)
						return nil
					},
				)
			}
		}(),
	)
}

func getAgentProperties(id string) agentProperties {
	result, _, err := util.Client.ConfigurationAPI.
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
	return agentProperties{
		id:                   agent.Id,
		serial:               agent.Serial,
		name:                 agent.Name,
		model_number:         agent.ModelNumber.Get(),
		wifi_mac_address:     agent.WifiMacAddress.Get(),
		ethernet_mac_address: agent.EthernetMacAddress.Get(),
		notes:                agent.Notes.Get(),
		pcapMode:             agent.PcapMode.Get(),
	}
}
