/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package shared

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func CheckStateAgainstWiredNetwork(
	t *testing.T,
	entity string,
	wiredNetwork config_api_client.WiredNetworksItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", wiredNetwork.Id),
		resource.TestCheckResourceAttrWith(
			entity,
			"name",
			func(value string) error {
				assert.Equal(t, value, wiredNetwork.Name)

				return nil
			},
		),
		resource.TestCheckResourceAttr(entity, "ip_version", wiredNetwork.IpVersion),
		TestOptionalValue(t, entity, "security", wiredNetwork.Security.Get()),
		TestOptionalValue(t, entity, "dns_lookup_domain", wiredNetwork.DnsLookupDomain.Get()),
		resource.TestCheckResourceAttr(
			entity,
			"disable_edns",
			strconv.FormatBool(wiredNetwork.DisableEdns),
		),
		resource.TestCheckResourceAttr(
			entity,
			"use_dns64",
			strconv.FormatBool(wiredNetwork.UseDns64),
		),
		resource.TestCheckResourceAttr(
			entity,
			"external_connectivity",
			strconv.FormatBool(wiredNetwork.ExternalConnectivity),
		),
		TestOptionalValue(
			t,
			entity,
			"vlan_id",
			Int32PtrToStringPtr(wiredNetwork.VLanId.Get()),
		),
	)
}
