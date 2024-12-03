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
) config_api_client.AgentGroupAssignmentPostRequest {
	return config_api_client.AgentGroupAssignmentPostRequest{
		GroupId: "group_id" + postfix,
		AgentId: "agent_id" + postfix,
	}
}

func GenerateAgentGroupAssignmentPostResponse(
	id string,
	postfix string,
) config_api_client.AgentGroupAssignmentPostResponse {
	return config_api_client.AgentGroupAssignmentPostResponse{
		Id:    id,
		Group: *config_api_client.NewAgentGroupAssignmentPostGroup("group_id" + postfix),
		Agent: *config_api_client.NewAgentGroupAssignmentPostAgent("agent_id" + postfix),
		Type:  shared.AgentGroupAssignmentType,
	}
}

func GenerateAgentGroupAssignmentsGetResponse(
	id string,
	postfix string,
) config_api_client.AgentGroupAssignmentsGetResponse {
	return config_api_client.AgentGroupAssignmentsGetResponse{
		Items: []config_api_client.AgentGroupAssignmentsGetItem{
			{
				Id:    id,
				Group: *config_api_client.NewAgentGroupAssignmentsGetGroup("group_id" + postfix),
				Agent: *config_api_client.NewAgentGroupAssignmentsGetAgent("agent_id" + postfix),
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
	request config_api_client.AgentGroupAssignmentPostRequest,
	response config_api_client.AgentGroupAssignmentPostResponse,
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
