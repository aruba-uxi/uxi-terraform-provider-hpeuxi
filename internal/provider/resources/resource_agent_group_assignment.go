/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

import (
	"context"
	"net/http"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/util"
	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &agentGroupAssignmentResource{}
	_ resource.ResourceWithConfigure = &agentGroupAssignmentResource{}
)

type agentGroupAssignmentResourceModel struct {
	ID      types.String `tfsdk:"id"`
	AgentID types.String `tfsdk:"agent_id"`
	GroupID types.String `tfsdk:"group_id"`
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
		Description: "Manages an agent group assignment. " +
			"\n\nNote: it is recommended to use a `hpeuxi_group` **resource** `id` as the `group_id`. " +
			"This will help maintain dependencies between resources. This is useful when a " +
			"destructive action is performed on an ancestor of the assigned group.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the agent group assignment.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"agent_id": schema.StringAttribute{
				Description: "The identifier of the agent to be assigned. " +
					"Use `hpeuxi_agent` resource id field; " +
					"`data.hpeuxi_agent` id field " +
					"or agent id here.",
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"group_id": schema.StringAttribute{
				Description: "The identifier of the group to be assigned to. " +
					"Use `hpeuxi_group` resource id field (recommended); " +
					"`data.hpeuxi_group` id field or " +
					"group id here.",
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
	var plan agentGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postRequest := config_api_client.NewAgentGroupAssignmentPostRequest(
		plan.GroupID.ValueString(),
		plan.AgentID.ValueString(),
	)
	request := r.client.ConfigurationAPI.
		AgentGroupAssignmentPost(ctx).
		AgentGroupAssignmentPostRequest(*postRequest)
	agentGroupAssignment, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	if errorPresent {
		resp.Diagnostics.AddError(
			util.GenerateErrorSummary("create", "hpeuxi_agent_group_assignment"),
			errorDetail,
		)

		return
	}

	defer response.Body.Close()

	plan.ID = types.StringValue(agentGroupAssignment.Id)
	plan.GroupID = types.StringValue(agentGroupAssignment.Group.Id)
	plan.AgentID = types.StringValue(agentGroupAssignment.Agent.Id)

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
	var state agentGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		AgentGroupAssignmentsGet(ctx).
		Id(state.ID.ValueString())
	agentGroupAssignmentResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	errorSummary := util.GenerateErrorSummary("read", "hpeuxi_agent_group_assignment")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)

		return
	}

	defer response.Body.Close()

	if len(agentGroupAssignmentResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)

		return
	}
	agentGroupAssignment := agentGroupAssignmentResponse.Items[0]

	state.ID = types.StringValue(agentGroupAssignment.Id)
	state.GroupID = types.StringValue(agentGroupAssignment.Group.Id)
	state.AgentID = types.StringValue(agentGroupAssignment.Agent.Id)

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
	var state agentGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		AgentGroupAssignmentDelete(ctx, state.ID.ValueString())
	_, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	if errorPresent {
		if response != nil && response.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)

			return
		}
		resp.Diagnostics.AddError(
			util.GenerateErrorSummary("delete", "hpeuxi_agent_group_assignment"),
			errorDetail,
		)

		return
	}

	defer response.Body.Close()
}

func (r *agentGroupAssignmentResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
