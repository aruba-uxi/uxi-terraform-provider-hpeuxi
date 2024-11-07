package resources

import (
	"context"
	"github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-configuration/internal/provider/util"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &agentGroupAssignmentResource{}
	_ resource.ResourceWithConfigure = &agentGroupAssignmentResource{}
)

type agentGroupAssignmentResourceModel struct {
	ID      types.String `tfsdk:"id"`
	AgentID types.String `tfsdk:"agent_id"`
	GroupID types.String `tfsdk:"group_id"`
}

type AgentGroupAssignmentResponseModel struct {
	UID      string //  <assignment_uid>
	GroupUID string //  <group_uid:str>,
	AgentUID string //  <agent_uid:str>
}

type AgentGroupAssignmentRequestModel struct {
	GroupUID string //  <group_uid:str>,
	AgentUID string //  <agent_uid:str>
}

func NewAgentGroupAssignmentResource() resource.Resource {
	return &agentGroupAssignmentResource{}
}

type agentGroupAssignmentResource struct {
	client *config_api_client.APIClient
}

func (r *agentGroupAssignmentResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_agent_group_assignment"
}

func (r *agentGroupAssignmentResource) Schema(
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
			"agent_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"group_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *agentGroupAssignmentResource) Configure(
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
			"Resource type: Network Group Assignment. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = client
}

func (r *agentGroupAssignmentResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	// Retrieve values from plan
	var plan agentGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client createAgentGroupAssignment method
	agentGroupAssignment := CreateAgentGroupAssignment(AgentGroupAssignmentRequestModel{
		GroupUID: plan.GroupID.ValueString(),
		AgentUID: plan.AgentID.ValueString(),
	})

	// Update the state to match the plan
	plan.ID = types.StringValue(agentGroupAssignment.UID)
	plan.GroupID = types.StringValue(agentGroupAssignment.GroupUID)
	plan.AgentID = types.StringValue(agentGroupAssignment.AgentUID)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *agentGroupAssignmentResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	// Get current state
	var state agentGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		AgentGroupAssignmentsGet(ctx).
		Id(state.ID.ValueString())
	agentGroupAssignmentResponse, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_agent_group_assignment")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(agentGroupAssignmentResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)
		return
	}
	agentGroupAssignment := agentGroupAssignmentResponse.Items[0]

	// Update state from client response
	state.GroupID = types.StringValue(agentGroupAssignment.Group.Id)
	state.AgentID = types.StringValue(agentGroupAssignment.Agent.Id)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *agentGroupAssignmentResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	// Retrieve values from plan
	var plan agentGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "updating an agent group assignment is not supported")
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *agentGroupAssignmentResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	// Retrieve values from state
	var state agentGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing agentGroupAssignment using the plan_id
}

func (r *agentGroupAssignmentResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var CreateAgentGroupAssignment = func(request AgentGroupAssignmentRequestModel) AgentGroupAssignmentResponseModel {
	// TODO: Query the agentGroupAssignment using the client

	return AgentGroupAssignmentResponseModel{
		UID:      "mock_uid",
		GroupUID: "mock_group_uid",
		AgentUID: "mock_agent_uid",
	}
}
