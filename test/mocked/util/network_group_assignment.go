/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateNetworkGroupAssignmentResponse(
	id string,
	postfix string,
) config_api_client.NetworkGroupAssignmentsResponse {
	return config_api_client.NetworkGroupAssignmentsResponse{
		Items: []config_api_client.NetworkGroupAssignmentsItem{
			{
				Id:      id,
				Group:   *config_api_client.NewGroup("group_id" + postfix),
				Network: *config_api_client.NewNetwork("network_id" + postfix),
				Type:    shared.NetworkGroupAssignmentType,
			},
		},
		Count: 1,
		Next:  *config_api_client.NewNullableString(nil),
	}
}

func GenerateNetworkGroupAssignmentPostRequest(
	id string,
	postfix string,
) config_api_client.NetworkGroupAssignmentsPostRequest {
	return config_api_client.NetworkGroupAssignmentsPostRequest{
		GroupId:   *config_api_client.PtrString("group_id" + postfix),
		NetworkId: *config_api_client.PtrString("network_id" + postfix),
	}
}

func GenerateNetworkGroupAssignmentPostResponse(
	id string,
	postfix string,
) config_api_client.NetworkGroupAssignmentsPostResponse {
	return config_api_client.NetworkGroupAssignmentsPostResponse{
		Id:      id,
		Group:   *config_api_client.NewGroup("group_id" + postfix),
		Network: *config_api_client.NewNetwork("network_id" + postfix),
		Type:    shared.NetworkGroupAssignmentType,
	}
}

func MockGetNetworkGroupAssignment(id string, response interface{}, times int) {
	gock.New(MockUxiUrl).
		Get(shared.NetworkGroupAssignmentPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostNetworkGroupAssignment(
	request config_api_client.NetworkGroupAssignmentsPostRequest,
	response config_api_client.NetworkGroupAssignmentsPostResponse,
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
