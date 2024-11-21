/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateSensorResponse(id string, postfix string) config_api_client.SensorsResponse {
	return config_api_client.SensorsResponse{
		Items: []config_api_client.SensorItem{
			{
				Id:                 id,
				Serial:             "serial" + postfix,
				Name:               "name" + postfix,
				ModelNumber:        "model_number" + postfix,
				WifiMacAddress:     *config_api_client.NewNullableString(config_api_client.PtrString("wifi_mac_address" + postfix)),
				EthernetMacAddress: *config_api_client.NewNullableString(config_api_client.PtrString("ethernet_mac_address" + postfix)),
				AddressNote:        *config_api_client.NewNullableString(config_api_client.PtrString("address_note" + postfix)),
				Longitude:          *config_api_client.NewNullableFloat32(config_api_client.PtrFloat32(0.0)),
				Latitude:           *config_api_client.NewNullableFloat32(config_api_client.PtrFloat32(0.0)),
				Notes:              *config_api_client.NewNullableString(config_api_client.PtrString("notes" + postfix)),
				PcapMode:           *config_api_client.NewNullableString(config_api_client.PtrString("light")),
				Type:               shared.SensorType,
			},
		},
		Count: 1,
		Next:  *config_api_client.NewNullableString(nil),
	}
}

func GenerateSensorPatchRequest(postfix string) config_api_client.SensorsPatchRequest {
	pcapMode, _ := config_api_client.NewPcapModeFromValue("light")
	return config_api_client.SensorsPatchRequest{
		Name:        config_api_client.PtrString("name" + postfix),
		AddressNote: config_api_client.PtrString("address_note" + postfix),
		Notes:       config_api_client.PtrString("notes" + postfix),
		PcapMode:    pcapMode,
	}
}

func GenerateSensorPatchResponse(id string, postfix string) config_api_client.SensorsPatchResponse {
	pcapMode, _ := config_api_client.NewPcapModeFromValue("light")
	return config_api_client.SensorsPatchResponse{
		Id:                 id,
		Serial:             "serial" + postfix,
		Name:               "name" + postfix,
		ModelNumber:        "model_number" + postfix,
		WifiMacAddress:     *config_api_client.NewNullableString(config_api_client.PtrString("wifi_mac_address" + postfix)),
		EthernetMacAddress: *config_api_client.NewNullableString(config_api_client.PtrString("ethernet_mac_address" + postfix)),
		AddressNote:        *config_api_client.NewNullableString(config_api_client.PtrString("address_note" + postfix)),
		Longitude:          *config_api_client.NewNullableFloat32(config_api_client.PtrFloat32(0.0)),
		Latitude:           *config_api_client.NewNullableFloat32(config_api_client.PtrFloat32(0.0)),
		Notes:              *config_api_client.NewNullableString(config_api_client.PtrString("notes" + postfix)),
		PcapMode:           *config_api_client.NewNullablePcapMode(pcapMode),
		Type:               shared.SensorType,
	}
}

func MockGetSensor(id string, response interface{}, times int) {
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
	request config_api_client.SensorsPatchRequest,
	response config_api_client.SensorsPatchResponse,
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
