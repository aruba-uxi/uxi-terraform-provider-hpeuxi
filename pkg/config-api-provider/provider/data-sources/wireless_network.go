package datasources

import (
	"context"

	config_api_client "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/util"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
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
	Alias                types.String `tfsdk:"alias"`
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

func (d *wirelessNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_network"
}

func (d *wirelessNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ssid": schema.StringAttribute{
				Computed: true,
			},
			"alias": schema.StringAttribute{
				Computed: true,
			},
			"ip_version": schema.StringAttribute{
				Computed: true,
			},
			"security": schema.StringAttribute{
				Computed: true,
			},
			"hidden": schema.BoolAttribute{
				Computed: true,
			},
			"band_locking": schema.StringAttribute{
				Computed: true,
			},
			"dns_lookup_domain": schema.StringAttribute{
				Computed: true,
			},
			"disable_edns": schema.BoolAttribute{
				Computed: true,
			},
			"use_dns64": schema.BoolAttribute{
				Computed: true,
			},
			"external_connectivity": schema.BoolAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"wireless_network_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *wirelessNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state wirelessNetworkDataSourceModel

	// Read configuration from request
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		GetUxiV1alpha1WirelessNetworksGet(ctx).
		Uid(state.Filter.WirelessNetworkID)
	networkResponse, _, err := util.RetryFor429(request.Execute)

	if err != nil || len(networkResponse.WirelessNetworks) != 1 {
		resp.Diagnostics.AddError(
			"Error reading Wireless Network",
			"Could not retrieve Wireless Network, unexpected error: "+err.Error(),
		)
		return
	}

	network := networkResponse.WirelessNetworks[0]
	state.ID = types.StringValue(network.Uid)
	state.Ssid = types.StringValue(network.Ssid)
	state.Alias = types.StringValue(network.Alias)
	state.IpVersion = types.StringValue(network.IpVersion)
	state.Security = types.StringValue(*network.Security.Get())
	state.Hidden = types.BoolValue(network.Hidden)
	state.BandLocking = types.StringValue(network.BandLocking)
	state.DnsLookupDomain = types.StringValue(*network.DnsLookupDomain.Get())
	state.DisableEdns = types.BoolValue(network.DisableEdns)
	state.UseDns64 = types.BoolValue(network.UseDns64)
	state.ExternalConnectivity = types.BoolValue(network.ExternalConnectivity)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *wirelessNetworkDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
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
