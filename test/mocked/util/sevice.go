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

func GenerateServiceTestResponse(
	id string,
	postfix string,
) config_api_client.ServiceTestsListResponse {
	return config_api_client.ServiceTestsListResponse{
		Items: []config_api_client.ServiceTestsListItem{
			{
				Id:        id,
				Category:  "external" + postfix,
				Name:      "name" + postfix,
				Target:    *config_api_client.NewNullableString(config_api_client.PtrString("target" + postfix)),
				Template:  "template" + postfix,
				IsEnabled: true,
				Type:      shared.ServiceTestType,
			},
		},
		Count: 1,
		Next:  *config_api_client.NewNullableString(nil),
	}
}

func MockGetServiceTest(id string, response interface{}, times int) {
	gock.New(MockUXIURL).
		Get(shared.ServiceTestPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}
