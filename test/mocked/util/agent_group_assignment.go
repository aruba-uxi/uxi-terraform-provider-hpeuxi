/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateAgentGroupAssignmentRequest(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"groupId": "group_id" + postfix,
		"agentId": "agent_id" + postfix,
	}
}

func GenerateAgentGroupAssignmentResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":    id,
		"group": map[string]string{"id": "group_id" + postfix},
		"agent": map[string]string{"id": "agent_id" + postfix},
		"type":  shared.AgentGroupAssignmentType,
	}
}

func MockGetAgentGroupAssignment(id string, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Get(shared.AgentGroupAssignmentPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostAgentGroupAssignment(
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	gock.New(MockUxiUrl).
		Post(shared.AgentGroupAssignmentPath).
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", mockToken).
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteAgentGroupAssignment(id string, times int) {
	gock.New(MockUxiUrl).
		Delete(shared.AgentGroupAssignmentPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}
