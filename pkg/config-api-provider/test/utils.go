package test

import (
	"encoding/json"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/h2non/gock"
)

var mockPaginationResponse = map[string]interface{}{
	"limit":    10,
	"first":    nil,
	"next":     nil,
	"previous": nil,
	"last":     nil,
}

func GenerateSensorResponseModel(uid string, postfix string) resources.SensorResponseModel {
	return resources.SensorResponseModel{
		UID:                uid,
		Serial:             "serial" + postfix,
		Name:               "name" + postfix,
		ModelNumber:        "model_number" + postfix,
		WifiMacAddress:     "wifi_mac_address" + postfix,
		EthernetMacAddress: "ethernet_mac_address" + postfix,
		AddressNote:        "address_note" + postfix,
		Longitude:          "longitude" + postfix,
		Latitude:           "latitude" + postfix,
		Notes:              "notes" + postfix,
		PCapMode:           "light" + postfix,
	}
}

func GenerateAgentResponseModel(uid string, postfix string) resources.AgentResponseModel {
	return resources.AgentResponseModel{
		UID:                uid,
		Serial:             "serial" + postfix,
		Name:               "name" + postfix,
		ModelNumber:        "model_number" + postfix,
		WifiMacAddress:     "wifi_mac_address" + postfix,
		EthernetMacAddress: "ethernet_mac_address" + postfix,
		Notes:              "notes" + postfix,
		PCapMode:           "light" + postfix,
	}
}

func GenerateGroupResponseModel(uid string, non_replacement_field_postfix string, replacement_field_postfix string) resources.GroupResponseModel {
	parent_uid := "parent_uid" + replacement_field_postfix
	return resources.GroupResponseModel{
		UID:       uid,
		Name:      "name" + non_replacement_field_postfix,
		ParentUid: &parent_uid,
		Path:      parent_uid + "." + uid,
	}
}

func GenerateGroupPaginatedResponse(Groups []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"groups":     Groups,
		"pagination": mockPaginationResponse,
	}
}

func GenerateRootGroupResponseModel(uid string) resources.GroupResponseModel {
	return resources.GroupResponseModel{
		UID:       uid,
		Name:      "root",
		ParentUid: nil,
		Path:      uid,
	}
}

func GenerateServiceTestResponseModel(uid string, postfix string) resources.ServiceTestResponseModel {
	return resources.ServiceTestResponseModel{
		Uid:       uid,
		Category:  "external" + postfix,
		Title:     "title" + postfix,
		Target:    "target" + postfix,
		Template:  "template" + postfix,
		IsEnabled: true,
	}
}

func GenerateWiredNetworkResponseModel(uid string, postfix string) resources.WiredNetworkResponseModel {
	return resources.WiredNetworkResponseModel{
		Uid:                  uid,
		Alias:                "alias" + postfix,
		DatetimeCreated:      "datetime_created" + postfix,
		DatetimeUpdated:      "datetime_updated" + postfix,
		IpVersion:            "ip_version" + postfix,
		Security:             "security" + postfix,
		DnsLookupDomain:      "dns_lookup_domain" + postfix,
		DisableEdns:          false,
		UseDns64:             false,
		ExternalConnectivity: false,
		VlanId:               123,
	}
}

func GenerateWiredNetworkPaginatedResponse(WiredNetworks []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"wired_networks": WiredNetworks,
		"pagination":     mockPaginationResponse,
	}
}

func GenerateWirelessNetworkResponseModel(uid string, postfix string) resources.WirelessNetworkResponseModel {
	return resources.WirelessNetworkResponseModel{
		Uid:                  uid,
		Ssid:                 "ssid" + postfix,
		DatetimeCreated:      "datetime_created" + postfix,
		DatetimeUpdated:      "datetime_updated" + postfix,
		Alias:                "alias" + postfix,
		IpVersion:            "ip_version" + postfix,
		Security:             "security" + postfix,
		Hidden:               false,
		BandLocking:          "band_locking" + postfix,
		DnsLookupDomain:      "dns_lookup_domain" + postfix,
		DisableEdns:          false,
		UseDns64:             false,
		ExternalConnectivity: false,
	}
}

func GeneratePaginatedResponse(wirelessNetworks []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"wireless_networks": wirelessNetworks,
		"pagination":        mockPaginationResponse,
	}
}

func GenerateSensorGroupAssignmentResponse(uid string, postfix string) resources.SensorGroupAssignmentResponseModel {
	return resources.SensorGroupAssignmentResponseModel{
		UID:       uid,
		GroupUID:  "group_uid" + postfix,
		SensorUID: "sensor_uid" + postfix,
	}
}

func GenerateSensorGroupAssignmentPaginatedResponse(sensorGroupAssignments []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"sensor_group_assignments": sensorGroupAssignments,
		"pagination":               mockPaginationResponse,
	}
}

func GenerateAgentGroupAssignmentResponse(uid string, postfix string) resources.AgentGroupAssignmentResponseModel {
	return resources.AgentGroupAssignmentResponseModel{
		UID:      uid,
		GroupUID: "group_uid" + postfix,
		AgentUID: "agent_uid" + postfix,
	}
}

func GenerateNetworkGroupAssignmentResponse(uid string, postfix string) resources.NetworkGroupAssignmentResponseModel {
	return resources.NetworkGroupAssignmentResponseModel{
		UID:        uid,
		GroupUID:   "group_uid" + postfix,
		NetworkUID: "network_uid" + postfix,
	}
}

func GenerateServiceTestGroupAssignmentResponse(uid string, postfix string) resources.ServiceTestGroupAssignmentResponseModel {
	return resources.ServiceTestGroupAssignmentResponseModel{
		UID:            uid,
		GroupUID:       "group_uid" + postfix,
		ServiceTestUID: "service_test_uid" + postfix,
	}
}

// Converts a struct to a map while maintaining the json alias as keys
func StructToMap(obj interface{}) map[string]interface{} {
	data, _ := json.Marshal(obj) // Convert to a json string

	newMap := map[string]interface{}{}

	_ = json.Unmarshal(data, &newMap) // Convert to a map
	return newMap
}

func MockOAuth() {
	gock.New("https://test.sso.common.cloud.hpe.com").
		Post("/as/token.oauth2").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		Persist().
		Reply(200).
		JSON(map[string]interface{}{
			"access_token": "mock_token",
			"token_type":   "bearer",
			"expires_in":   3600,
		})

}

func MockPostGroup(response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Post("/configuration/app/v1/groups").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(200).
		JSON(response)
}

func MockGetGroup(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/configuration/app/v1/groups").
		MatchHeader("Authorization", "mock_token").
		MatchParam("uid", uid).
		Times(times).
		Reply(200).
		JSON(response)
}

func MockGetWiredNetwork(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/configuration/app/v1/wired-networks").
		MatchHeader("Authorization", "mock_token").
		MatchParam("uid", uid).
		Times(times).
		Reply(200).
		JSON(response)
}

func MockGetWirelessNetwork(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/configuration/app/v1/wireless-networks").
		MatchHeader("Authorization", "mock_token").
		MatchParam("uid", uid).
		Times(times).
		Reply(200).
		JSON(response)
}

func MockGetSensorGroupAssociation(uid string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/configuration/app/v1/sensor-group-assignments").
		MatchHeader("Authorization", "mock_token").
		MatchParam("uid", uid).
		Times(times).
		Reply(200).
		JSON(response)
}
