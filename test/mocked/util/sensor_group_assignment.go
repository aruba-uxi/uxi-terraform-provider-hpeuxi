/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/h2non/gock"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func GenerateSensorGroupAssignmentsGetResponse(
	id string,
	postfix string,
) config_api_client.SensorGroupAssignmentsGetResponse {
	return config_api_client.SensorGroupAssignmentsGetResponse{
		Items: []config_api_client.SensorGroupAssignmentsGetItem{
			{
				Id:     id,
				Group:  *config_api_client.NewSensorGroupAssignmentsGetGroup("group_id" + postfix),
				Sensor: *config_api_client.NewSensorGroupAssignmentsGetSensor("sensor_id" + postfix),
				Type:   shared.SensorGroupAssignmentType,
			},
		},
		Count: 1,
		Next:  *config_api_client.NewNullableString(nil),
	}
}

func GenerateSensorGroupAssignmentPostRequest(
	id string,
	postfix string,
) config_api_client.SensorGroupAssignmentPostRequest {
	return config_api_client.SensorGroupAssignmentPostRequest{
		GroupId:  "group_id" + postfix,
		SensorId: "sensor_id" + postfix,
	}
}

func GenerateSensorGroupAssignmentPostResponse(
	id string,
	postfix string,
) config_api_client.SensorGroupAssignmentPostResponse {
	return config_api_client.SensorGroupAssignmentPostResponse{
		Id:     id,
		Group:  *config_api_client.NewSensorGroupAssignmentPostGroup("group_id" + postfix),
		Sensor: *config_api_client.NewSensorGroupAssignmentPostSensor("sensor_id" + postfix),
		Type:   shared.SensorGroupAssignmentType,
	}
}

func MockGetSensorGroupAssignment(id string, response interface{}, times int) {
	gock.New(MockUXIURL).
		Get(shared.SensorGroupAssignmentPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostSensorGroupAssignment(
	request config_api_client.SensorGroupAssignmentPostRequest,
	response config_api_client.SensorGroupAssignmentPostResponse,
	times int,
) {
	gock.New(MockUXIURL).
		Post(shared.SensorGroupAssignmentPath).
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", mockToken).
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteSensorGroupAssignment(id string, times int) {
	gock.New(MockUXIURL).
		Delete(shared.SensorGroupAssignmentPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}
