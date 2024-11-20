/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package shared

import (
	"strconv"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
)

func CheckStateAgainstWiredNetwork(
	t *testing.T,
	entity string,
	wired_network config_api_client.WiredNetworksItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", config.WiredNetworkId),
		resource.TestCheckResourceAttrWith(
			entity,
			"name",
			func(value string) error {
				assert.Equal(t, value, wired_network.Name)
				return nil
			},
		),
		resource.TestCheckResourceAttr(entity, "ip_version", wired_network.IpVersion),
		TestOptionalValue(t, entity, "security", wired_network.Security.Get()),
		TestOptionalValue(t, entity, "dns_lookup_domain", wired_network.DnsLookupDomain.Get()),
		resource.TestCheckResourceAttr(
			entity,
			"disable_edns",
			strconv.FormatBool(wired_network.DisableEdns),
		),
		resource.TestCheckResourceAttr(
			entity,
			"use_dns64",
			strconv.FormatBool(wired_network.UseDns64),
		),
		resource.TestCheckResourceAttr(
			entity,
			"external_connectivity",
			strconv.FormatBool(wired_network.ExternalConnectivity),
		),
		TestOptionalValue(
			t,
			entity,
			"vlan_id",
			Int32PtrToStringPtr(wired_network.VLanId.Get()),
		),
	)
}
