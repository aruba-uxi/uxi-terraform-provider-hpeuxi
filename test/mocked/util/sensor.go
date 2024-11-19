/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateSensorResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":                 id,
		"serial":             "serial" + postfix,
		"name":               "name" + postfix,
		"modelNumber":        "model_number" + postfix,
		"wifiMacAddress":     "wifi_mac_address" + postfix,
		"ethernetMacAddress": "ethernet_mac_address" + postfix,
		"addressNote":        "address_note" + postfix,
		"longitude":          0.0,
		"latitude":           0.0,
		"notes":              "notes" + postfix,
		"pcapMode":           "light",
		"type":               shared.SensorType,
	}
}

func GenerateSensorUpdateRequest(postfix string) map[string]interface{} {
	return map[string]interface{}{
		"name":        "name" + postfix,
		"addressNote": "address_note" + postfix,
		"notes":       "notes" + postfix,
		"pcapMode":    "light",
	}
}

func MockGetSensor(id string, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Get(shared.SensorPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockUpdateSensor(
	id string,
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	gock.New(MockUxiUrl).
		Patch(shared.SensorPath+"/"+id).
		MatchHeader("Content-Type", "application/merge-patch+json").
		MatchHeader("Authorization", mockToken).
		JSON(request).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}
