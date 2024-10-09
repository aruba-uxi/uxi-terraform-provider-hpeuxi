package resources

import (
	"context"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/util"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &wirelessNetworkResource{}
	_ resource.ResourceWithConfigure = &wirelessNetworkResource{}
)

type wirelessNetworkResourceModel struct {
	ID    types.String `tfsdk:"id"`
	Alias types.String `tfsdk:"alias"`
}

func NewWirelessNetworkResource() resource.Resource {
	return &wirelessNetworkResource{}
}

type wirelessNetworkResource struct {
	client *config_api_client.APIClient
}

func (r *wirelessNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_network"
}

func (r *wirelessNetworkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"alias": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *wirelessNetworkResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
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

func (r *wirelessNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan wirelessNetworkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "creating a wireless_network is not supported; wireless_networks can only be imported")
	resp.Diagnostics.Append(diags...)
}

func (r *wirelessNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state wirelessNetworkResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		GetUxiV1alpha1WirelessNetworksGet(ctx).
		Uid(state.ID.ValueString())
	networkResponse, _, err := util.RetryFor429(request.Execute)

	if err != nil || len(networkResponse.Items) != 1 {
		resp.Diagnostics.AddError(
			"Error reading Wireless Networks",
			"Could not retrieve Wireless Network, unexpected error: "+err.Error(),
		)
		return
	}

	network := networkResponse.Items[0]

	// Update state from client response
	state.Alias = types.StringValue(network.Name)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *wirelessNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan wirelessNetworkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "updating a wireless_network is not supported; wireless_networks can only be updated through the dashboard")
	resp.Diagnostics.Append(diags...)
}

func (r *wirelessNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state wirelessNetworkResourceModel
	diags := req.State.Get(ctx, &state)
	diags.AddError("operation not supported", "deleting a wireless_network is not supported; wireless_networks can only removed from state")
	resp.Diagnostics.Append(diags...)
}

func (r *wirelessNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
