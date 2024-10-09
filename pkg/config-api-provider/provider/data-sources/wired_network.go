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
	_ datasource.DataSource              = &wiredNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &wiredNetworkDataSource{}
)

func NewWiredNetworkDataSource() datasource.DataSource {
	return &wiredNetworkDataSource{}
}

type wiredNetworkDataSource struct {
	client *config_api_client.APIClient
}

type wiredNetworkDataSourceModel struct {
	ID                   types.String `tfsdk:"id"`
	Alias                types.String `tfsdk:"alias"`
	IpVersion            types.String `tfsdk:"ip_version"`
	Security             types.String `tfsdk:"security"`
	DnsLookupDomain      types.String `tfsdk:"dns_lookup_domain"`
	DisableEdns          types.Bool   `tfsdk:"disable_edns"`
	UseDns64             types.Bool   `tfsdk:"use_dns64"`
	ExternalConnectivity types.Bool   `tfsdk:"external_connectivity"`
	VlanId               types.Int64  `tfsdk:"vlan_id"`

	Filter struct {
		WiredNetworkID string `tfsdk:"wired_network_id"`
	} `tfsdk:"filter"`
}

func (d *wiredNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wired_network"
}

func (d *wiredNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
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
			"vlan_id": schema.Int64Attribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"wired_network_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *wiredNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state wiredNetworkDataSourceModel

	// Read configuration from request
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		GetUxiV1alpha1WiredNetworksGet(ctx).
		Uid(state.Filter.WiredNetworkID)
	networkResponse, _, err := util.RetryFor429(request.Execute)

	if err != nil || len(networkResponse.Items) != 1 {
		resp.Diagnostics.AddError(
			"Error reading Wired Network",
			"Could not retrieve Wired Network, unexpected error: "+err.Error(),
		)
		return
	}

	network := networkResponse.Items[0]
	state.ID = types.StringValue(network.Id)
	state.Alias = types.StringValue(network.Name)
	state.IpVersion = types.StringValue(network.IpVersion)
	state.Security = types.StringValue(*network.Security.Get())
	state.DnsLookupDomain = types.StringValue(*network.DnsLookupDomain.Get())
	state.DisableEdns = types.BoolValue(network.DisableEdns)
	state.UseDns64 = types.BoolValue(network.UseDns64)
	state.ExternalConnectivity = types.BoolValue(network.ExternalConnectivity)
	state.VlanId = types.Int64Value(int64(*network.VLanId.Get()))

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *wiredNetworkDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Data Source type: Wired Network. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
