package resources

import (
	"context"
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"

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
				Computed: true,
			},
			"pcap_mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
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
	agentResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_agent")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(agentResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)
		return
	}
	agent := agentResponse.Items[0]

	state.ID = types.StringValue(agent.Id)
	state.Name = types.StringValue(agent.Name)
	state.Notes = types.StringPointerValue(agent.Notes.Get())
	state.PCapMode = types.StringPointerValue(agent.PcapMode.Get())

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

	patchRequest := config_api_client.NewAgentsPatchRequest()
	patchRequest.Name = plan.Name.ValueStringPointer()
	if !plan.Notes.IsUnknown() {
		patchRequest.Notes = plan.Notes.ValueStringPointer()
	}
	if !plan.PCapMode.IsUnknown() {
		patchRequest.PcapMode = plan.PCapMode.ValueStringPointer()
	}
	request := r.client.ConfigurationAPI.
		AgentsPatch(ctx, plan.ID.ValueString()).
		AgentsPatchRequest(*patchRequest)
	agent, response, err := util.RetryForTooManyRequests(request.Execute)

	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(util.GenerateErrorSummary("update", "uxi_agent"), errorDetail)
		return
	}

	// Update the state to match the plan (replace with response from client)
	plan.ID = types.StringValue(agent.Id)
	plan.Name = types.StringValue(agent.Name)
	if agent.Notes.Get() != nil {
		plan.Notes = types.StringValue(*agent.Notes.Get())
	}
	if agent.PcapMode.Get() != nil {
		plan.PCapMode = types.StringValue(*agent.PcapMode.Get())
	}

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

	_, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		if response.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
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
