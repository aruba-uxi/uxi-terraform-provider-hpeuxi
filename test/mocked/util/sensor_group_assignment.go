/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateSensorGroupAssignmentResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":     id,
		"group":  map[string]string{"id": "group_id" + postfix},
		"sensor": map[string]string{"id": "sensor_id" + postfix},
		"type":   shared.SensorGroupAssignmentType,
	}
}

func GenerateSensorGroupAssignmentRequest(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"groupId":  "group_id" + postfix,
		"sensorId": "sensor_id" + postfix,
	}
}

func MockGetSensorGroupAssignment(id string, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Get(shared.SensorGroupAssignmentPath).
		MatchHeader("Authorization", mockToken).
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
	gock.New(MockUxiUrl).
		Post(shared.SensorGroupAssignmentPath).
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", mockToken).
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteSensorGroupAssignment(id string, times int) {
	gock.New(MockUxiUrl).
		Delete(shared.SensorGroupAssignmentPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}
