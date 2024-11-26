/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/h2non/gock"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func GenerateAgentPatchRequest(postfix string) config_api_client.AgentsPatchRequest {
	name := "name" + postfix
	notes := "notes" + postfix
	pcapMode, _ := config_api_client.NewPcapModeFromValue("light")

	return config_api_client.AgentsPatchRequest{
		Name:     &name,
		Notes:    &notes,
		PcapMode: pcapMode,
	}
}

func GenerateAgentPatchResponse(id string, postfix string) config_api_client.AgentsPatchResponse {
	modelNumber := "model_number" + postfix
	wifiMacAddress := "wifi_mac_address" + postfix
	ethernetMacAddress := "ethernet_mac_address" + postfix
	notes := "notes" + postfix
	pcapMode, _ := config_api_client.NewPcapModeFromValue("light")

	return config_api_client.AgentsPatchResponse{
		Id:                 "id",
		Serial:             "serial" + postfix,
		Name:               "name" + postfix,
		ModelNumber:        *config_api_client.NewNullableString(&modelNumber),
		WifiMacAddress:     *config_api_client.NewNullableString(&wifiMacAddress),
		EthernetMacAddress: *config_api_client.NewNullableString(&ethernetMacAddress),
		Notes:              *config_api_client.NewNullableString(&notes),
		PcapMode:           *config_api_client.NewNullablePcapMode(pcapMode),
		Type:               shared.AgentType,
	}
}

func GenerateAgentResponse(id string, postfix string) config_api_client.AgentsResponse {
	modelNumber := "model_number" + postfix
	wifiMacAddress := "wifi_mac_address" + postfix
	ethernetMacAddress := "ethernet_mac_address" + postfix
	notes := "notes" + postfix
	pcapMode := "light"

	return config_api_client.AgentsResponse{
		Items: []config_api_client.AgentItem{
			{
				Id:                 id,
				Serial:             "serial" + postfix,
				Name:               "name" + postfix,
				ModelNumber:        *config_api_client.NewNullableString(&modelNumber),
				WifiMacAddress:     *config_api_client.NewNullableString(&wifiMacAddress),
				EthernetMacAddress: *config_api_client.NewNullableString(&ethernetMacAddress),
				Notes:              *config_api_client.NewNullableString(&notes),
				PcapMode:           *config_api_client.NewNullableString(&pcapMode),
				Type:               shared.AgentType,
			},
		},
		Count: 1,
		Next:  *config_api_client.NewNullableString(nil),
	}
}

func MockGetAgent(id string, response interface{}, times int) {
	gock.New(MockUXIURL).
		Get(shared.AgentPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteAgent(id string, times int) {
	gock.New(MockUXIURL).
		Delete(shared.AgentPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}

func MockUpdateAgent(
	id string,
	request config_api_client.AgentsPatchRequest,
	response config_api_client.AgentsPatchResponse,
	times int,
) {
	gock.New(MockUXIURL).
		Patch(shared.AgentPath+"/"+id).
		MatchHeader("Content-Type", "application/merge-patch+json").
		MatchHeader("Authorization", mockToken).
		JSON(request).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}
