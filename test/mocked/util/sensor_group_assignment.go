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

func GenerateSensorGroupAssignmentResponse(
	id string,
	postfix string,
) config_api_client.SensorGroupAssignmentsResponse {
	return config_api_client.SensorGroupAssignmentsResponse{
		Items: []config_api_client.SensorGroupAssignmentsItem{
			{
				Id:     id,
				Group:  *config_api_client.NewGroup("group_id" + postfix),
				Sensor: *config_api_client.NewSensor("sensor_id" + postfix),
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
) config_api_client.SensorGroupAssignmentsPostRequest {
	return config_api_client.SensorGroupAssignmentsPostRequest{
		GroupId:  "group_id" + postfix,
		SensorId: "sensor_id" + postfix,
	}
}

func GenerateSensorGroupAssignmentPostResponse(
	id string,
	postfix string,
) config_api_client.SensorGroupAssignmentResponse {
	return config_api_client.SensorGroupAssignmentResponse{
		Id:     id,
		Group:  *config_api_client.NewGroup("group_id" + postfix),
		Sensor: *config_api_client.NewSensor("sensor_id" + postfix),
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
	request config_api_client.SensorGroupAssignmentsPostRequest,
	response config_api_client.SensorGroupAssignmentResponse,
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
