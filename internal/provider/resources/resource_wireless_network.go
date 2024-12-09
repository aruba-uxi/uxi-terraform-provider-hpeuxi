/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/util"
	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

var (
	_ resource.Resource              = &wirelessNetworkResource{}
	_ resource.ResourceWithConfigure = &wirelessNetworkResource{}
)

type wirelessNetworkResourceModel struct {
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
}

func NewWirelessNetworkResource() resource.Resource {
	return &wirelessNetworkResource{}
}

type wirelessNetworkResource struct {
	client *config_api_client.APIClient
}

func (r *wirelessNetworkResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_wireless_network"
}

func (r *wirelessNetworkResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Manages a wireless network.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the wireless network.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ssid": schema.StringAttribute{
				Description: "The SSID of the wireless network.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the wireless network.",
				Required:    true,
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
		},
	}
}

func (r *wirelessNetworkResource) Configure(
	_ context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Resource type: Wireless Network. Please report this issue to the provider developers.",
		)

		return
	}

	r.client = client
}

func (r *wirelessNetworkResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan wirelessNetworkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"creating a wireless_network is not supported; wireless_networks can only be imported",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *wirelessNetworkResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state wirelessNetworkResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		WirelessNetworksGet(ctx).
		Id(state.ID.ValueString())
	networkResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	errorSummary := util.GenerateErrorSummary("read", "hpeuxi_wireless_network")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)

		return
	}

	defer response.Body.Close()

	if len(networkResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)

		return
	}

	network := networkResponse.Items[0]

	state.ID = types.StringValue(network.Id)
	state.SSID = types.StringValue(network.Ssid)
	state.Name = types.StringValue(network.Name)
	state.IPVersion = types.StringValue(string(network.IpVersion))
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

func (r *wirelessNetworkResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan wirelessNetworkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"updating a wireless_network is not supported; wireless_networks can only be updated through the dashboard",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *wirelessNetworkResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state wirelessNetworkResourceModel
	diags := req.State.Get(ctx, &state)
	diags.AddError(
		"operation not supported",
		"deleting a wireless_network is not supported; wireless_networks can only removed from state",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *wirelessNetworkResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
