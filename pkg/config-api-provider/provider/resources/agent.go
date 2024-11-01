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
	_ resource.Resource              = &agentResource{}
	_ resource.ResourceWithConfigure = &agentResource{}
)

type agentResourceModel struct {
	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Notes    types.String `tfsdk:"notes"`
	PCapMode types.String `tfsdk:"pcap_mode"`
}

// TODO: Switch this to use the Client Model when that becomes available
type AgentResponseModel struct {
	UID                string
	Serial             string
	Name               string
	ModelNumber        string
	WifiMacAddress     string
	EthernetMacAddress string
	Notes              string
	PCapMode           string
}

// TODO: Switch this to use the Client Model when that becomes available
type AgentUpdateRequestModel struct {
	Name     string
	Notes    string
	PCapMode string
}

func NewAgentResource() resource.Resource {
	return &agentResource{}
}

type agentResource struct {
	client *config_api_client.APIClient
}

func (r *agentResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_agent"
}

func (r *agentResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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

func (r *agentResource) Configure(
	_ context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Resource type: Group. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = client
}

func (r *agentResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	// Retrieve values from plan
	var plan agentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"creating an agent is not supported; agents can only be imported",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *agentResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state agentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		AgentsGet(ctx).
		Id(state.ID.ValueString())
	sensorResponse, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_agent")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(sensorResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)
		return
	}
	sensor := sensorResponse.Items[0]

	state.Name = types.StringValue(sensor.Name)
	state.Notes = types.StringPointerValue(sensor.Notes.Get())
	state.PCapMode = types.StringPointerValue(sensor.PcapMode.Get())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *agentResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	// Retrieve values from plan
	var plan agentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing item
	response := UpdateAgent(AgentUpdateRequestModel{
		Name:     plan.Name.ValueString(),
		Notes:    plan.Notes.ValueString(),
		PCapMode: plan.PCapMode.ValueString(),
	})

	// Update resource state with updated items
	plan.Name = types.StringValue(response.Name)
	plan.Notes = types.StringValue(response.Notes)
	plan.PCapMode = types.StringValue(response.PCapMode)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *agentResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	// Retrieve values from state
	var state agentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.AgentsDelete(ctx, state.ID.ValueString())

	_, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(util.GenerateErrorSummary("delete", "uxi_agent"), errorDetail)
		return
	}
}

func (r *agentResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Update the agent using the configuration-api client
var UpdateAgent = func(request AgentUpdateRequestModel) AgentResponseModel {
	// TODO: Query the agent using the client

	return AgentResponseModel{
		UID:                "mock_uid",
		Serial:             "mock_serial",
		Name:               request.Name,
		ModelNumber:        "mock_model_number",
		WifiMacAddress:     "mock_wifi_mac_address",
		EthernetMacAddress: "mock_ethernet_mac_address",
		Notes:              request.Notes,
		PCapMode:           request.PCapMode,
	}
}
