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

const MockRootGroupID = "root_group_id"

func GenerateGroupPostResponse(
	id string,
	nameSuffix string,
	parentIDSuffix string,
) config_api_client.GroupsPostResponse {
	parentID := "parent_id" + parentIDSuffix

	return config_api_client.GroupsPostResponse{
		Id:     id,
		Name:   "name" + nameSuffix,
		Parent: *config_api_client.NewParent(parentID),
		Path:   parentID + "." + id,
		Type:   shared.GroupType,
	}
}

func GenerateGroupAttachedToRootGroupPostResponse(
	id string,
	nameSuffix string,
) config_api_client.GroupsPostResponse {
	return config_api_client.GroupsPostResponse{
		Id:     id,
		Name:   "name" + nameSuffix,
		Parent: *config_api_client.NewParent(MockRootGroupID),
		Path:   "root_group_id." + id,
		Type:   shared.GroupType,
	}
}

func GenerateGroupPatchResponse(
	id string,
	nameSuffix string,
	parentIDSuffix string,
) config_api_client.GroupsPatchResponse {
	parentID := "parent_id" + parentIDSuffix

	return config_api_client.GroupsPatchResponse{
		Id:     id,
		Name:   "name" + nameSuffix,
		Parent: *config_api_client.NewParent(parentID),
		Path:   parentID + "." + id,
		Type:   shared.GroupType,
	}
}

func GenerateGroupGetResponse(
	id string,
	nameSuffix string,
	parentIDSuffix string,
) config_api_client.GroupsGetResponse {
	parentID := "parent_id" + parentIDSuffix

	return config_api_client.GroupsGetResponse{
		Items: []config_api_client.GroupsGetItem{
			{
				Id:     id,
				Name:   "name" + nameSuffix,
				Parent: *config_api_client.NewNullableParent(config_api_client.NewParent(parentID)),
				Path:   parentID + "." + id,
				Type:   shared.GroupType,
			},
		},
		Count: 1,
		Next:  *config_api_client.NewNullableString(nil),
	}
}

func GenerateGroupAttachedToRootGroupGetResponse(
	id string,
	nameSuffix string,
) config_api_client.GroupsGetResponse {
	return config_api_client.GroupsGetResponse{
		Items: []config_api_client.GroupsGetItem{
			{
				Id:     id,
				Name:   "name" + nameSuffix,
				Parent: *config_api_client.NewNullableParent(config_api_client.NewParent(MockRootGroupID)),
				Path:   "root_group_id." + id,
				Type:   shared.GroupType,
			},
		},
		Count: 1,
		Next:  *config_api_client.NewNullableString(nil),
	}
}

func GenerateRootGroupGetResponse() config_api_client.GroupsGetResponse {
	return config_api_client.GroupsGetResponse{
		Items: []config_api_client.GroupsGetItem{
			{
				Id:     MockRootGroupID,
				Name:   "root",
				Parent: *config_api_client.NewNullableParent(nil),
				Path:   MockRootGroupID,
				Type:   shared.GroupType,
			},
		},
		Count: 1,
		Next:  *config_api_client.NewNullableString(nil),
	}
}

func GenerateNonRootGroupPostRequest(
	id string,
	namePostfix string,
	parentIDPostfix string,
) config_api_client.GroupsPostRequest {
	parentID := "parent_id" + parentIDPostfix

	return config_api_client.GroupsPostRequest{
		Name:     "name" + namePostfix,
		ParentId: *config_api_client.NewNullableString(&parentID),
	}
}

func GenerateGroupAttachedToRootGroupPostRequest(
	id string,
	namePostfix string,
) config_api_client.GroupsPostRequest {
	return config_api_client.GroupsPostRequest{
		Name: "name" + namePostfix,
	}
}

func GenerateGroupPatchRequest(postfix string) config_api_client.GroupsPatchRequest {
	name := "name" + postfix

	return config_api_client.GroupsPatchRequest{
		Name: &name,
	}
}

func MockPostGroup(
	request config_api_client.GroupsPostRequest,
	response config_api_client.GroupsPostResponse,
	times int,
) {
	gock.New(MockUXIURL).
		Post("/networking-uxi/v1alpha1/groups").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", mockToken).
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockGetGroup(id string, response interface{}, times int) {
	gock.New(MockUXIURL).
		Get("/networking-uxi/v1alpha1/groups").
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockPatchGroup(
	id string,
	request config_api_client.GroupsPatchRequest,
	response config_api_client.GroupsPatchResponse,
	times int,
) {
	gock.New(MockUXIURL).
		Patch(shared.GroupPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		MatchHeader("Content-Type", "application/merge-patch+json").
		JSON(request).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteGroup(id string, times int) {
	gock.New(MockUXIURL).
		Delete(shared.GroupPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}
