/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateWiredNetworkResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":                   id,
		"name":                 "name" + postfix,
		"createdAt":            "2024-09-11T12:00:00.000Z",
		"updatedAt":            "2024-09-11T12:00:00.000Z",
		"ipVersion":            "ip_version" + postfix,
		"security":             "security" + postfix,
		"dnsLookupDomain":      "dns_lookup_domain" + postfix,
		"disableEdns":          false,
		"useDns64":             false,
		"externalConnectivity": false,
		"vLanId":               123,
		"type":                 shared.WiredNetworkType,
	}
}

func MockGetWiredNetwork(id string, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Get(shared.WiredNetworkPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}
