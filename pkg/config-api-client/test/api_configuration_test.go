package config_api_client

import (
	"context"
	"testing"
	"time"

	openapiclient "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigurationAPI(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	configuration.Host = "localhost:80"
	configuration.Scheme = "http"
	apiClient := openapiclient.NewAPIClient(configuration)

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
		assert.Equal(t, resp, &openapiclient.AgentsResponse{
			Items: []openapiclient.AgentItem{
				{
					Id:                 "uid",
					Serial:             "serial",
					Name:               "name",
					ModelNumber:        *openapiclient.NewNullableString(&modelNumber),
					WifiMacAddress:     *openapiclient.NewNullableString(&wifiMacAddress),
					EthernetMacAddress: *openapiclient.NewNullableString(&ethernetMacAddress),
					Notes:              *openapiclient.NewNullableString(&notes),
					PcapMode:           *openapiclient.NewNullableString(&pcapMode),
					Type:               "networking-uxi/sensor",
				},
			},
			Next:  *openapiclient.NewNullableString(nil),
			Count: 1,
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
		assert.Equal(t, resp, &openapiclient.GroupsGetResponse{
			Items: []openapiclient.GroupsGetItem{
				{
					Id:     "uid",
					Name:   "name",
					Parent: *openapiclient.NewNullableParent(openapiclient.NewParent("parent_uid")),
					Path:   "root_uid.parent_uid.uid",
					Type:   "networking-uxi/group",
				},
			},
			Next:  *openapiclient.NewNullableString(nil),
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
		groupsPostRequest := openapiclient.NewGroupsPostRequest("name")
		groupsPostRequest.SetParentId("parent.uid")
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GroupsPost(context.Background()).
			GroupsPostRequest(*groupsPostRequest).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.GroupsPostResponse{
			Id:     "node",
			Name:   "name",
			Parent: *openapiclient.NewParent("parent.uid"),
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
		groupsPatchRequest := openapiclient.NewGroupsPatchRequest("new_name")
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GroupsPatch(context.Background(), "node").
			GroupsPatchRequest(*groupsPatchRequest).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.GroupsPatchResponse{
			Id:     "node",
			Name:   "new_name",
			Parent: *openapiclient.NewParent("parent.uid"),
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
		assert.Equal(t, resp, &openapiclient.SensorsResponse{
			Items: []openapiclient.SensorItem{
				{
					Id:                 "uid",
					Serial:             "serial",
					Name:               "name",
					ModelNumber:        "model_number",
					WifiMacAddress:     *openapiclient.NewNullableString(&WifiMacAddress),
					EthernetMacAddress: *openapiclient.NewNullableString(&EthernetMacAddress),
					AddressNote:        *openapiclient.NewNullableString(&AddressNote),
					Longitude:          *openapiclient.NewNullableFloat32(&Longitude),
					Latitude:           *openapiclient.NewNullableFloat32(&Latitude),
					Notes:              *openapiclient.NewNullableString(&Notes),
					PcapMode:           *openapiclient.NewNullableString(&PcapMode),
					Type:               "networking-uxi/sensor",
				},
			},
			Next:  *openapiclient.NewNullableString(nil),
			Count: 1,
		})
	})

	t.Run("Test ConfigurationAPI GroupsPatch", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Patch("/networking-uxi/v1alpha1/sensors/uid").
			// TODO: uncomment this once spec has been updated to merge-patch+json
			// MatchHeader("Content-Type", "application/merge-patch+json").
			// TODO: Change these fields to camelCase once the spec has been updated
			BodyString(`{"name":"new_name","address_note":"new_address_note","notes":"new_notes","pcap_mode":"off"}`).
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
		sensorsPatchRequest := openapiclient.SensorsPatchRequest{
			Name:        &name,
			AddressNote: *openapiclient.NewNullableString(&addressNote),
			Notes:       *openapiclient.NewNullableString(&notes),
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
		assert.Equal(t, resp, &openapiclient.SensorsPatchResponse{
			Id:                 "uid",
			Serial:             "serial",
			Name:               "new_name",
			ModelNumber:        "model_number",
			WifiMacAddress:     *openapiclient.NewNullableString(&wifiMacAddress),
			EthernetMacAddress: *openapiclient.NewNullableString(&ethernetMacAddress),
			AddressNote:        *openapiclient.NewNullableString(&addressNote),
			Longitude:          *openapiclient.NewNullableFloat32(&longitude),
			Latitude:           *openapiclient.NewNullableFloat32(&latitude),
			Notes:              *openapiclient.NewNullableString(&notes),
			PcapMode:           *openapiclient.NewNullableString(&pcapMode),
			Type:               "networking-uxi/sensor",
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
		assert.Equal(t, resp, &openapiclient.SensorGroupAssignmentsResponse{
			Items: []openapiclient.SensorGroupAssignmentsItem{
				{
					Id:     "uid",
					Group:  *openapiclient.NewGroup("group_uid"),
					Sensor: *openapiclient.NewSensor("sensor_uid"),
					Type:   "networking-uxi/sensor-group-assignment",
				},
			},
			Count: 1,
			Next:  *openapiclient.NewNullableString(nil),
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

		postRequest := openapiclient.NewSensorGroupAssignmentsPostRequest("group_uid", "sensor_uid")
		resp, httpRes, err := apiClient.ConfigurationAPI.
			SensorGroupAssignmentsPost(context.Background()).
			SensorGroupAssignmentsPostRequest(*postRequest).
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.SensorGroupAssignmentResponse{
			Id:     "uid",
			Group:  *openapiclient.NewGroup("group_uid"),
			Sensor: *openapiclient.NewSensor("sensor_uid"),
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
		assert.Equal(t, resp, &openapiclient.WiredNetworksResponse{
			Items: []openapiclient.WiredNetworksItem{
				{
					Id:                   "uid",
					Name:                 "alias",
					IpVersion:            "ip_version",
					UpdatedAt:            time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					CreatedAt:            time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					Security:             *openapiclient.NewNullableString(&security),
					DnsLookupDomain:      *openapiclient.NewNullableString(&dnsLookupDomain),
					DisableEdns:          true,
					UseDns64:             false,
					ExternalConnectivity: true,
					VLanId:               *openapiclient.NewNullableInt32(&vlanId),
					Type:                 "networking-uxi/wired-network",
				},
			},
			Count: 1,
			Next:  *openapiclient.NewNullableString(nil),
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
		assert.Equal(t, resp, &openapiclient.WirelessNetworksResponse{
			Items: []openapiclient.WirelessNetworksItem{
				{
					Id:                   "uid",
					Name:                 "alias",
					Ssid:                 "ssid",
					Security:             *openapiclient.NewNullableString(&security),
					IpVersion:            "ip_version",
					CreatedAt:            time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					UpdatedAt:            time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					Hidden:               false,
					BandLocking:          "band_locking",
					DnsLookupDomain:      *openapiclient.NewNullableString(&dnsLookupDomain),
					DisableEdns:          true,
					UseDns64:             false,
					ExternalConnectivity: true,
					Type:                 "networking-uxi/wireless-network",
				},
			},
			Count: 1,
			Next:  *openapiclient.NewNullableString(nil),
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
		assert.Equal(t, resp, &openapiclient.NetworkGroupAssignmentsResponse{
			Items: []openapiclient.NetworkGroupAssignmentsItem{
				{
					Id:      "uid",
					Group:   *openapiclient.NewGroup("group_uid"),
					Network: *openapiclient.NewNetwork("network_uid"),
					Type:    "networking-uxi/network-group-assignment",
				},
			},
			Count: 1,
			Next:  *openapiclient.NewNullableString(nil),
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

		postRequest := openapiclient.NewNetworkGroupAssignmentsPostRequest(
			"group_uid",
			"network_uid",
		)
		resp, httpRes, err := apiClient.ConfigurationAPI.
			NetworkGroupAssignmentsPost(context.Background()).
			NetworkGroupAssignmentsPostRequest(*postRequest).
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.NetworkGroupAssignmentsPostResponse{
			Id:      "uid",
			Group:   *openapiclient.NewGroup("group_uid"),
			Network: *openapiclient.NewNetwork("network_uid"),
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

		postRequest := openapiclient.NewServiceTestGroupAssignmentsPostRequest(
			"group_uid",
			"service_test_uid",
		)
		resp, httpRes, err := apiClient.ConfigurationAPI.
			ServiceTestGroupAssignmentsPost(context.Background()).
			ServiceTestGroupAssignmentsPostRequest(*postRequest).
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.ServiceTestGroupAssignmentsPostResponse{
			Id:          "uid",
			Group:       *openapiclient.NewGroup("group_uid"),
			ServiceTest: *openapiclient.NewServiceTest("service_test_uid"),
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
