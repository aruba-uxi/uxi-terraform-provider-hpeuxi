/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateNonRootGroupResponse(
	id string,
	nonReplacementFieldPostfix string,
	replacementFieldPostfix string,
) map[string]interface{} {
	parentId := "parent_id" + replacementFieldPostfix

	return map[string]interface{}{
		"id":     id,
		"name":   "name" + nonReplacementFieldPostfix,
		"parent": map[string]string{"id": parentId},
		"path":   parentId + "." + id,
		"type":   shared.GroupType,
	}
}

func GenerateGroupRequest(
	id string,
	nonReplacementFieldPostfix string,
	replacementFieldPostfix string,
) map[string]interface{} {
	return map[string]interface{}{
		"name":     "name" + nonReplacementFieldPostfix,
		"parentId": "parent_id" + replacementFieldPostfix,
	}
}

func MockPostGroup(request map[string]interface{}, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Post("/networking-uxi/v1alpha1/groups").
		MatchHeader("Content-Type", "application/json").
		MatchHeader("Authorization", mockToken).
		Times(times).
		JSON(request).
		Reply(http.StatusOK).
		JSON(response)
}

func MockGetGroup(id string, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Get("/networking-uxi/v1alpha1/groups").
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockUpdateGroup(
	id string,
	request map[string]interface{},
	response map[string]interface{},
	times int,
) {
	gock.New(MockUxiUrl).
		Patch(shared.GroupPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		MatchHeader("Content-Type", "application/merge-patch+json").
		JSON(request).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}

func MockDeleteGroup(id string, times int) {
	gock.New(MockUxiUrl).
		Delete(shared.GroupPath+"/"+id).
		MatchHeader("Authorization", mockToken).
		Times(times).
		Reply(http.StatusNoContent)
}
