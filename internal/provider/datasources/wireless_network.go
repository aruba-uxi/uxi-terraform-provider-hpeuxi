/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package datasources

import (
	"context"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &wirelessNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &wirelessNetworkDataSource{}
)

func NewWirelessNetworkDataSource() datasource.DataSource {
	return &wirelessNetworkDataSource{}
}

type wirelessNetworkDataSource struct {
	client *config_api_client.APIClient
}

type wirelessNetworkDataSourceModel struct {
	ID                   types.String `tfsdk:"id"`
	Ssid                 types.String `tfsdk:"ssid"`
	Name                 types.String `tfsdk:"name"`
	IpVersion            types.String `tfsdk:"ip_version"`
	Security             types.String `tfsdk:"security"`
	Hidden               types.Bool   `tfsdk:"hidden"`
	BandLocking          types.String `tfsdk:"band_locking"`
	DnsLookupDomain      types.String `tfsdk:"dns_lookup_domain"`
	DisableEdns          types.Bool   `tfsdk:"disable_edns"`
	UseDns64             types.Bool   `tfsdk:"use_dns64"`
	ExternalConnectivity types.Bool   `tfsdk:"external_connectivity"`

	Filter struct {
		WirelessNetworkID string `tfsdk:"wireless_network_id"`
	} `tfsdk:"filter"`
}

func (d *wirelessNetworkDataSource) Metadata(
	_ context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_wireless_network"
}

func (d *wirelessNetworkDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The identifier of the wireless network.",
			},
			"ssid": schema.StringAttribute{
				Computed:    true,
				Description: "The SSID of the wireless network.",
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the wireless network.",
			},
			"ip_version": schema.StringAttribute{
				Computed:    true,
				Description: "The IP version of the wireless network.",
			},
			"security": schema.StringAttribute{
				Computed:    true,
				Description: "The security protocol of the wireless network.",
			},
			"hidden": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the wireless network is hidden.",
			},
			"band_locking": schema.StringAttribute{
				Computed:    true,
				Description: "The frequency band the wireless network is locked to.",
			},
			"dns_lookup_domain": schema.StringAttribute{
				Computed:    true,
				Description: "The DNS lookup domain of the wireless network.",
			},
			"disable_edns": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether EDNS is disabled on the wireless network.",
			},
			"use_dns64": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the wireless network is configured to use DNS64.",
			},
			"external_connectivity": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the wireless network has external connectivity.",
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"wireless_network_id": schema.StringAttribute{
						Required: true,
					},
				},
				Description: "The filter used to filter the specific wired network by id.",
			},
		},
	}
}

func (d *wirelessNetworkDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var state wirelessNetworkDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		WirelessNetworksGet(ctx).
		Id(state.Filter.WirelessNetworkID)
	networkResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_wireless_network")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(networkResponse.Items) != 1 {
		resp.Diagnostics.AddError(errorSummary, "Could not find specified data source")
		resp.State.RemoveResource(ctx)
		return
	}

	network := networkResponse.Items[0]
	state.ID = types.StringValue(network.Id)
	state.Ssid = types.StringValue(network.Ssid)
	state.Name = types.StringValue(network.Name)
	state.IpVersion = types.StringValue(network.IpVersion)
	state.Security = types.StringPointerValue(network.Security.Get())
	state.Hidden = types.BoolValue(network.Hidden)
	state.BandLocking = types.StringValue(network.BandLocking)
	state.DnsLookupDomain = types.StringPointerValue(network.DnsLookupDomain.Get())
	state.DisableEdns = types.BoolValue(network.DisableEdns)
	state.UseDns64 = types.BoolValue(network.UseDns64)
	state.ExternalConnectivity = types.BoolValue(network.ExternalConnectivity)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *wirelessNetworkDataSource) Configure(
	_ context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Data Source type: Wireless Network. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
