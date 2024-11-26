/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"
	"time"

	"github.com/h2non/gock"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
)

func GenerateWiredNetworkResponse(
	id string,
	postfix string,
) config_api_client.WiredNetworksResponse {
	createdAt, _ := time.Parse(time.RFC3339, "2024-09-11T12:00:00.000Z")
	updatedAt, _ := time.Parse(time.RFC3339, "2024-09-11T12:00:00.000Z")

	return config_api_client.WiredNetworksResponse{
		Items: []config_api_client.WiredNetworksItem{
			{
				Id:                   id,
				Name:                 "name" + postfix,
				CreatedAt:            createdAt,
				UpdatedAt:            updatedAt,
				IpVersion:            "ip_version" + postfix,
				Security:             *config_api_client.NewNullableString(config_api_client.PtrString("security" + postfix)),
				DnsLookupDomain:      *config_api_client.NewNullableString(config_api_client.PtrString("dns_lookup_domain" + postfix)),
				DisableEdns:          false,
				UseDns64:             false,
				ExternalConnectivity: false,
				VLanId:               *config_api_client.NewNullableInt32(config_api_client.PtrInt32(123)),
				Type:                 shared.WiredNetworkType,
			},
		},
	}
}

func MockGetWiredNetwork(id string, response interface{}, times int) {
	gock.New(MockUXIURL).
		Get(shared.WiredNetworkPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}
