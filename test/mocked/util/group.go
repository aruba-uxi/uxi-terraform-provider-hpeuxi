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
) config_api_client.GroupPostResponse {
	parentID := "parent_id" + parentIDSuffix

	return config_api_client.GroupPostResponse{
		Id:     id,
		Name:   "name" + nameSuffix,
		Parent: *config_api_client.NewGroupPostParent(parentID),
		Path:   parentID + "." + id,
		Type:   shared.GroupType,
	}
}

func GenerateGroupAttachedToRootGroupPostResponse(
	id string,
	nameSuffix string,
) config_api_client.GroupPostResponse {
	return config_api_client.GroupPostResponse{
		Id:     id,
		Name:   "name" + nameSuffix,
		Parent: *config_api_client.NewGroupPostParent(MockRootGroupID),
		Path:   "root_group_id." + id,
		Type:   shared.GroupType,
	}
}

func GenerateGroupPatchResponse(
	id string,
	nameSuffix string,
	parentIDSuffix string,
) config_api_client.GroupPatchResponse {
	parentID := "parent_id" + parentIDSuffix

	return config_api_client.GroupPatchResponse{
		Id:     id,
		Name:   "name" + nameSuffix,
		Parent: *config_api_client.NewGroupPatchParent(parentID),
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
				Parent: *config_api_client.NewNullableGroupsGetParent(config_api_client.NewGroupsGetParent(parentID)),
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
				Parent: *config_api_client.NewNullableGroupsGetParent(config_api_client.NewGroupsGetParent(MockRootGroupID)),
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
				Parent: *config_api_client.NewNullableGroupsGetParent(nil),
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
) config_api_client.GroupPostRequest {
	parentID := "parent_id" + parentIDPostfix

	return config_api_client.GroupPostRequest{
		Name:     "name" + namePostfix,
		ParentId: *config_api_client.NewNullableString(&parentID),
	}
}

func GenerateGroupAttachedToRootGroupPostRequest(
	id string,
	namePostfix string,
) config_api_client.GroupPostRequest {
	return config_api_client.GroupPostRequest{
		Name: "name" + namePostfix,
	}
}

func GenerateGroupPatchRequest(postfix string) config_api_client.GroupPatchRequest {
	name := "name" + postfix

	return config_api_client.GroupPatchRequest{
		Name: &name,
	}
}

func MockPostGroup(
	request config_api_client.GroupPostRequest,
	response config_api_client.GroupPostResponse,
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
	request config_api_client.GroupPatchRequest,
	response config_api_client.GroupPatchResponse,
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
