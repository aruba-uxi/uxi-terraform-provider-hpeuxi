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

func GenerateNetworkGroupAssignmentsGetResponse(
	id string,
	postfix string,
) config_api_client.NetworkGroupAssignmentsGetResponse {
	return config_api_client.NetworkGroupAssignmentsGetResponse{
		Items: []config_api_client.NetworkGroupAssignmentsGetItem{
			{
				Id:      id,
				Group:   *config_api_client.NewNetworkGroupAssignmentsGetGroup("group_id" + postfix),
				Network: *config_api_client.NewNetworkGroupAssignmentsGetNetwork("network_id" + postfix),
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
) config_api_client.NetworkGroupAssignmentPostRequest {
	return config_api_client.NetworkGroupAssignmentPostRequest{
		GroupId:   *config_api_client.PtrString("group_id" + postfix),
		NetworkId: *config_api_client.PtrString("network_id" + postfix),
	}
}

func GenerateNetworkGroupAssignmentPostResponse(
	id string,
	postfix string,
) config_api_client.NetworkGroupAssignmentPostResponse {
	return config_api_client.NetworkGroupAssignmentPostResponse{
		Id:      id,
		Group:   *config_api_client.NewNetworkGroupAssignmentPostGroup("group_id" + postfix),
		Network: *config_api_client.NewNetworkGroupAssignmentPostNetwork("network_id" + postfix),
		Type:    shared.NetworkGroupAssignmentType,
	}
}

func MockGetNetworkGroupAssignment(id string, response interface{}, times int) {
	gock.New(MockUXIURL).
		Get(shared.NetworkGroupAssignmentPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostNetworkGroupAssignment(
	request config_api_client.NetworkGroupAssignmentPostRequest,
	response config_api_client.NetworkGroupAssignmentPostResponse,
	times int,
) {
	gock.New(MockUXIURL).
		Post(shared.NetworkGroupAssignmentPath).
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", mockToken).
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteNetworkGroupAssignment(id string, times int) {
	gock.New(MockUXIURL).
		Delete(shared.NetworkGroupAssignmentPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}
