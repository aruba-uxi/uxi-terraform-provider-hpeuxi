/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"context"
	"strconv"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
)

func GetWiredNetwork(id string) config_api_client.WiredNetworksItem {
	result, _, err := Client.ConfigurationAPI.
		WiredNetworksGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("wired_network with id `" + id + "` could not be found")
	}
	return result.Items[0]
}

func CheckStateAgainstWiredNetwork(
	t *testing.T,
	wired_network config_api_client.WiredNetworksItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(
			"data.uxi_wired_network.my_wired_network",
			"id",
			config.WiredNetworkId,
		),
		resource.TestCheckResourceAttrWith(
			"data.uxi_wired_network.my_wired_network",
			"name",
			func(value string) error {
				assert.Equal(t, value, wired_network.Name)
				return nil
			},
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wired_network.my_wired_network",
			"ip_version",
			wired_network.IpVersion,
		),
		TestOptionalValue(
			t,
			"data.uxi_wired_network.my_wired_network",
			"security",
			wired_network.Security.Get(),
		),
		TestOptionalValue(
			t,
			"data.uxi_wired_network.my_wired_network",
			"dns_lookup_domain",
			wired_network.DnsLookupDomain.Get(),
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wired_network.my_wired_network",
			"disable_edns",
			strconv.FormatBool(wired_network.DisableEdns),
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wired_network.my_wired_network",
			"use_dns64",
			strconv.FormatBool(wired_network.UseDns64),
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wired_network.my_wired_network",
			"external_connectivity",
			strconv.FormatBool(wired_network.ExternalConnectivity),
		),
		TestOptionalValue(
			t,
			"data.uxi_wired_network.my_wired_network",
			"vlan_id",
			Int32PtrToStringPtr(wired_network.VLanId.Get()),
		),
	)
}
