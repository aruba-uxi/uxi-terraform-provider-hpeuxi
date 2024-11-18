/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"encoding/json"
	"net/http"

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
		"type":               "networking-uxi/sensor",
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
		"type":               "networking-uxi/sensor",
	}
}

func GenerateNonRootGroupResponse(
	id string,
	nonReplacementFieldPostfix string,
	replacementFieldPostfix string,
) map[string]interface{} {
	parentId := "parent_id" + replacementFieldPostfix

	return map[string]interface{}{
		"id":     id,
		"name":   "name" + nonReplacementFieldPostfix,
		"parent": map[string]string{"id": parentId},
		"path":   parentId + "." + id,
		"type":   "networking-uxi/group",
	}
}

func GenerateGroupRequest(
	id string,
	nonReplacementFieldPostfix string,
	replacementFieldPostfix string,
) map[string]interface{} {
	return map[string]interface{}{
		"name":     "name" + nonReplacementFieldPostfix,
		"parentId": "parent_id" + replacementFieldPostfix,
	}
}

func GeneratePaginatedResponse(items []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"items": items,
		"next":  nil,
		"count": len(items),
	}
}

func GenerateServiceTestResponse(
	id string,
	postfix string,
) map[string]interface{} {
	return map[string]interface{}{
		"id":        id,
		"category":  "external" + postfix,
		"name":      "name" + postfix,
		"target":    "target" + postfix,
		"template":  "template" + postfix,
		"isEnabled": true,
		"type":      "networking-uxi/service-test",
	}
}

func GenerateWiredNetworkResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":                   id,
		"name":                 "name" + postfix,
		"createdAt":            "2024-09-11T12:00:00.000Z",
		"updatedAt":            "2024-09-11T12:00:00.000Z",
		"ipVersion":            "ip_version" + postfix,
		"security":             "security" + postfix,
		"dnsLookupDomain":      "dns_lookup_domain" + postfix,
		"disableEdns":          false,
		"useDns64":             false,
		"externalConnectivity": false,
		"vLanId":               123,
		"type":                 "networking-uxi/wired-network",
	}
}

func GenerateWirelessNetworkResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":                   id,
		"ssid":                 "ssid" + postfix,
		"createdAt":            "2024-09-11T12:00:00.000Z",
		"updatedAt":            "2024-09-11T12:00:00.000Z",
		"name":                 "name" + postfix,
		"ipVersion":            "ip_version" + postfix,
		"security":             "security" + postfix,
		"hidden":               false,
		"bandLocking":          "band_locking" + postfix,
		"dnsLookupDomain":      "dns_lookup_domain" + postfix,
		"disableEdns":          false,
		"useDns64":             false,
		"externalConnectivity": false,
		"type":                 "networking-uxi/wireless-network",
	}
}

func GenerateSensorGroupAssignmentResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":     id,
		"group":  map[string]string{"id": "group_id" + postfix},
		"sensor": map[string]string{"id": "sensor_id" + postfix},
		"type":   "networking-uxi/sensor-group-assignment",
	}
}

func GenerateSensorGroupAssignmentRequest(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"groupId":  "group_id" + postfix,
		"sensorId": "sensor_id" + postfix,
	}
}

func GenerateAgentGroupAssignmentRequest(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"groupId": "group_id" + postfix,
		"agentId": "agent_id" + postfix,
	}
}

func GenerateAgentGroupAssignmentResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":    id,
		"group": map[string]string{"id": "group_id" + postfix},
		"agent": map[string]string{"id": "agent_id" + postfix},
		"type":  "networking-uxi/agent-group-assignment",
	}
}

func GenerateNetworkGroupAssignmentResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":      id,
		"group":   map[string]string{"id": "group_id" + postfix},
		"network": map[string]string{"id": "network_id" + postfix},
		"type":    "networking-uxi/network-group-assignment",
	}
}

func GenerateNetworkGroupAssignmentRequest(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"groupId":   "group_id" + postfix,
		"networkId": "network_id" + postfix,
	}
}

func GenerateServiceTestGroupAssignmentResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":          id,
		"group":       map[string]string{"id": "group_id" + postfix},
		"serviceTest": map[string]string{"id": "service_test_id" + postfix},
		"type":        "networking-uxi/service-test-group-assignment",
	}
}

func GenerateServiceTestGroupAssignmentRequest(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"groupId":       "group_id" + postfix,
		"serviceTestId": "service_test_id" + postfix,
	}
}

func MockOAuth() *gock.Response {
	return gock.New("https://test.sso.common.cloud.hpe.com").
		Post("/as/token.oauth2").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		Persist().
		Reply(http.StatusOK).
		JSON(map[string]interface{}{
			"access_token": "mock_token",
			"token_type":   "bearer",
			"expires_in":   3600,
		})

}

func MockGetAgent(id string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/networking-uxi/v1alpha1/agents").
		MatchHeader("Authorization", "mock_token").
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteAgent(id string, times int) {
	gock.New("https://test.api.capenetworks.com").
		Delete("/networking-uxi/v1alpha1/agents/"+id).
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(http.StatusNoContent)
}

func MockUpdateAgent(
	id string,
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	gock.New("https://test.api.capenetworks.com").
		Patch("/networking-uxi/v1alpha1/agents/"+id).
		MatchHeader("Content-Type", "application/merge-patch+json").
		MatchHeader("Authorization", "mock_token").
		JSON(request).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostGroup(request map[string]interface{}, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Post("/networking-uxi/v1alpha1/groups").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", "mock_token").
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockGetGroup(id string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/networking-uxi/v1alpha1/groups").
		MatchHeader("Authorization", "mock_token").
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockUpdateGroup(
	id string,
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	body, _ := json.Marshal(request)
	gock.New("https://test.api.capenetworks.com").
		Patch("/networking-uxi/v1alpha1/groups/"+id).
		MatchHeader("Authorization", "mock_token").
		MatchHeader("Content-Type", "application/merge-patch+json").
		BodyString(string(body)).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteGroup(id string, times int) {
	gock.New("https://test.api.capenetworks.com").
		Delete("/networking-uxi/v1alpha1/groups/"+id).
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(http.StatusNoContent)
}

func MockGetSensor(id string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/networking-uxi/v1alpha1/sensors").
		MatchHeader("Authorization", "mock_token").
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
	gock.New("https://test.api.capenetworks.com").
		Patch("/networking-uxi/v1alpha1/sensors/"+id).
		MatchHeader("Content-Type", "application/merge-patch+json").
		MatchHeader("Authorization", "mock_token").
		JSON(request).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockGetWiredNetwork(id string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/networking-uxi/v1alpha1/wired-networks").
		MatchHeader("Authorization", "mock_token").
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockGetWirelessNetwork(id string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/networking-uxi/v1alpha1/wireless-networks").
		MatchHeader("Authorization", "mock_token").
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockGetServiceTest(id string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/networking-uxi/v1alpha1/service-tests").
		MatchHeader("Authorization", "mock_token").
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockGetAgentGroupAssignment(id string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/networking-uxi/v1alpha1/agent-group-assignments").
		MatchHeader("Authorization", "mock_token").
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostAgentGroupAssignment(
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	gock.New("https://test.api.capenetworks.com").
		Post("/networking-uxi/v1alpha1/agent-group-assignments").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", "mock_token").
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteAgentGroupAssignment(id string, times int) {
	gock.New("https://test.api.capenetworks.com").
		Delete("/networking-uxi/v1alpha1/agent-group-assignments/"+id).
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(http.StatusNoContent)
}

func MockGetSensorGroupAssignment(id string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/networking-uxi/v1alpha1/sensor-group-assignments").
		MatchHeader("Authorization", "mock_token").
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostSensorGroupAssignment(
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	gock.New("https://test.api.capenetworks.com").
		Post("/networking-uxi/v1alpha1/sensor-group-assignments").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", "mock_token").
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteSensorGroupAssignment(id string, times int) {
	gock.New("https://test.api.capenetworks.com").
		Delete("/networking-uxi/v1alpha1/sensor-group-assignments/"+id).
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(http.StatusNoContent)
}

func MockGetNetworkGroupAssignment(id string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/networking-uxi/v1alpha1/network-group-assignments").
		MatchHeader("Authorization", "mock_token").
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostNetworkGroupAssignment(
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	gock.New("https://test.api.capenetworks.com").
		Post("/networking-uxi/v1alpha1/network-group-assignments").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", "mock_token").
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteNetworkGroupAssignment(id string, times int) {
	gock.New("https://test.api.capenetworks.com").
		Delete("/networking-uxi/v1alpha1/network-group-assignments/"+id).
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(http.StatusNoContent)
}

func MockGetServiceTestGroupAssignment(id string, response map[string]interface{}, times int) {
	gock.New("https://test.api.capenetworks.com").
		Get("/networking-uxi/v1alpha1/service-test-group-assignments").
		MatchHeader("Authorization", "mock_token").
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostServiceTestGroupAssignment(
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	gock.New("https://test.api.capenetworks.com").
		Post("/networking-uxi/v1alpha1/service-test-group-assignments").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", "mock_token").
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteServiceTestGroupAssignment(id string, times int) {
	gock.New("https://test.api.capenetworks.com").
		Delete("/networking-uxi/v1alpha1/service-test-group-assignments/"+id).
		MatchHeader("Authorization", "mock_token").
		Times(times).
		Reply(http.StatusNoContent)
}

var RateLimitingHeaders = map[string]string{
	"X-RateLimit-Limit":     "100",
	"X-RateLimit-Remaining": "0",
	"X-RateLimit-Reset":     "0.01",
}
