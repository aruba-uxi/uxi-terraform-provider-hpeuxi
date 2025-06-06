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

func GenerateAgentPatchRequest(postfix string) config_api_client.AgentPatchRequest {
	name := "name" + postfix
	notes := "notes" + postfix
	pcapMode := config_api_client.AGENTPCAPMODE_LIGHT

	return config_api_client.AgentPatchRequest{
		Name:     &name,
		Notes:    &notes,
		PcapMode: &pcapMode,
	}
}

func GenerateAgentPatchResponse(id string, postfix string) config_api_client.AgentPatchResponse {
	modelNumber := "model_number" + postfix
	wifiMacAddress := "wifi_mac_address" + postfix
	ethernetMacAddress := "ethernet_mac_address" + postfix
	notes := "notes" + postfix
	pcapMode := config_api_client.AGENTPCAPMODE_LIGHT
	groupPath := "group_path" + postfix
	groupName := "group_name" + postfix

	return config_api_client.AgentPatchResponse{
		Id:                 "id",
		Serial:             "serial" + postfix,
		Name:               "name" + postfix,
		ModelNumber:        *config_api_client.NewNullableString(&modelNumber),
		WifiMacAddress:     *config_api_client.NewNullableString(&wifiMacAddress),
		EthernetMacAddress: *config_api_client.NewNullableString(&ethernetMacAddress),
		Notes:              *config_api_client.NewNullableString(&notes),
		PcapMode:           *config_api_client.NewNullableAgentPcapMode(&pcapMode),
		GroupPath:          *config_api_client.NewNullableString(&groupPath),
		GroupName:          *config_api_client.NewNullableString(&groupName),
		Type:               shared.AgentType,
	}
}

func GenerateAgentsGetResponse(id string, postfix string) config_api_client.AgentsGetResponse {
	modelNumber := "model_number" + postfix
	wifiMacAddress := "wifi_mac_address" + postfix
	ethernetMacAddress := "ethernet_mac_address" + postfix
	notes := "notes" + postfix
	pcapMode := config_api_client.AGENTPCAPMODE_LIGHT
	groupPath := "group_path" + postfix
	groupName := "group_name" + postfix

	return config_api_client.AgentsGetResponse{
		Items: []config_api_client.AgentsGetItem{
			{
				Id:                 id,
				Serial:             "serial" + postfix,
				Name:               "name" + postfix,
				ModelNumber:        *config_api_client.NewNullableString(&modelNumber),
				WifiMacAddress:     *config_api_client.NewNullableString(&wifiMacAddress),
				EthernetMacAddress: *config_api_client.NewNullableString(&ethernetMacAddress),
				Notes:              *config_api_client.NewNullableString(&notes),
				PcapMode:           *config_api_client.NewNullableAgentPcapMode(&pcapMode),
				GroupPath:          *config_api_client.NewNullableString(&groupPath),
				GroupName:          *config_api_client.NewNullableString(&groupName),
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
	request config_api_client.AgentPatchRequest,
	response config_api_client.AgentPatchResponse,
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
