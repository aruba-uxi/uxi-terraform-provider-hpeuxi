package config_api_client

import (
	"context"
	openapiclient "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_config_api_client_ConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	configuration.Host = "localhost:80"
	configuration.Scheme = "http"
	apiClient := openapiclient.NewAPIClient(configuration)

	defer gock.Off()

	t.Run("Test ConfigurationAPIService GetLivezHealthLivezGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).Get("/health/livez").Reply(200).JSON(map[string]string{"status": "OK"})

		resp, httpRes, err := apiClient.ConfigurationAPI.GetLivezHealthLivezGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, httpRes.StatusCode, 200)
		assert.Equal(t, resp, &openapiclient.LivenessResponse{Status: "OK"})
	})

	t.Run("Test ConfigurationAPIService GetReadyzHealthReadyzGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).Get("/health/readyz").Reply(200).JSON(map[string]interface{}{"status": "OK", "data": map[string]string{}})

		resp, httpRes, err := apiClient.ConfigurationAPI.GetReadyzHealthReadyzGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.ReadinessResponse{Status: "OK", Data: map[string]string{}})
	})

	t.Run("Test ConfigurationAPIService GetStatusHealthStatusGet", func(t *testing.T) {

		gock.New(configuration.Scheme + "://" + configuration.Host).Get("/health/status").Reply(200).JSON(map[string]string{"name": "Configuration-API", "version": "1.0.0"})

		resp, httpRes, err := apiClient.ConfigurationAPI.GetStatusHealthStatusGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, resp, &openapiclient.StatusResponse{Name: "Configuration-API", Version: "1.0.0"})
	})

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
			SensorGroupAssignments: []openapiclient.SensorGroupAssignment{
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

		gock.New(configuration.Scheme + "://" + configuration.Host).
			Get("/configuration/app/v1/groups").
			MatchParams(map[string]string{"uid": "uid", "limit": "10", "cursor": "some-cursor"}).
			Reply(200).
			JSON(map[string]interface{}{
				"groups": []map[string]string{
					{
						"uid":        "uid",
						"name":       "name",
						"parent_uid": "parent_uid",
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
					ParentUid: "parent_uid",
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
}
