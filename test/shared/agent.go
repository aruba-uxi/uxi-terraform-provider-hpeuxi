/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package shared

import (
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
)

func CheckStateAgainstAgent(
	t *testing.T,
	stateEntity string,
	agent config_api_client.AgentItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(stateEntity, "id", agent.Id),
		resource.TestCheckResourceAttr(stateEntity, "serial", agent.Serial),
		TestOptionalValue(t, stateEntity, "model_number", agent.ModelNumber.Get()),
		resource.TestCheckResourceAttrWith(
			stateEntity,
			"name",
			func(value string) error {
				assert.Equal(t, value, agent.Name)
				return nil
			},
		),
		TestOptionalValue(t, stateEntity, "wifi_mac_address", agent.WifiMacAddress.Get()),
		TestOptionalValue(t, stateEntity, "ethernet_mac_address", agent.EthernetMacAddress.Get()),
		TestOptionalValue(t, stateEntity, "notes", agent.Notes.Get()),
		TestOptionalValue(t, stateEntity, "pcap_mode", agent.PcapMode.Get()),
	)
}
