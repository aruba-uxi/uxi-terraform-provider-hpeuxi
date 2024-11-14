package util

import (
	"context"
	"strconv"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
)

func GetWirelessNetwork(id string) config_api_client.WirelessNetworksItem {
	result, _, err := Client.ConfigurationAPI.
		WirelessNetworksGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("wireless_network with id `" + id + "` could not be found")
	}
	return result.Items[0]
}

func CheckStateAgainstWirelessNetwork(
	t st.Fatalf,
	wireless_network config_api_client.WirelessNetworksItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(
			"data.uxi_wireless_network.my_wireless_network",
			"id",
			config.WirelessNetworkId,
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wireless_network.my_wireless_network",
			"ssid",
			wireless_network.Ssid,
		),
		resource.TestCheckResourceAttrWith(
			"data.uxi_wireless_network.my_wireless_network",
			"name",
			func(value string) error {
				st.Assert(t, value, wireless_network.Name)
				return nil
			},
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wireless_network.my_wireless_network",
			"ip_version",
			wireless_network.IpVersion,
		),
		TestOptionalValue(
			t,
			"data.uxi_wireless_network.my_wireless_network",
			"security",
			wireless_network.Security.Get(),
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wireless_network.my_wireless_network",
			"hidden",
			strconv.FormatBool(wireless_network.Hidden),
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wireless_network.my_wireless_network",
			"band_locking",
			wireless_network.BandLocking,
		),
		TestOptionalValue(
			t,
			"data.uxi_wireless_network.my_wireless_network",
			"dns_lookup_domain",
			wireless_network.DnsLookupDomain.Get(),
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wireless_network.my_wireless_network",
			"disable_edns",
			strconv.FormatBool(wireless_network.DisableEdns),
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wireless_network.my_wireless_network",
			"use_dns64",
			strconv.FormatBool(wireless_network.UseDns64),
		),
		resource.TestCheckResourceAttr(
			"data.uxi_wireless_network.my_wireless_network",
			"external_connectivity",
			strconv.FormatBool(wireless_network.ExternalConnectivity),
		),
	)
}
