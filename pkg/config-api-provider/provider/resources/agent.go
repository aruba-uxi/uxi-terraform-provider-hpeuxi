package resources

import (
	"context"

	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &agentResource{}
	_ resource.ResourceWithConfigure = &agentResource{}
)

type agentResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Notes       types.String `tfsdk:"notes"`
	PCapMode    types.String `tfsdk:"pcap_mode"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

type agentResponseModel struct {
	Uid                string
	Serial             string
	Name               string
	ModelNumber        string
	WifiMacAddress     string
	EthernetMacAddress string
	Notes              string
	PCapMode           string
}

func NewAgentResource() resource.Resource {
	return &agentResource{}
}

type agentResource struct{}

func (r *agentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_agent"
}

func (r *agentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"notes": schema.StringAttribute{
				Optional: true,
			},
			"pcap_mode": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *agentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
}

func (r *agentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan agentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "creating an agent resource is not supported; agents can only be imported")
	resp.Diagnostics.Append(diags...)
}

func (r *agentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state agentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client create-agent method
	// We are mocking the response of the client for this early stage of development
	response := GetAgent()

	// Update state from client response
	state.Name = types.StringValue(response.Name)
	state.Notes = types.StringValue(response.Notes)
	state.PCapMode = types.StringValue(response.PCapMode)
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *agentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan agentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client update-agent method

	// Update the state to match the plan (replace with response from client)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *agentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state agentResourceModel
	diags := req.State.Get(ctx, &state)
	diags.AddError("operation not supported", "deleting an agent resource is not supported; agents can only removed from state")
	resp.Diagnostics.Append(diags...)
}

func (r *agentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Get the agent using the configuration-api client
func GetAgent() agentResponseModel {
	// TODO: Query the agent using the client

	return agentResponseModel{
		Serial:             "temporary_serial",
		Name:               "temporary_name",
		ModelNumber:        "temporary_model_number",
		WifiMacAddress:     "temporary_wifi_mac_address",
		EthernetMacAddress: "temporary_ethernet_mac_address",
		Notes:              "temporary_notes",
		PCapMode:           "temporary_pcap_mode",
	}
}
