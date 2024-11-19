package util

import (
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateServiceTestGroupAssignmentResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":          id,
		"group":       map[string]string{"id": "group_id" + postfix},
		"serviceTest": map[string]string{"id": "service_test_id" + postfix},
		"type":        shared.ServiceTestGroupAssignmentType,
	}
}

func GenerateServiceTestGroupAssignmentRequest(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"groupId":       "group_id" + postfix,
		"serviceTestId": "service_test_id" + postfix,
	}
}

func MockGetServiceTestGroupAssignment(id string, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Get(shared.ServiceTestGroupAssignmentPath).
		MatchHeader("Authorization", mockToken).
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
	gock.New(MockUxiUrl).
		Post(shared.ServiceTestGroupAssignmentPath).
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", mockToken).
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteServiceTestGroupAssignment(id string, times int) {
	gock.New(MockUxiUrl).
		Delete(shared.ServiceTestGroupAssignmentPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}
