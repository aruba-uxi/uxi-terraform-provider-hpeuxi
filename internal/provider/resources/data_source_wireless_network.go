/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

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
	SSID                 types.String `tfsdk:"ssid"`
	Name                 types.String `tfsdk:"name"`
	IPVersion            types.String `tfsdk:"ip_version"`
	Security             types.String `tfsdk:"security"`
	Hidden               types.Bool   `tfsdk:"hidden"`
	BandLocking          types.String `tfsdk:"band_locking"`
	DNSLookupDomain      types.String `tfsdk:"dns_lookup_domain"`
	DisableEDNS          types.Bool   `tfsdk:"disable_edns"`
	UseDNS64             types.Bool   `tfsdk:"use_dns64"`
	ExternalConnectivity types.Bool   `tfsdk:"external_connectivity"`
	Filter               struct {
		ID string `tfsdk:"id"`
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
		Description: "Retrieves a specific wireless network.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the wireless network.",
				Computed:    true,
			},
			"ssid": schema.StringAttribute{
				Description: "The SSID of the wireless network.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the wireless network.",
				Computed:    true,
			},
			"ip_version": schema.StringAttribute{
				Description: "The IP version of the wireless network.",
				Computed:    true,
			},
			"security": schema.StringAttribute{
				Description: "The security protocol of the wireless network.",
				Computed:    true,
			},
			"hidden": schema.BoolAttribute{
				Description: "Whether the wireless network is hidden.",
				Computed:    true,
			},
			"band_locking": schema.StringAttribute{
				Description: "The frequency band the wireless network is locked to.",
				Computed:    true,
			},
			"dns_lookup_domain": schema.StringAttribute{
				Description: "The DNS lookup domain of the wireless network.",
				Computed:    true,
			},
			"disable_edns": schema.BoolAttribute{
				Description: "Whether EDNS is disabled on the wireless network.",
				Computed:    true,
			},
			"use_dns64": schema.BoolAttribute{
				Description: "Whether the wireless network is configured to use DNS64.",
				Computed:    true,
			},
			"external_connectivity": schema.BoolAttribute{
				Description: "Whether the wireless network has external connectivity.",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Description: "The filter used to filter the specific wireless network.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The identifier of the wireless network.",
						Required:    true,
					},
				},
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
		Id(state.Filter.ID)
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
	state.SSID = types.StringValue(network.Ssid)
	state.Name = types.StringValue(network.Name)
	state.IPVersion = types.StringValue(network.IpVersion)
	state.Security = types.StringPointerValue(network.Security.Get())
	state.Hidden = types.BoolValue(network.Hidden)
	state.BandLocking = types.StringValue(network.BandLocking)
	state.DNSLookupDomain = types.StringPointerValue(network.DnsLookupDomain.Get())
	state.DisableEDNS = types.BoolValue(network.DisableEdns)
	state.UseDNS64 = types.BoolValue(network.UseDns64)
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
