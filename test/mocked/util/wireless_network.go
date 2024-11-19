/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/shared"
	"github.com/h2non/gock"
)

func GenerateWirelessNetworkResponse(id string, postfix string) map[string]interface{} {
	return map[string]interface{}{
		"id":                   id,
		"ssid":                 "ssid" + postfix,
		"createdAt":            "2024-09-11T12:00:00.000Z",
		"updatedAt":            "2024-09-11T12:00:00.000Z",
		"name":                 "name" + postfix,
		"ipVersion":            "ip_version" + postfix,
		"security":             "security" + postfix,
		"hidden":               false,
		"bandLocking":          "band_locking" + postfix,
		"dnsLookupDomain":      "dns_lookup_domain" + postfix,
		"disableEdns":          false,
		"useDns64":             false,
		"externalConnectivity": false,
		"type":                 shared.WirelessNetworkType,
	}
}

func MockGetWirelessNetwork(id string, response map[string]interface{}, times int) {
	gock.New(MockUxiUrl).
		Get(shared.WirelessNetworkPath).
		MatchHeader("Authorization", mockToken).
		MatchParam("id", id).
		Times(times).
		Reply(http.StatusOK).
		JSON(response)
}
