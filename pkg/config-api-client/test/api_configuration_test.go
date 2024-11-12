package config_api_client

import (
	"context"
	"testing"
	"time"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			Reply(200).
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
						"pcapMode":           "pcap_mode",
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
		pcapMode := "pcap_mode"

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.AgentsResponse{
			Items: []config_api_client.AgentItem{
				{
					Id:                 "uid",
					Serial:             "serial",
					Name:               "name",
					ModelNumber:        *config_api_client.NewNullableString(&modelNumber),
					WifiMacAddress:     *config_api_client.NewNullableString(&wifiMacAddress),
					EthernetMacAddress: *config_api_client.NewNullableString(&ethernetMacAddress),
					Notes:              *config_api_client.NewNullableString(&notes),
					PcapMode:           *config_api_client.NewNullableString(&pcapMode),
					Type:               "networking-uxi/sensor",
				},
			},
			Next:  *config_api_client.NewNullableString(nil),
			Count: 1,
		})
	})

	t.Run("Test ConfigurationAPI AgentsPatchRequest", func(t *testing.T) {
		gock.New(configuration.Scheme+"://"+configuration.Host).
			Patch("/networking-uxi/v1alpha1/agents/uid").
			MatchHeader("Content-Type", "application/merge-patch+json").
			JSON(map[string]interface{}{"name": "new_name", "notes": "new_notes", "pcapMode": "off"}).
			Reply(200).
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
		pcapMode := "off"
		agentsPatchRequest := config_api_client.AgentsPatchRequest{
			Name:     &name,
			Notes:    &notes,
			PcapMode: &pcapMode,
		}
		resp, httpRes, err := apiClient.ConfigurationAPI.
			AgentsPatch(context.Background(), "uid").
			AgentsPatchRequest(agentsPatchRequest).
			Execute()

		wifiMacAddress := "wifi_mac_address"
		ethernetMacAddress := "ethernet_mac_address"
		modelNumber := "model_number"

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.AgentsPatchResponse{
			Id:                 "uid",
			Serial:             "serial",
			Name:               "new_name",
			ModelNumber:        *config_api_client.NewNullableString(&modelNumber),
			WifiMacAddress:     *config_api_client.NewNullableString(&wifiMacAddress),
			EthernetMacAddress: *config_api_client.NewNullableString(&ethernetMacAddress),
			Notes:              *config_api_client.NewNullableString(&notes),
			PcapMode:           *config_api_client.NewNullableString(&pcapMode),
			Type:               "networking-uxi/agent",
		})
	})

	t.Run("Test ConfigurationAPI AgentsDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/agents/uid").
			Reply(204)

		_, httpRes, err := apiClient.ConfigurationAPI.
			AgentsDelete(context.Background(), "uid").
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPI GroupsGet", func(t *testing.T) {
		parent_uid := "parent_uid"
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/groups").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(200).
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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.GroupsGetResponse{
			Items: []config_api_client.GroupsGetItem{
				{
					Id:     "uid",
					Name:   "name",
					Parent: *config_api_client.NewNullableParent(config_api_client.NewParent("parent_uid")),
					Path:   "root_uid.parent_uid.uid",
					Type:   "networking-uxi/group",
				},
			},
			Next:  *config_api_client.NewNullableString(nil),
			Count: 1,
		})
	})

	t.Run("Test ConfigurationAPI GroupsPost", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/networking-uxi/v1alpha1/groups").
			JSON(map[string]interface{}{
				"name":     "name",
				"parentId": "parent.uid",
			}).
			Reply(200).
			JSON(map[string]interface{}{
				"id":     "node",
				"name":   "name",
				"parent": map[string]string{"id": "parent.uid"},
				"path":   "parent.uid.node",
				"type":   "networking-uxi/group",
			})
		groupsPostRequest := config_api_client.NewGroupsPostRequest("name")
		groupsPostRequest.SetParentId("parent.uid")
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GroupsPost(context.Background()).
			GroupsPostRequest(*groupsPostRequest).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.GroupsPostResponse{
			Id:     "node",
			Name:   "name",
			Parent: *config_api_client.NewParent("parent.uid"),
			Path:   "parent.uid.node",
			Type:   "networking-uxi/group",
		})
	})

	t.Run("Test ConfigurationAPI GroupsPatch", func(t *testing.T) {
		gock.New(configuration.Scheme+"://"+configuration.Host).
			Patch("/networking-uxi/v1alpha1/groups/node").
			MatchHeader("Content-Type", "application/merge-patch+json").
			Reply(200).
			JSON(map[string]interface{}{
				"id":     "node",
				"name":   "new_name",
				"parent": map[string]string{"id": "parent.uid"},
				"path":   "parent.uid.node",
				"type":   "networking-uxi/group",
			})
		groupsPatchRequest := config_api_client.NewGroupsPatchRequest("new_name")
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GroupsPatch(context.Background(), "node").
			GroupsPatchRequest(*groupsPatchRequest).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.GroupsPatchResponse{
			Id:     "node",
			Name:   "new_name",
			Parent: *config_api_client.NewParent("parent.uid"),
			Path:   "parent.uid.node",
			Type:   "networking-uxi/group",
		})
	})

	t.Run("Test ConfigurationAPI GroupsDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/groups/uid").
			Reply(204)

		_, httpRes, err := apiClient.ConfigurationAPI.
			GroupsDelete(context.Background(), "uid").
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPI SensorsGet", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/sensors").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(200).
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
						"pcapMode":           "pcap_mode",
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
		PcapMode := "pcap_mode"

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.SensorsResponse{
			Items: []config_api_client.SensorItem{
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
					PcapMode:           *config_api_client.NewNullableString(&PcapMode),
					Type:               "networking-uxi/sensor",
				},
			},
			Next:  *config_api_client.NewNullableString(nil),
			Count: 1,
		})
	})

	t.Run("Test ConfigurationAPI SensorsPatchRequest", func(t *testing.T) {
		gock.New(configuration.Scheme+"://"+configuration.Host).
			Patch("/networking-uxi/v1alpha1/sensors/uid").
			MatchHeader("Content-Type", "application/merge-patch+json").
			JSON(map[string]interface{}{"name": "new_name", "addressNote": "new_address_note", "notes": "new_notes", "pcapMode": "off"}).
			Reply(200).
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
		pcapMode := "off"
		sensorsPatchRequest := config_api_client.SensorsPatchRequest{
			Name:        &name,
			AddressNote: &addressNote,
			Notes:       &notes,
			PcapMode:    &pcapMode,
		}
		resp, httpRes, err := apiClient.ConfigurationAPI.
			SensorsPatch(context.Background(), "uid").
			SensorsPatchRequest(sensorsPatchRequest).
			Execute()

		wifiMacAddress := "wifi_mac_address"
		ethernetMacAddress := "ethernet_mac_address"
		var longitude float32 = 0.0
		var latitude float32 = 0.0

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.SensorsPatchResponse{
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
			PcapMode:           *config_api_client.NewNullableString(&pcapMode),
			Type:               "networking-uxi/sensor",
		})
	})

	t.Run("Test ConfigurationAPI AgentGroupAssignmentsGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/agent-group-assignments").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(200).
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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.AgentGroupAssignmentsResponse{
			Items: []config_api_client.AgentGroupAssignmentsItem{
				{
					Id:    "uid",
					Group: *config_api_client.NewGroup("group_uid"),
					Agent: *config_api_client.NewAgent("agent_uid"),
					Type:  "networking-uxi/agent-group-assignment",
				},
			},
			Count: 1,
			Next:  *config_api_client.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPI AgentGroupAssignmentsPost", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/networking-uxi/v1alpha1/agent-group-assignments").
			JSON(map[string]interface{}{
				"groupId": "group_uid",
				"agentId": "agent_uid",
			}).
			Reply(200).
			JSON(map[string]interface{}{
				"id":    "uid",
				"group": map[string]string{"id": "group_uid"},
				"agent": map[string]string{"id": "agent_uid"},
				"type":  "networking-uxi/agent-group-assignment",
			})

		postRequest := config_api_client.NewAgentGroupAssignmentsPostRequest(
			"group_uid",
			"agent_uid",
		)
		resp, httpRes, err := apiClient.ConfigurationAPI.
			AgentGroupAssignmentsPost(context.Background()).
			AgentGroupAssignmentsPostRequest(*postRequest).
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.AgentGroupAssignmentResponse{
			Id:    "uid",
			Group: *config_api_client.NewGroup("group_uid"),
			Agent: *config_api_client.NewAgent("agent_uid"),
			Type:  "networking-uxi/agent-group-assignment",
		})
	})

	t.Run("Test ConfigurationAPI SensorGroupAssignmentsGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/sensor-group-assignments").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(200).
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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.SensorGroupAssignmentsResponse{
			Items: []config_api_client.SensorGroupAssignmentsItem{
				{
					Id:     "uid",
					Group:  *config_api_client.NewGroup("group_uid"),
					Sensor: *config_api_client.NewSensor("sensor_uid"),
					Type:   "networking-uxi/sensor-group-assignment",
				},
			},
			Count: 1,
			Next:  *config_api_client.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPI SensorGroupAssignmentsPost", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/networking-uxi/v1alpha1/sensor-group-assignments").
			JSON(map[string]interface{}{
				"groupId":  "group_uid",
				"sensorId": "sensor_uid",
			}).
			Reply(200).
			JSON(map[string]interface{}{
				"id":     "uid",
				"group":  map[string]string{"id": "group_uid"},
				"sensor": map[string]string{"id": "sensor_uid"},
				"type":   "networking-uxi/sensor-group-assignment",
			})

		postRequest := config_api_client.NewSensorGroupAssignmentsPostRequest(
			"group_uid",
			"sensor_uid",
		)
		resp, httpRes, err := apiClient.ConfigurationAPI.
			SensorGroupAssignmentsPost(context.Background()).
			SensorGroupAssignmentsPostRequest(*postRequest).
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.SensorGroupAssignmentResponse{
			Id:     "uid",
			Group:  *config_api_client.NewGroup("group_uid"),
			Sensor: *config_api_client.NewSensor("sensor_uid"),
			Type:   "networking-uxi/sensor-group-assignment",
		})
	})

	t.Run("Test ConfigurationAPI SensorGroupAssignmentsDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/sensor-group-assignments/uid").
			Reply(204)

		_, httpRes, err := apiClient.ConfigurationAPI.
			SensorGroupAssignmentsDelete(context.Background(), "uid").
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPI WiredNetworksGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/wired-networks").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(200).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":                   "uid",
						"name":                 "alias",
						"createdAt":            "2024-09-11T12:00:00.000Z",
						"updatedAt":            "2024-09-11T12:00:00.000Z",
						"ipVersion":            "ip_version",
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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.WiredNetworksResponse{
			Items: []config_api_client.WiredNetworksItem{
				{
					Id:                   "uid",
					Name:                 "alias",
					IpVersion:            "ip_version",
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
			Reply(200).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":                   "uid",
						"ssid":                 "ssid",
						"name":                 "alias",
						"createdAt":            "2024-09-11T12:00:00.000Z",
						"updatedAt":            "2024-09-11T12:00:00.000Z",
						"ipVersion":            "ip_version",
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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.WirelessNetworksResponse{
			Items: []config_api_client.WirelessNetworksItem{
				{
					Id:                   "uid",
					Name:                 "alias",
					Ssid:                 "ssid",
					Security:             *config_api_client.NewNullableString(&security),
					IpVersion:            "ip_version",
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
			Reply(200).
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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.NetworkGroupAssignmentsResponse{
			Items: []config_api_client.NetworkGroupAssignmentsItem{
				{
					Id:      "uid",
					Group:   *config_api_client.NewGroup("group_uid"),
					Network: *config_api_client.NewNetwork("network_uid"),
					Type:    "networking-uxi/network-group-assignment",
				},
			},
			Count: 1,
			Next:  *config_api_client.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPI NetworkGroupAssignmentsPost", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/networking-uxi/v1alpha1/network-group-assignments").
			JSON(map[string]interface{}{
				"groupId":   "group_uid",
				"networkId": "network_uid",
			}).
			Reply(200).
			JSON(map[string]interface{}{
				"id":      "uid",
				"group":   map[string]string{"id": "group_uid"},
				"network": map[string]string{"id": "network_uid"},
				"type":    "networking-uxi/network-group-assignment",
			})

		postRequest := config_api_client.NewNetworkGroupAssignmentsPostRequest(
			"group_uid",
			"network_uid",
		)
		resp, httpRes, err := apiClient.ConfigurationAPI.
			NetworkGroupAssignmentsPost(context.Background()).
			NetworkGroupAssignmentsPostRequest(*postRequest).
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.NetworkGroupAssignmentsPostResponse{
			Id:      "uid",
			Group:   *config_api_client.NewGroup("group_uid"),
			Network: *config_api_client.NewNetwork("network_uid"),
			Type:    "networking-uxi/network-group-assignment",
		})
	})

	t.Run("Test ConfigurationAPI NetworkGroupAssignmentsDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/network-group-assignments/uid").
			Reply(204)

		_, httpRes, err := apiClient.ConfigurationAPI.
			NetworkGroupAssignmentsDelete(context.Background(), "uid").
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPI ServiceTestGroupAssignmentsGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/networking-uxi/v1alpha1/service-test-group-assignments").
			MatchParams(map[string]string{"id": "uid", "limit": "10", "next": "some-cursor"}).
			Reply(200).
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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.ServiceTestGroupAssignmentsResponse{
			Items: []config_api_client.ServiceTestGroupAssignmentsItem{
				{
					Id:          "uid",
					Group:       *config_api_client.NewGroup("group_uid"),
					ServiceTest: *config_api_client.NewServiceTest("service_test_uid"),
					Type:        "networking-uxi/service-test-group-assignment",
				},
			},
			Count: 1,
			Next:  *config_api_client.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPI ServiceTestGroupAssignmentsPost", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/networking-uxi/v1alpha1/service-test-group-assignments").
			JSON(map[string]interface{}{
				"groupId":       "group_uid",
				"serviceTestId": "service_test_uid",
			}).
			Reply(200).
			JSON(map[string]interface{}{
				"id":          "uid",
				"group":       map[string]string{"id": "group_uid"},
				"serviceTest": map[string]string{"id": "service_test_uid"},
				"type":        "networking-uxi/service-test-group-assignment",
			})

		postRequest := config_api_client.NewServiceTestGroupAssignmentsPostRequest(
			"group_uid",
			"service_test_uid",
		)
		resp, httpRes, err := apiClient.ConfigurationAPI.
			ServiceTestGroupAssignmentsPost(context.Background()).
			ServiceTestGroupAssignmentsPostRequest(*postRequest).
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &config_api_client.ServiceTestGroupAssignmentsPostResponse{
			Id:          "uid",
			Group:       *config_api_client.NewGroup("group_uid"),
			ServiceTest: *config_api_client.NewServiceTest("service_test_uid"),
			Type:        "networking-uxi/service-test-group-assignment",
		})
	})

	t.Run("Test ConfigurationAPI ServiceTestGroupAssignmentsDelete", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/networking-uxi/v1alpha1/service-test-group-assignments/uid").
			Reply(204)

		_, httpRes, err := apiClient.ConfigurationAPI.
			ServiceTestGroupAssignmentsDelete(context.Background(), "uid").
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})
}
