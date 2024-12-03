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

func GenerateServiceTestGroupAssignmentsGetResponse(
	id string,
	postfix string,
) config_api_client.ServiceTestGroupAssignmentsGetResponse {
	return config_api_client.ServiceTestGroupAssignmentsGetResponse{
		Items: []config_api_client.ServiceTestGroupAssignmentsGetItem{
			{
				Id:          id,
				Group:       *config_api_client.NewServiceTestGroupAssignmentsGetGroup("group_id" + postfix),
				ServiceTest: *config_api_client.NewServiceTestGroupAssignmentsGetServiceTest("service_test_id" + postfix),
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
) config_api_client.ServiceTestGroupAssignmentPostRequest {
	return config_api_client.ServiceTestGroupAssignmentPostRequest{
		GroupId:       "group_id" + postfix,
		ServiceTestId: "service_test_id" + postfix,
	}
}

func GenerateServiceTestGroupAssignmentPostResponse(
	id string,
	postfix string,
) config_api_client.ServiceTestGroupAssignmentPostResponse {
	return config_api_client.ServiceTestGroupAssignmentPostResponse{
		Id:          id,
		Group:       *config_api_client.NewServiceTestGroupAssignmentPostGroup("group_id" + postfix),
		ServiceTest: *config_api_client.NewServiceTestGroupAssignmentPostServiceTest("service_test_id" + postfix),
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
	request config_api_client.ServiceTestGroupAssignmentPostRequest,
	response config_api_client.ServiceTestGroupAssignmentPostResponse,
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
