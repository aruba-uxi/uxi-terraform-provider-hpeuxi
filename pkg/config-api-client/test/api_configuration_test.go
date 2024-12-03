/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package config_api_client

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func TestConfigurationAPI(t *testing.T) {
	configuration := config_api_client.NewConfiguration()
	configuration.Host = "localhost:80"
	configuration.Scheme = "http"
	apiClient := config_api_client.NewAPIClient(configuration)

	defer gock.Off()

	t.Run("Test ConfigurationAPI AgentsGet", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/agents").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":                 "uid",
						"serial":             "serial",
						"name":               "name",
						"modelNumber":        "model_number",
						"wifiMacAddress":     "wifi_mac_address",
						"ethernetMacAddress": "ethernet_mac_address",
						"notes":              "notes",
						"pcapMode":           "light",
						"type":               "networking-uxi/sensor",
					},
				},
				"next":  nil,
				"count": 1,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			AgentsGet(context.Background()).
			Id("uid").
			Limit(10).
			Next("some-cursor").
			Execute()
		modelNumber := "model_number"

		wifiMacAddress := "wifi_mac_address"
		ethernetMacAddress := "ethernet_mac_address"
		notes := "notes"
		pcapMode := config_api_client.AGENTPCAPMODE_LIGHT
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.AgentsGetResponse{
			Items: []config_api_client.AgentsGetItem{
				{
					Id:                 "uid",
					Serial:             "serial",
					Name:               "name",
					ModelNumber:        *config_api_client.NewNullableString(&modelNumber),
					WifiMacAddress:     *config_api_client.NewNullableString(&wifiMacAddress),
					EthernetMacAddress: *config_api_client.NewNullableString(&ethernetMacAddress),
					Notes:              *config_api_client.NewNullableString(&notes),
					PcapMode:           *config_api_client.NewNullableAgentPcapMode(&pcapMode),
					Type:               "networking-uxi/sensor",
				},
			},
			Next:  *config_api_client.NewNullableString(nil),
			Count: 1,
		})
	})

	t.Run("Test ConfigurationAPI AgentPatchRequest", func(t *testing.T) {
		gock.New(configuration.Scheme+"://"+configuration.Host).
			Patch("/networking-uxi/v1alpha1/agents/uid").
			MatchHeader("Content-Type", "application/merge-patch+json").
			JSON(map[string]interface{}{"name": "new_name", "notes": "new_notes", "pcapMode": "off"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"id":                 "uid",
				"serial":             "serial",
				"name":               "new_name",
				"modelNumber":        "model_number",
				"wifiMacAddress":     "wifi_mac_address",
				"ethernetMacAddress": "ethernet_mac_address",
				"notes":              "new_notes",
				"pcapMode":           "off",
				"type":               "networking-uxi/agent",
			},
			)
		name := "new_name"
		notes := "new_notes"
		pcapMode := config_api_client.AGENTPCAPMODE_OFF
		agentsPatchRequest := config_api_client.AgentPatchRequest{
			Name:     &name,
			Notes:    &notes,
			PcapMode: &pcapMode,
		}
		resp, httpRes, err := apiClient.ConfigurationAPI.
			AgentPatch(context.Background(), "uid").
			AgentPatchRequest(agentsPatchRequest).
			Execute()
		wifiMacAddress := "wifi_mac_address"

		ethernetMacAddress := "ethernet_mac_address"
		modelNumber := "model_number"
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.AgentPatchResponse{
			Id:                 "uid",
			Serial:             "serial",
			Name:               "new_name",
			ModelNumber:        *config_api_client.NewNullableString(&modelNumber),
			WifiMacAddress:     *config_api_client.NewNullableString(&wifiMacAddress),
			EthernetMacAddress: *config_api_client.NewNullableString(&ethernetMacAddress),
			Notes:              *config_api_client.NewNullableString(&notes),
			PcapMode:           *config_api_client.NewNullableAgentPcapMode(&pcapMode),
			Type:               "networking-uxi/agent",
		})
	})

	t.Run("Test ConfigurationAPI AgentDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/agents/uid").
			Reply(http.StatusNoContent)

		_, httpRes, err := apiClient.ConfigurationAPI.
			AgentDelete(context.Background(), "uid").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusNoContent, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPI GroupsGet", func(t *testing.T) {
		parent_uid := "parent_uid"
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/groups").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":     "uid",
						"name":   "name",
						"parent": map[string]string{"id": parent_uid},
						"path":   "root_uid.parent_uid.uid",
						"type":   "networking-uxi/group",
					},
				},
				"next":  nil,
				"count": 1,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GroupsGet(context.Background()).
			Id("uid").
			Limit(10).
			Next("some-cursor").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.GroupsGetResponse{
			Items: []config_api_client.GroupsGetItem{
				{
					Id:     "uid",
					Name:   "name",
					Parent: *config_api_client.NewNullableGroupsGetParent(config_api_client.NewGroupsGetParent("parent_uid")),
					Path:   "root_uid.parent_uid.uid",
					Type:   "networking-uxi/group",
				},
			},
			Next:  *config_api_client.NewNullableString(nil),
			Count: 1,
		})
	})

	t.Run("Test ConfigurationAPI GroupPost", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/networking-uxi/v1alpha1/groups").
			JSON(map[string]interface{}{
				"name":     "name",
				"parentId": "parent.uid",
			}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"id":     "node",
				"name":   "name",
				"parent": map[string]string{"id": "parent.uid"},
				"path":   "parent.uid.node",
				"type":   "networking-uxi/group",
			})
		groupsPostRequest := config_api_client.NewGroupPostRequest("name")
		groupsPostRequest.SetParentId("parent.uid")
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GroupPost(context.Background()).
			GroupPostRequest(*groupsPostRequest).Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.GroupPostResponse{
			Id:     "node",
			Name:   "name",
			Parent: *config_api_client.NewGroupPostParent("parent.uid"),
			Path:   "parent.uid.node",
			Type:   "networking-uxi/group",
		})
	})

	t.Run("Test ConfigurationAPI GroupPatch", func(t *testing.T) {
		gock.New(configuration.Scheme+"://"+configuration.Host).
			Patch("/networking-uxi/v1alpha1/groups/node").
			MatchHeader("Content-Type", "application/merge-patch+json").
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"id":     "node",
				"name":   "new_name",
				"parent": map[string]string{"id": "parent.uid"},
				"path":   "parent.uid.node",
				"type":   "networking-uxi/group",
			})
		name := "new_name"
		groupsPatchRequest := config_api_client.NewGroupPatchRequest()
		groupsPatchRequest.Name = &name
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GroupPatch(context.Background(), "node").
			GroupPatchRequest(*groupsPatchRequest).Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.GroupPatchResponse{
			Id:     "node",
			Name:   "new_name",
			Parent: *config_api_client.NewGroupPatchParent("parent.uid"),
			Path:   "parent.uid.node",
			Type:   "networking-uxi/group",
		})
	})

	t.Run("Test ConfigurationAPI GroupDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/groups/uid").
			Reply(http.StatusNoContent)

		_, httpRes, err := apiClient.ConfigurationAPI.
			GroupDelete(context.Background(), "uid").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusNoContent, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPI SensorsGet", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/sensors").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":                 "uid",
						"serial":             "serial",
						"name":               "name",
						"modelNumber":        "model_number",
						"wifiMacAddress":     "wifi_mac_address",
						"ethernetMacAddress": "ethernet_mac_address",
						"addressNote":        "address_note",
						"longitude":          0.0,
						"latitude":           0.0,
						"notes":              "notes",
						"pcapMode":           "light",
						"type":               "networking-uxi/sensor",
					},
				},
				"count": 1,
				"next":  nil,
			})

		resp, httpRes, err := apiClient.ConfigurationAPI.
			SensorsGet(context.Background()).
			Id("uid").
			Limit(10).
			Next("some-cursor").
			Execute()
		WifiMacAddress := "wifi_mac_address"

		EthernetMacAddress := "ethernet_mac_address"
		AddressNote := "address_note"
		var Longitude float32 = 0.0
		var Latitude float32 = 0.0
		Notes := "notes"
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.SensorsGetResponse{
			Items: []config_api_client.SensorsGetItem{
				{
					Id:                 "uid",
					Serial:             "serial",
					Name:               "name",
					ModelNumber:        "model_number",
					WifiMacAddress:     *config_api_client.NewNullableString(&WifiMacAddress),
					EthernetMacAddress: *config_api_client.NewNullableString(&EthernetMacAddress),
					AddressNote:        *config_api_client.NewNullableString(&AddressNote),
					Longitude:          *config_api_client.NewNullableFloat32(&Longitude),
					Latitude:           *config_api_client.NewNullableFloat32(&Latitude),
					Notes:              *config_api_client.NewNullableString(&Notes),
					PcapMode:           *config_api_client.NewNullableSensorPcapMode(config_api_client.SENSORPCAPMODE_LIGHT.Ptr()),
					Type:               "networking-uxi/sensor",
				},
			},
			Next:  *config_api_client.NewNullableString(nil),
			Count: 1,
		})
	})

	t.Run("Test ConfigurationAPI SensorPatchRequest", func(t *testing.T) {
		gock.New(configuration.Scheme+"://"+configuration.Host).
			Patch("/networking-uxi/v1alpha1/sensors/uid").
			MatchHeader("Content-Type", "application/merge-patch+json").
			JSON(map[string]interface{}{"name": "new_name", "addressNote": "new_address_note", "notes": "new_notes", "pcapMode": "off"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"id":                 "uid",
				"serial":             "serial",
				"name":               "new_name",
				"modelNumber":        "model_number",
				"wifiMacAddress":     "wifi_mac_address",
				"ethernetMacAddress": "ethernet_mac_address",
				"addressNote":        "new_address_note",
				"longitude":          0.0,
				"latitude":           0.0,
				"notes":              "new_notes",
				"pcapMode":           "off",
				"type":               "networking-uxi/sensor",
			},
			)
		name := "new_name"
		addressNote := "new_address_note"
		notes := "new_notes"
		sensorsPatchRequest := config_api_client.SensorPatchRequest{
			Name:        &name,
			AddressNote: &addressNote,
			Notes:       &notes,
			PcapMode:    config_api_client.SENSORPCAPMODE_OFF.Ptr(),
		}
		resp, httpRes, err := apiClient.ConfigurationAPI.
			SensorPatch(context.Background(), "uid").
			SensorPatchRequest(sensorsPatchRequest).
			Execute()
		wifiMacAddress := "wifi_mac_address"

		ethernetMacAddress := "ethernet_mac_address"
		var longitude float32 = 0.0
		var latitude float32 = 0.0
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.SensorPatchResponse{
			Id:                 "uid",
			Serial:             "serial",
			Name:               "new_name",
			ModelNumber:        "model_number",
			WifiMacAddress:     *config_api_client.NewNullableString(&wifiMacAddress),
			EthernetMacAddress: *config_api_client.NewNullableString(&ethernetMacAddress),
			AddressNote:        *config_api_client.NewNullableString(&addressNote),
			Longitude:          *config_api_client.NewNullableFloat32(&longitude),
			Latitude:           *config_api_client.NewNullableFloat32(&latitude),
			Notes:              *config_api_client.NewNullableString(&notes),
			PcapMode:           *config_api_client.NewNullableSensorPcapMode(config_api_client.SENSORPCAPMODE_OFF.Ptr()),
			Type:               "networking-uxi/sensor",
		})
	})

	t.Run("Test ConfigurationAPI AgentGroupAssignmentsGet", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/agent-group-assignments").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":    "uid",
						"group": map[string]string{"id": "group_uid"},
						"agent": map[string]string{"id": "agent_uid"},
						"type":  "networking-uxi/agent-group-assignment",
					},
				},
				"count": 1,
				"next":  nil,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			AgentGroupAssignmentsGet(context.Background()).
			Id("uid").
			Limit(10).
			Next("some-cursor").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.AgentGroupAssignmentsGetResponse{
			Items: []config_api_client.AgentGroupAssignmentsGetItem{
				{
					Id:    "uid",
					Group: *config_api_client.NewAgentGroupAssignmentsGetGroup("group_uid"),
					Agent: *config_api_client.NewAgentGroupAssignmentsGetAgent("agent_uid"),
					Type:  "networking-uxi/agent-group-assignment",
				},
			},
			Count: 1,
			Next:  *config_api_client.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPI AgentGroupAssignmentPost", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/networking-uxi/v1alpha1/agent-group-assignments").
			JSON(map[string]interface{}{
				"groupId": "group_uid",
				"agentId": "agent_uid",
			}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"id":    "uid",
				"group": map[string]string{"id": "group_uid"},
				"agent": map[string]string{"id": "agent_uid"},
				"type":  "networking-uxi/agent-group-assignment",
			})

		postRequest := config_api_client.NewAgentGroupAssignmentPostRequest(
			"group_uid",
			"agent_uid",
		)
		resp, httpRes, err := apiClient.ConfigurationAPI.
			AgentGroupAssignmentPost(context.Background()).
			AgentGroupAssignmentPostRequest(*postRequest).
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.AgentGroupAssignmentPostResponse{
			Id:    "uid",
			Group: *config_api_client.NewAgentGroupAssignmentPostGroup("group_uid"),
			Agent: *config_api_client.NewAgentGroupAssignmentPostAgent("agent_uid"),
			Type:  "networking-uxi/agent-group-assignment",
		})
	})

	t.Run("Test ConfigurationAPI AgentGroupAssignmentDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/agent-group-assignments/uid").
			Reply(http.StatusNoContent)

		_, httpRes, err := apiClient.ConfigurationAPI.
			AgentGroupAssignmentDelete(context.Background(), "uid").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusNoContent, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPI SensorGroupAssignmentsGet", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/sensor-group-assignments").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":     "uid",
						"group":  map[string]string{"id": "group_uid"},
						"sensor": map[string]string{"id": "sensor_uid"},
						"type":   "networking-uxi/sensor-group-assignment",
					},
				},
				"count": 1,
				"next":  nil,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			SensorGroupAssignmentsGet(context.Background()).
			Id("uid").
			Limit(10).
			Next("some-cursor").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.SensorGroupAssignmentsGetResponse{
			Items: []config_api_client.SensorGroupAssignmentsGetItem{
				{
					Id:     "uid",
					Group:  *config_api_client.NewSensorGroupAssignmentsGetGroup("group_uid"),
					Sensor: *config_api_client.NewSensorGroupAssignmentsGetSensor("sensor_uid"),
					Type:   "networking-uxi/sensor-group-assignment",
				},
			},
			Count: 1,
			Next:  *config_api_client.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPI SensorGroupAssignmentPost", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/networking-uxi/v1alpha1/sensor-group-assignments").
			JSON(map[string]interface{}{
				"groupId":  "group_uid",
				"sensorId": "sensor_uid",
			}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"id":     "uid",
				"group":  map[string]string{"id": "group_uid"},
				"sensor": map[string]string{"id": "sensor_uid"},
				"type":   "networking-uxi/sensor-group-assignment",
			})

		postRequest := config_api_client.NewSensorGroupAssignmentPostRequest(
			"group_uid",
			"sensor_uid",
		)
		resp, httpRes, err := apiClient.ConfigurationAPI.
			SensorGroupAssignmentPost(context.Background()).
			SensorGroupAssignmentPostRequest(*postRequest).
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.SensorGroupAssignmentPostResponse{
			Id:     "uid",
			Group:  *config_api_client.NewSensorGroupAssignmentPostGroup("group_uid"),
			Sensor: *config_api_client.NewSensorGroupAssignmentPostSensor("sensor_uid"),
			Type:   "networking-uxi/sensor-group-assignment",
		})
	})

	t.Run("Test ConfigurationAPI SensorGroupAssignmentDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/sensor-group-assignments/uid").
			Reply(http.StatusNoContent)

		_, httpRes, err := apiClient.ConfigurationAPI.
			SensorGroupAssignmentDelete(context.Background(), "uid").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusNoContent, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPI WiredNetworksGet", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/wired-networks").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":                   "uid",
						"name":                 "alias",
						"createdAt":            "2024-09-11T12:00:00.000Z",
						"updatedAt":            "2024-09-11T12:00:00.000Z",
						"ipVersion":            config_api_client.IPVERSION_IPV4,
						"security":             "security",
						"dnsLookupDomain":      "dns_lookup_domain",
						"disableEdns":          true,
						"useDns64":             false,
						"externalConnectivity": true,
						"vLanId":               1,
						"type":                 "networking-uxi/wired-network",
					},
				},
				"count": 1,
				"next":  nil,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			WiredNetworksGet(context.Background()).
			Id("uid").
			Limit(10).
			Next("some-cursor").
			Execute()
		security := "security"

		dnsLookupDomain := "dns_lookup_domain"
		var vlanId int32 = 1
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.WiredNetworksGetResponse{
			Items: []config_api_client.WiredNetworksGetItem{
				{
					Id:                   "uid",
					Name:                 "alias",
					IpVersion:            config_api_client.IPVERSION_IPV4,
					UpdatedAt:            time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					CreatedAt:            time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					Security:             *config_api_client.NewNullableString(&security),
					DnsLookupDomain:      *config_api_client.NewNullableString(&dnsLookupDomain),
					DisableEdns:          true,
					UseDns64:             false,
					ExternalConnectivity: true,
					VLanId:               *config_api_client.NewNullableInt32(&vlanId),
					Type:                 "networking-uxi/wired-network",
				},
			},
			Count: 1,
			Next:  *config_api_client.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPI WirelessNetworksGet", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/wireless-networks").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":                   "uid",
						"ssid":                 "ssid",
						"name":                 "alias",
						"createdAt":            "2024-09-11T12:00:00.000Z",
						"updatedAt":            "2024-09-11T12:00:00.000Z",
						"ipVersion":            config_api_client.IPVERSION_IPV4,
						"security":             "security",
						"hidden":               false,
						"bandLocking":          "band_locking",
						"dnsLookupDomain":      "dns_lookup_domain",
						"disableEdns":          true,
						"useDns64":             false,
						"externalConnectivity": true,
						"type":                 "networking-uxi/wireless-network",
					},
				},
				"count": 1,
				"next":  nil,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			WirelessNetworksGet(context.Background()).
			Id("uid").
			Limit(10).
			Next("some-cursor").
			Execute()
		security := "security"

		dnsLookupDomain := "dns_lookup_domain"
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.WirelessNetworksGetResponse{
			Items: []config_api_client.WirelessNetworksGetItem{
				{
					Id:                   "uid",
					Name:                 "alias",
					Ssid:                 "ssid",
					Security:             *config_api_client.NewNullableString(&security),
					IpVersion:            config_api_client.IPVERSION_IPV4,
					CreatedAt:            time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					UpdatedAt:            time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					Hidden:               false,
					BandLocking:          "band_locking",
					DnsLookupDomain:      *config_api_client.NewNullableString(&dnsLookupDomain),
					DisableEdns:          true,
					UseDns64:             false,
					ExternalConnectivity: true,
					Type:                 "networking-uxi/wireless-network",
				},
			},
			Count: 1,
			Next:  *config_api_client.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPI NetworkGroupAssignmentsGet", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/network-group-assignments").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":      "uid",
						"group":   map[string]string{"id": "group_uid"},
						"network": map[string]string{"id": "network_uid"},
						"type":    "networking-uxi/network-group-assignment",
					},
				},
				"count": 1,
				"next":  nil,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			NetworkGroupAssignmentsGet(context.Background()).
			Id("uid").
			Limit(10).
			Next("some-cursor").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.NetworkGroupAssignmentsGetResponse{
			Items: []config_api_client.NetworkGroupAssignmentsGetItem{
				{
					Id:      "uid",
					Group:   *config_api_client.NewNetworkGroupAssignmentsGetGroup("group_uid"),
					Network: *config_api_client.NewNetworkGroupAssignmentsGetNetwork("network_uid"),
					Type:    "networking-uxi/network-group-assignment",
				},
			},
			Count: 1,
			Next:  *config_api_client.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPI NetworkGroupAssignmentPost", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/networking-uxi/v1alpha1/network-group-assignments").
			JSON(map[string]interface{}{
				"groupId":   "group_uid",
				"networkId": "network_uid",
			}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"id":      "uid",
				"group":   map[string]string{"id": "group_uid"},
				"network": map[string]string{"id": "network_uid"},
				"type":    "networking-uxi/network-group-assignment",
			})

		postRequest := config_api_client.NewNetworkGroupAssignmentPostRequest(
			"group_uid",
			"network_uid",
		)
		resp, httpRes, err := apiClient.ConfigurationAPI.
			NetworkGroupAssignmentPost(context.Background()).
			NetworkGroupAssignmentPostRequest(*postRequest).
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.NetworkGroupAssignmentPostResponse{
			Id:      "uid",
			Group:   *config_api_client.NewNetworkGroupAssignmentPostGroup("group_uid"),
			Network: *config_api_client.NewNetworkGroupAssignmentPostNetwork("network_uid"),
			Type:    "networking-uxi/network-group-assignment",
		})
	})

	t.Run("Test ConfigurationAPI NetworkGroupAssignmentDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/network-group-assignments/uid").
			Reply(http.StatusNoContent)

		_, httpRes, err := apiClient.ConfigurationAPI.
			NetworkGroupAssignmentDelete(context.Background(), "uid").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusNoContent, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPI ServiceTestGroupAssignmentsGet", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/service-test-group-assignments").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":          "uid",
						"group":       map[string]string{"id": "group_uid"},
						"serviceTest": map[string]string{"id": "service_test_uid"},
						"type":        "networking-uxi/service-test-group-assignment",
					},
				},
				"count": 1,
				"next":  nil,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			ServiceTestGroupAssignmentsGet(context.Background()).
			Id("uid").
			Limit(10).
			Next("some-cursor").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.ServiceTestGroupAssignmentsGetResponse{
			Items: []config_api_client.ServiceTestGroupAssignmentsGetItem{
				{
					Id:          "uid",
					Group:       *config_api_client.NewServiceTestGroupAssignmentsGetGroup("group_uid"),
					ServiceTest: *config_api_client.NewServiceTestGroupAssignmentsGetServiceTest("service_test_uid"),
					Type:        "networking-uxi/service-test-group-assignment",
				},
			},
			Count: 1,
			Next:  *config_api_client.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPI ServiceTestGroupAssignmentPost", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/networking-uxi/v1alpha1/service-test-group-assignments").
			JSON(map[string]interface{}{
				"groupId":       "group_uid",
				"serviceTestId": "service_test_uid",
			}).
			Reply(http.StatusOK).
			JSON(map[string]interface{}{
				"id":          "uid",
				"group":       map[string]string{"id": "group_uid"},
				"serviceTest": map[string]string{"id": "service_test_uid"},
				"type":        "networking-uxi/service-test-group-assignment",
			})

		postRequest := config_api_client.NewServiceTestGroupAssignmentPostRequest(
			"group_uid",
			"service_test_uid",
		)
		resp, httpRes, err := apiClient.ConfigurationAPI.
			ServiceTestGroupAssignmentPost(context.Background()).
			ServiceTestGroupAssignmentPostRequest(*postRequest).
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.ServiceTestGroupAssignmentPostResponse{
			Id:          "uid",
			Group:       *config_api_client.NewServiceTestGroupAssignmentPostGroup("group_uid"),
			ServiceTest: *config_api_client.NewServiceTestGroupAssignmentPostServiceTest("service_test_uid"),
			Type:        "networking-uxi/service-test-group-assignment",
		})
	})

	t.Run("Test ConfigurationAPI ServiceTestGroupAssignmentDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/service-test-group-assignments/uid").
			Reply(http.StatusNoContent)

		_, httpRes, err := apiClient.ConfigurationAPI.
			ServiceTestGroupAssignmentDelete(context.Background(), "uid").
			Execute()
		require.Nil(t, err)

		defer httpRes.Body.Close()
		assert.Equal(t, http.StatusNoContent, httpRes.StatusCode)
	})
}
