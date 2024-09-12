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

	t.Run("Test ConfigurationAPIService GroupsPostConfigurationAppV1GroupsPost", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).Post("/configuration/app/v1/groups").Reply(200).JSON(map[string]string{"parent_uid": "parent.uid", "name": "name", "uid": "node", "path": "parent.uid.node"})
		groups_post_request := openapiclient.NewGroupsPostRequest("parent_uid", "name")
		resp, httpRes, err := apiClient.ConfigurationAPI.GroupsPostConfigurationAppV1GroupsPost(context.Background()).GroupsPostRequest(*groups_post_request).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.GroupsPostResponse{ParentUid: "parent.uid", Name: "name", Uid: "node", Path: "parent.uid.node"})
	})

	t.Run("Test ConfigurationAPIService GetConfigurationAppV1SensorGroupAssignmentsGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/configuration/app/v1/sensor-group-assignments").
			MatchParams(map[string]string{"uid": "uid", "limit": "10", "cursor": "some-cursor"}).
			Reply(200).
			JSON(map[string]interface{}{
				"sensor_group_assignments": []map[string]string{
					{
						"uid":        "uid",
						"group_uid":  "group_uid",
						"sensor_uid": "sensor_uid",
					},
				},
				"pagination": map[string]interface{}{
					"limit":    10,
					"first":    nil,
					"next":     nil,
					"previous": nil,
					"last":     nil,
				},
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GetConfigurationAppV1SensorGroupAssignmentsGet(context.Background()).
			Uid("uid").
			Limit(10).
			Cursor("some-cursor").
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.SensorGroupAssignmentsResponse{
			SensorGroupAssignments: []openapiclient.SensorGroupAssignmentsItem{
				{
					Uid:       "uid",
					GroupUid:  "group_uid",
					SensorUid: "sensor_uid",
				},
			},
			Pagination: openapiclient.PaginationDetails{
				Limit:    10,
				Next:     *openapiclient.NewNullableString(nil),
				Previous: *openapiclient.NewNullableString(nil),
				First:    *openapiclient.NewNullableString(nil),
				Last:     *openapiclient.NewNullableString(nil),
			},
		})
	})

	t.Run("Test ConfigurationAPIService GetConfigurationAppV1SensorGroupAssignmentsGet", func(t *testing.T) {
		parent_uid := "parent_uid"
		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/configuration/app/v1/groups").
			MatchParams(map[string]string{"uid": "uid", "limit": "10", "cursor": "some-cursor"}).
			Reply(200).
			JSON(map[string]interface{}{
				"groups": []map[string]string{
					{
						"uid":        "uid",
						"name":       "name",
						"parent_uid": parent_uid,
						"path":       "root_uid.parent_uid.uid",
					},
				},
				"pagination": map[string]interface{}{
					"limit":    10,
					"first":    nil,
					"next":     nil,
					"previous": nil,
					"last":     nil,
				},
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GroupsGetConfigurationAppV1GroupsGet(context.Background()).
			Uid("uid").
			Limit(10).
			Cursor("some-cursor").
			Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.GroupsGetResponse{
			Groups: []openapiclient.GroupsGetItem{
				{
					Uid:       "uid",
					Name:      "name",
					ParentUid: *openapiclient.NewNullableString(&parent_uid),
					Path:      "root_uid.parent_uid.uid",
				},
			},
			Pagination: openapiclient.PaginationDetails{
				Limit:    10,
				Next:     *openapiclient.NewNullableString(nil),
				Previous: *openapiclient.NewNullableString(nil),
				First:    *openapiclient.NewNullableString(nil),
				Last:     *openapiclient.NewNullableString(nil),
			},
		})
	})

	t.Run("Test ConfigurationAPIService GetConfigurationAppV1WiredNetworksGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/configuration/app/v1/wired-networks").
			MatchParams(map[string]string{"uid": "uid", "limit": "10", "cursor": "some-cursor"}).
			Reply(200).
			JSON(map[string]interface{}{
				"wired_networks": []map[string]interface{}{
					{
						"uid":                   "uid",
						"alias":                 "alias",
						"datetime_created":      "2024-09-11T12:00:00.000Z",
						"datetime_updated":      "2024-09-11T12:00:00.000Z",
						"ip_version":            "ip_version",
						"security":              "security",
						"dns_lookup_domain":     "dns_lookup_domain",
						"disable_edns":          true,
						"use_dns64":             false,
						"external_connectivity": true,
						"vlan_id":               1,
					},
				},
				"pagination": map[string]interface{}{
					"limit":    10,
					"first":    nil,
					"next":     nil,
					"previous": nil,
					"last":     nil,
				},
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GetConfigurationAppV1WiredNetworksGet(context.Background()).
			Uid("uid").
			Limit(10).
			Cursor("some-cursor").
			Execute()

		security := "security"
		dnsLookupDomain := "dns_lookup_domain"
		var vlanId int32 = 1

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.WiredNetworksResponse{
			WiredNetworks: []openapiclient.WiredNetworksItem{
				{
					Uid:                  "uid",
					Alias:                "alias",
					IpVersion:            "ip_version",
					DatetimeUpdated:      time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					DatetimeCreated:      time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					Security:             *openapiclient.NewNullableString(&security),
					DnsLookupDomain:      *openapiclient.NewNullableString(&dnsLookupDomain),
					DisableEdns:          true,
					UseDns64:             false,
					ExternalConnectivity: true,
					VlanId:               *openapiclient.NewNullableInt32(&vlanId),
				},
			},
			Pagination: openapiclient.PaginationDetails{
				Limit:    10,
				Next:     *openapiclient.NewNullableString(nil),
				Previous: *openapiclient.NewNullableString(nil),
				First:    *openapiclient.NewNullableString(nil),
				Last:     *openapiclient.NewNullableString(nil),
			},
		})
	})

	t.Run("Test ConfigurationAPIService GetConfigurationAppV1WiredNetworksGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/configuration/app/v1/wireless-networks").
			MatchParams(map[string]string{"uid": "uid", "limit": "10", "cursor": "some-cursor"}).
			Reply(200).
			JSON(map[string]interface{}{
				"wireless_networks": []map[string]interface{}{
					{
						"uid":                   "uid",
						"ssid":                  "ssid",
						"alias":                 "alias",
						"datetime_created":      "2024-09-11T12:00:00.000Z",
						"datetime_updated":      "2024-09-11T12:00:00.000Z",
						"ip_version":            "ip_version",
						"security":              "security",
						"hidden":                false,
						"band_locking":          "band_locking",
						"dns_lookup_domain":     "dns_lookup_domain",
						"disable_edns":          true,
						"use_dns64":             false,
						"external_connectivity": true,
					},
				},
				"pagination": map[string]interface{}{
					"limit":    10,
					"first":    nil,
					"next":     nil,
					"previous": nil,
					"last":     nil,
				},
			})
		resp, httpRes, err := apiClient.ConfigurationAPI.
			GetConfigurationAppV1WirelessNetworksGet(context.Background()).
			Uid("uid").
			Limit(10).
			Cursor("some-cursor").
			Execute()

		security := "security"
		dnsLookupDomain := "dns_lookup_domain"

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.WirelessNetworksResponse{
			WirelessNetworks: []openapiclient.WirelessNetworksItem{
				{
					Uid:                  "uid",
					Alias:                "alias",
					Ssid:                 "ssid",
					Security:             *openapiclient.NewNullableString(&security),
					IpVersion:            "ip_version",
					DatetimeCreated:      time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					DatetimeUpdated:      time.Date(2024, 9, 11, 12, 0, 0, 0, time.UTC),
					Hidden:               false,
					BandLocking:          "band_locking",
					DnsLookupDomain:      *openapiclient.NewNullableString(&dnsLookupDomain),
					DisableEdns:          true,
					UseDns64:             false,
					ExternalConnectivity: true,
				},
			},
			Pagination: openapiclient.PaginationDetails{
				Limit:    10,
				Next:     *openapiclient.NewNullableString(nil),
				Previous: *openapiclient.NewNullableString(nil),
				First:    *openapiclient.NewNullableString(nil),
				Last:     *openapiclient.NewNullableString(nil),
			},
		})
	})

}
