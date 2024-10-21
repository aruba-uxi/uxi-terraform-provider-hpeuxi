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

func Test_config_api_client_ConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	configuration.Host = "localhost:80"
	configuration.Scheme = "http"
	apiClient := openapiclient.NewAPIClient(configuration)

	defer gock.Off()

	t.Run("Test ConfigurationAPIService GroupsGetUxiV1alpha1GroupsGet", func(t *testing.T) {
		parent_uid := "parent_uid"
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/uxi/v1alpha1/groups").
			MatchParams(map[string]string{"uid": "uid", "limit": "10", "cursor": "some-cursor"}).
			Reply(200).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":     "uid",
						"name":   "name",
						"parent": map[string]string{"id": parent_uid},
						"path":   "root_uid.parent_uid.uid",
						"type":   "uxi/group",
					},
				},
				"next":  nil,
				"count": 1,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GroupsGetUxiV1alpha1GroupsGet(context.Background()).
			Uid("uid").
			Limit(10).
			Cursor("some-cursor").
			Execute()

		resourceType := "uxi/group"

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.GroupsGetResponse{
			Items: []openapiclient.GroupsGetItem{
				{
					Id:     "uid",
					Name:   "name",
					Parent: *openapiclient.NewNullableParent(openapiclient.NewParent("parent_uid")),
					Path:   "root_uid.parent_uid.uid",
					Type:   &resourceType,
				},
			},
			Next:  *openapiclient.NewNullableString(nil),
			Count: 1,
		})
	})

	t.Run("Test ConfigurationAPIService GroupsPostUxiV1alpha1GroupsPost", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/uxi/v1alpha1/groups").
			MatchType("json").
			BodyString(`{"parentId": "parent.uid", "name": "name"}`).
			Reply(200).
			JSON(map[string]interface{}{
				"id":     "node",
				"name":   "name",
				"parent": map[string]string{"id": "parent.uid"},
				"path":   "parent.uid.node",
				"type":   "uxi/group",
			})
		groupsPostRequest := openapiclient.NewGroupsPostRequest("name")
		groupsPostRequest.SetParentId("parent.uid")
		resp, httpRes, err := apiClient.ConfigurationAPI.GroupsPostUxiV1alpha1GroupsPost(context.Background()).GroupsPostRequest(*groupsPostRequest).Execute()

		resourceType := "uxi/group"

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.GroupsPostResponse{
			Id:     "node",
			Name:   "name",
			Parent: *openapiclient.NewParent("parent.uid"),
			Path:   "parent.uid.node",
			Type:   &resourceType,
		})
	})

	t.Run("Test ConfigurationAPIService GroupsPatchUxiV1alpha1GroupsGroupUidPatch", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Patch("/uxi/v1alpha1/groups/node").
			MatchType("json").
			BodyString(`{"name": "new_name"}`).
			Reply(200).
			JSON(map[string]interface{}{
				"id":     "node",
				"name":   "new_name",
				"parent": map[string]string{"id": "parent.uid"},
				"path":   "parent.uid.node",
				"type":   "uxi/group",
			})
		groupsPatchRequest := openapiclient.NewGroupsPatchRequest("new_name")
		resp, httpRes, err := apiClient.ConfigurationAPI.GroupsPatchUxiV1alpha1GroupsGroupUidPatch(context.Background(), "node").GroupsPatchRequest(*groupsPatchRequest).Execute()

		resourceType := "uxi/group"

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.GroupsPatchResponse{
			Id:     "node",
			Name:   "new_name",
			Parent: *openapiclient.NewParent("parent.uid"),
			Path:   "parent.uid.node",
			Type:   &resourceType,
		})
	})

	t.Run("Test ConfigurationAPIService GetUxiV1alpha1SensorGroupAssignmentsGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/uxi/v1alpha1/sensor-group-assignments").
			MatchParams(map[string]string{"uid": "uid", "limit": "10", "cursor": "some-cursor"}).
			Reply(200).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":     "uid",
						"group":  map[string]string{"id": "group_uid"},
						"sensor": map[string]string{"id": "sensor_uid"},
						"type":   "uxi/sensor-group-assignment",
					},
				},
				"count": 1,
				"next":  nil,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GetUxiV1alpha1SensorGroupAssignmentsGet(context.Background()).
			Uid("uid").
			Limit(10).
			Cursor("some-cursor").
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		resourceType := "uxi/sensor-group-assignment"
		assert.Equal(t, resp, &openapiclient.SensorGroupAssignmentsResponse{
			Items: []openapiclient.SensorGroupAssignmentsItem{
				{
					Id:     "uid",
					Group:  *openapiclient.NewGroup("group_uid"),
					Sensor: *openapiclient.NewSensor("sensor_uid"),
					Type:   &resourceType,
				},
			},
			Count: 1,
			Next:  *openapiclient.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPIService PostUxiV1alpha1SensorGroupAssignmentsPost", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/uxi/v1alpha1/sensor-group-assignments").
			MatchType("json").
			BodyString(`{"groupId": "group_uid", "sensorId": "sensor_uid"}`).
			Reply(200).
			JSON(map[string]interface{}{
				"id":     "uid",
				"group":  map[string]string{"id": "group_uid"},
				"sensor": map[string]string{"id": "sensor_uid"},
				"type":   "uxi/sensor-group-assignment",
			})

		postRequest := openapiclient.NewSensorGroupAssignmentsPostRequest("group_uid", "sensor_uid")
		resp, httpRes, err := apiClient.ConfigurationAPI.
			PostUxiV1alpha1SensorGroupAssignmentsPost(context.Background()).
			SensorGroupAssignmentsPostRequest(*postRequest).
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		resourceType := "uxi/sensor-group-assignment"
		assert.Equal(t, resp, &openapiclient.SensorGroupAssignmentResponse{
			Id:     "uid",
			Group:  *openapiclient.NewGroup("group_uid"),
			Sensor: *openapiclient.NewSensor("sensor_uid"),
			Type:   &resourceType,
		})
	})

	t.Run("Test ConfigurationAPIService DeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/uxi/v1alpha1/sensor-group-assignments/uid").
			Reply(204)

		_, httpRes, err := apiClient.ConfigurationAPI.
			DeleteSensorGroupAssignmentUxiV1alpha1SensorGroupAssignmentsIdDelete(context.Background(), "uid").
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPIService GetConfigurationAppV1WiredNetworksGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/uxi/v1alpha1/wired-networks").
			MatchParams(map[string]string{"uid": "uid", "limit": "10", "cursor": "some-cursor"}).
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
						"type":                 "uxi/wired-network",
					},
				},
				"count": 1,
				"next":  nil,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GetUxiV1alpha1WiredNetworksGet(context.Background()).
			Uid("uid").
			Limit(10).
			Cursor("some-cursor").
			Execute()

		security := "security"
		dnsLookupDomain := "dns_lookup_domain"
		var vlanId int32 = 1
		resourceType := "uxi/wired-network"

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
					Type:                 &resourceType,
				},
			},
			Count: 1,
			Next:  *openapiclient.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPIService GetUxiV1alpha1WirelessNetworksGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/uxi/v1alpha1/wireless-networks").
			MatchParams(map[string]string{"uid": "uid", "limit": "10", "cursor": "some-cursor"}).
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
						"type":                 "uxi/wireless-network",
					},
				},
				"count": 1,
				"next":  nil,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GetUxiV1alpha1WirelessNetworksGet(context.Background()).
			Uid("uid").
			Limit(10).
			Cursor("some-cursor").
			Execute()

		security := "security"
		dnsLookupDomain := "dns_lookup_domain"
		resourceType := "uxi/wireless-network"

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
					Type:                 &resourceType,
				},
			},
			Count: 1,
			Next:  *openapiclient.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPIService GetUxiV1alpha1NetworkGroupAssignmentsGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/uxi/v1alpha1/network-group-assignments").
			MatchParams(map[string]string{"uid": "uid", "limit": "10", "cursor": "some-cursor"}).
			Reply(200).
			JSON(map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":      "uid",
						"group":   map[string]string{"id": "group_uid"},
						"network": map[string]string{"id": "network_uid"},
						"type":    "uxi/network-group-assignment",
					},
				},
				"count": 1,
				"next":  nil,
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GetUxiV1alpha1NetworkGroupAssignmentsGet(context.Background()).
			Uid("uid").
			Limit(10).
			Cursor("some-cursor").
			Execute()

		resourceType := "uxi/network-group-assignment"

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.NetworkGroupAssignmentsResponse{
			Items: []openapiclient.NetworkGroupAssignmentsItem{
				{
					Id:      "uid",
					Group:   *openapiclient.NewGroup("group_uid"),
					Network: *openapiclient.NewNetwork("network_uid"),
					Type:    &resourceType,
				},
			},
			Count: 1,
			Next:  *openapiclient.NewNullableString(nil),
		})
	})

	t.Run("Test ConfigurationAPIService PostUxiV1alpha1NetworkGroupAssignmentsPost", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/uxi/v1alpha1/network-group-assignments").
			MatchType("json").
			BodyString(`{"groupId": "group_uid", "networkId": "network_uid"}`).
			Reply(200).
			JSON(map[string]interface{}{
				"id":      "uid",
				"group":   map[string]string{"id": "group_uid"},
				"network": map[string]string{"id": "network_uid"},
				"type":    "uxi/network-group-assignment",
			})

		postRequest := openapiclient.NewNetworkGroupAssignmentsPostRequest("group_uid", "network_uid")
		resp, httpRes, err := apiClient.ConfigurationAPI.
			PostUxiV1alpha1NetworkGroupAssignmentsPost(context.Background()).
			NetworkGroupAssignmentsPostRequest(*postRequest).
			Execute()

		resourceType := "uxi/network-group-assignment"

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.NetworkGroupAssignmentsPostResponse{
			Id:      "uid",
			Group:   *openapiclient.NewGroup("group_uid"),
			Network: *openapiclient.NewNetwork("network_uid"),
			Type:    &resourceType,
		})
	})

	t.Run("Test ConfigurationAPIService DeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDelete", func(t *testing.T) {
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Delete("/uxi/v1alpha1/network-group-assignments/uid").
			Reply(204)

		_, httpRes, err := apiClient.ConfigurationAPI.
			DeleteNetworkGroupAssignmentUxiV1alpha1NetworkGroupAssignmentsIdDelete(context.Background(), "uid").
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 204, httpRes.StatusCode)
	})

	t.Run("Test ConfigurationAPIService PostUxiV1alpha1ServiceTestGroupAssignmentsPost", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Post("/uxi/v1alpha1/service-test-group-assignments").
			MatchType("json").
			BodyString(`{"groupId": "group_uid", "serviceTestId": "service_test_uid"}`).
			Reply(200).
			JSON(map[string]interface{}{
				"id":           "uid",
				"group":        map[string]string{"id": "group_uid"},
				"service_test": map[string]string{"id": "service_test_uid"},
				"type":         "uxi/service-test-group-assignment",
			})

		postRequest := openapiclient.NewServiceTestGroupAssignmentsPostRequest("group_uid", "service_test_uid")
		resp, httpRes, err := apiClient.ConfigurationAPI.
			PostUxiV1alpha1ServiceTestGroupAssignmentsPost(context.Background()).
			ServiceTestGroupAssignmentsPostRequest(*postRequest).
			Execute()

		resourceType := "uxi/service-test-group-assignment"

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.ServiceTestGroupAssignmentsPostResponse{
			Id:          "uid",
			Group:       *openapiclient.NewGroup("group_uid"),
			ServiceTest: *openapiclient.NewServiceTest("service_test_uid"),
			Type:        &resourceType,
		})
	})

}
