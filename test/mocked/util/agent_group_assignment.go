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

func GenerateAgentGroupAssignmentPostRequest(
	id string,
	postfix string,
) config_api_client.AgentGroupAssignmentsPostRequest {
	return config_api_client.AgentGroupAssignmentsPostRequest{
		GroupId: "group_id" + postfix,
		AgentId: "agent_id" + postfix,
	}
}

func GenerateAgentGroupAssignmentPostResponse(
	id string,
	postfix string,
) config_api_client.AgentGroupAssignmentResponse {
	return config_api_client.AgentGroupAssignmentResponse{
		Id:    id,
		Group: *config_api_client.NewGroup("group_id" + postfix),
		Agent: *config_api_client.NewAgent("agent_id" + postfix),
		Type:  shared.AgentGroupAssignmentType,
	}
}

func GenerateAgentGroupAssignmentsResponse(
	id string,
	postfix string,
) config_api_client.AgentGroupAssignmentsResponse {
	return config_api_client.AgentGroupAssignmentsResponse{
		Items: []config_api_client.AgentGroupAssignmentsItem{
			{
				Id:    id,
				Group: *config_api_client.NewGroup("group_id" + postfix),
				Agent: *config_api_client.NewAgent("agent_id" + postfix),
				Type:  shared.AgentGroupAssignmentType,
			},
		},
		Count: 1,
		Next:  *config_api_client.NewNullableString(nil),
	}
}

func MockGetAgentGroupAssignment(id string, response interface{}, times int) {
	gock.New(MockUXIURL).
		Get(shared.AgentGroupAssignmentPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostAgentGroupAssignment(
	request config_api_client.AgentGroupAssignmentsPostRequest,
	response config_api_client.AgentGroupAssignmentResponse,
	times int,
) {
	gock.New(MockUXIURL).
		Post(shared.AgentGroupAssignmentPath).
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", mockToken).
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteAgentGroupAssignment(id string, times int) {
	gock.New(MockUXIURL).
		Delete(shared.AgentGroupAssignmentPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}
