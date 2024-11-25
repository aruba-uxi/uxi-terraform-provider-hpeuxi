/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package shared

import (
	"strconv"
	"testing"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
)

func CheckStateAgainstWirelessNetwork(
	t *testing.T,
	entity string,
	wirelessNetwork config_api_client.WirelessNetworksItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", wirelessNetwork.Id),
		resource.TestCheckResourceAttr(entity, "ssid", wirelessNetwork.Ssid),
		resource.TestCheckResourceAttrWith(
			entity,
			"name",
			func(value string) error {
				assert.Equal(t, value, wirelessNetwork.Name)
				return nil
			},
		),
		resource.TestCheckResourceAttr(entity, "ip_version", wirelessNetwork.IpVersion),
		TestOptionalValue(
			t,
			entity,
			"security",
			wirelessNetwork.Security.Get(),
		),
		resource.TestCheckResourceAttr(
			entity,
			"hidden",
			strconv.FormatBool(wirelessNetwork.Hidden),
		),
		resource.TestCheckResourceAttr(entity, "band_locking", wirelessNetwork.BandLocking),
		TestOptionalValue(
			t,
			entity,
			"dns_lookup_domain",
			wirelessNetwork.DnsLookupDomain.Get(),
		),
		resource.TestCheckResourceAttr(
			entity,
			"disable_edns",
			strconv.FormatBool(wirelessNetwork.DisableEdns),
		),
		resource.TestCheckResourceAttr(
			entity,
			"use_dns64",
			strconv.FormatBool(wirelessNetwork.UseDns64),
		),
		resource.TestCheckResourceAttr(
			entity,
			"external_connectivity",
			strconv.FormatBool(wirelessNetwork.ExternalConnectivity),
		),
	)
}
