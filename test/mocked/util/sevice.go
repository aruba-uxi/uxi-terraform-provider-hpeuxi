/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateServiceTestResponse(
	id string,
	postfix string,
) map[string]interface{} {
	return map[string]interface{}{
		"id":        id,
		"category":  "external" + postfix,
		"name":      "name" + postfix,
		"target":    "target" + postfix,
		"template":  "template" + postfix,
		"isEnabled": true,
		"type":      shared.ServiceTestType,
	}
}

func MockGetServiceTest(id string, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Get(shared.ServiceTestPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}
