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

func GenerateServiceTestGroupAssignmentResponse(
	id string,
	postfix string,
) config_api_client.ServiceTestGroupAssignmentsResponse {
	return config_api_client.ServiceTestGroupAssignmentsResponse{
		Items: []config_api_client.ServiceTestGroupAssignmentsItem{
			{
				Id:          id,
				Group:       *config_api_client.NewGroup("group_id" + postfix),
				ServiceTest: *config_api_client.NewServiceTest("service_test_id" + postfix),
				Type:        shared.ServiceTestGroupAssignmentType,
			},
		},
		Count: 1,
		Next:  *config_api_client.NewNullableString(nil),
	}
}

func GenerateServiceTestGroupAssignmentPostRequest(
	id string,
	postfix string,
) config_api_client.ServiceTestGroupAssignmentsPostRequest {
	return config_api_client.ServiceTestGroupAssignmentsPostRequest{
		GroupId:       "group_id" + postfix,
		ServiceTestId: "service_test_id" + postfix,
	}
}

func GenerateServiceTestGroupAssignmentPostResponse(
	id string,
	postfix string,
) config_api_client.ServiceTestGroupAssignmentsPostResponse {
	return config_api_client.ServiceTestGroupAssignmentsPostResponse{
		Id:          id,
		Group:       *config_api_client.NewGroup("group_id" + postfix),
		ServiceTest: *config_api_client.NewServiceTest("service_test_id" + postfix),
		Type:        shared.ServiceTestGroupAssignmentType,
	}
}

func MockGetServiceTestGroupAssignment(id string, response interface{}, times int) {
	gock.New(MockUXIURL).
		Get(shared.ServiceTestGroupAssignmentPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPostServiceTestGroupAssignment(
	request config_api_client.ServiceTestGroupAssignmentsPostRequest,
	response config_api_client.ServiceTestGroupAssignmentsPostResponse,
	times int,
) {
	gock.New(MockUXIURL).
		Post(shared.ServiceTestGroupAssignmentPath).
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", mockToken).
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteServiceTestGroupAssignment(id string, times int) {
	gock.New(MockUXIURL).
		Delete(shared.ServiceTestGroupAssignmentPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}
