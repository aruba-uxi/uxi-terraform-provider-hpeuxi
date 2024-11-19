/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateAgentUpdateRequest(postfix string) map[string]interface{} {
	return map[string]interface{}{
		"name":     "name" + postfix,
		"notes":    "notes" + postfix,
		"pcapMode": "light",
	}
}

func GenerateAgentResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":                 id,
		"serial":             "serial" + postfix,
		"name":               "name" + postfix,
		"modelNumber":        "model_number" + postfix,
		"wifiMacAddress":     "wifi_mac_address" + postfix,
		"ethernetMacAddress": "ethernet_mac_address" + postfix,
		"notes":              "notes" + postfix,
		"pcapMode":           "light",
		"type":               shared.AgentType,
	}
}

func MockGetAgent(id string, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Get(shared.AgentPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteAgent(id string, times int) {
	gock.New(MockUxiUrl).
		Delete(shared.AgentPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}

func MockUpdateAgent(
	id string,
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	gock.New(MockUxiUrl).
		Patch(shared.AgentPath+"/"+id).
		MatchHeader("Content-Type", "application/merge-patch+json").
		MatchHeader("Authorization", mockToken).
		JSON(request).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}
