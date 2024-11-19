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
	Name                 types.String `tfsdk:"name"`
	IpVersion            types.String `tfsdk:"ip_version"`
	Security             types.String `tfsdk:"security"`
	DnsLookupDomain      types.String `tfsdk:"dns_lookup_domain"`
	DisableEdns          types.Bool   `tfsdk:"disable_edns"`
	UseDns64             types.Bool   `tfsdk:"use_dns64"`
	ExternalConnectivity types.Bool   `tfsdk:"external_connectivity"`
	VlanId               types.Int32  `tfsdk:"vlan_id"`

	Filter struct {
		WiredNetworkID string `tfsdk:"wired_network_id"`
	} `tfsdk:"filter"`
}

func (d *wiredNetworkDataSource) Metadata(
	_ context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_wired_network"
}

func (d *wiredNetworkDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Retrieves a specific wired network.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the wired network.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the wired network.",
				Computed:    true,
			},
			"ip_version": schema.StringAttribute{
				Description: "The ip version of the wired network.",
				Computed:    true,
			},
			"security": schema.StringAttribute{
				Description: "The security protocol of the wired network.",
				Computed:    true,
			},
			"dns_lookup_domain": schema.StringAttribute{
				Description: "The DNS lookup domain of the wired network.",
				Computed:    true,
			},
			"disable_edns": schema.BoolAttribute{
				Description: "Whether EDNS is disabled on the wired network.",
				Computed:    true,
			},
			"use_dns64": schema.BoolAttribute{
				Description: "Whether the wired network is configured to use DNS64.",
				Computed:    true,
			},
			"external_connectivity": schema.BoolAttribute{
				Description: "Whether the wired network has external connectivity.",
				Computed:    true,
			},
			"vlan_id": schema.Int32Attribute{
				Description: "The VLAN identifier of the wired network.",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Description: "The filter used to filter the specific wired network.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"wired_network_id": schema.StringAttribute{
						Description: "The identifier of the wired network.",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *wiredNetworkDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var state wiredNetworkDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		WiredNetworksGet(ctx).
		Id(state.Filter.WiredNetworkID)
	networkResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_wired_network")

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
	state.Name = types.StringValue(network.Name)
	state.IpVersion = types.StringValue(network.IpVersion)
	state.Security = types.StringPointerValue(network.Security.Get())
	state.DnsLookupDomain = types.StringPointerValue(network.DnsLookupDomain.Get())
	state.DisableEdns = types.BoolValue(network.DisableEdns)
	state.UseDns64 = types.BoolValue(network.UseDns64)
	state.ExternalConnectivity = types.BoolValue(network.ExternalConnectivity)
	state.VlanId = types.Int32PointerValue(network.VLanId.Get())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *wiredNetworkDataSource) Configure(
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
			"Data Source type: Wired Network. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
