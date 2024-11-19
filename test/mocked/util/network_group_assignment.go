package util

import (
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateNetworkGroupAssignmentResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":      id,
		"group":   map[string]string{"id": "group_id" + postfix},
		"network": map[string]string{"id": "network_id" + postfix},
		"type":    shared.NetworkGroupAssignmentType,
	}
}

func GenerateNetworkGroupAssignmentRequest(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"groupId":   "group_id" + postfix,
		"networkId": "network_id" + postfix,
	}
}

func MockGetNetworkGroupAssignment(id string, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Get(shared.NetworkGroupAssignmentPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostNetworkGroupAssignment(
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	gock.New(MockUxiUrl).
		Post(shared.NetworkGroupAssignmentPath).
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", mockToken).
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteNetworkGroupAssignment(id string, times int) {
	gock.New(MockUxiUrl).
		Delete(shared.NetworkGroupAssignmentPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}
