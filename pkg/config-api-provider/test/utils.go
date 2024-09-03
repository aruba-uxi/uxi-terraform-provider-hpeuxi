package test

import (
	"encoding/json"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/h2non/gock"
)

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

func GenerateSensorGroupAssignmentResponse(uid string, postfix string) resources.SensorGroupAssignmentResponseModel {
	return resources.SensorGroupAssignmentResponseModel{
		UID:       "sensor_group_assignment_uid" + postfix,
		GroupUID:  "group_uid" + postfix,
		SensorUID: "sensor_uid" + postfix,
	}
}

func GenerateAgentGroupAssignmentResponse(uid string, postfix string) resources.AgentGroupAssignmentResponseModel {
	return resources.AgentGroupAssignmentResponseModel{
		UID:      "agent_group_assignment_uid" + postfix,
		GroupUID: "group_uid" + postfix,
		AgentUID: "agent_uid" + postfix,
	}
}

func GenerateNetworkGroupAssignmentResponse(uid string, postfix string) resources.NetworkGroupAssignmentResponseModel {
	return resources.NetworkGroupAssignmentResponseModel{
		UID:        "network_group_assignment_uid" + postfix,
		GroupUID:   "group_uid" + postfix,
		NetworkUID: "network_uid" + postfix,
	}
}

func GenerateServiceTestGroupAssignmentResponse(uid string, postfix string) resources.ServiceTestGroupAssignmentResponseModel {
	return resources.ServiceTestGroupAssignmentResponseModel{
		UID:            "service_test_group_assignment_uid" + postfix,
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
	gock.New("https://sso.common.cloud.hpe.com").
		Post("/as/token.oauth2").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		Reply(200).
		JSON(map[string]interface{}{
			"access_token": "mock_token",
			"token_type":   "bearer",
			"expires_in":   3600,
		})
}

func MockPostGroup(response map[string]interface{}) {
	gock.New("https://test.api.capenetworks.com").
		Post("/configuration/app/v1/groups").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", "mock_token").
		Reply(200).
		JSON(response)

}
