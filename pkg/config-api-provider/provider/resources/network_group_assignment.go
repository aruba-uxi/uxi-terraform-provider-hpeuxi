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
	_ resource.Resource              = &networkGroupAssignmentResource{}
	_ resource.ResourceWithConfigure = &networkGroupAssignmentResource{}
)

type networkGroupAssignmentResourceModel struct {
	ID        types.String `tfsdk:"id"`
	NetworkID types.String `tfsdk:"network_id"`
	GroupID   types.String `tfsdk:"group_id"`
}

type NetworkGroupAssignmentResponseModel struct {
	UID        string //  <assignment_uid>
	GroupUID   string //  <group_uid:str>,
	NetworkUID string //  <network_uid:str>
}

type NetworkGroupAssignmentRequestModel struct {
	GroupUID   string //  <group_uid:str>,
	NetworkUID string //  <network_uid:str>
}

func NewNetworkGroupAssignmentResource() resource.Resource {
	return &networkGroupAssignmentResource{}
}

type networkGroupAssignmentResource struct {
	client *config_api_client.APIClient
}

func (r *networkGroupAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_group_assignment"
}

func (r *networkGroupAssignmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
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

func (r *networkGroupAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *networkGroupAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan networkGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client createNetworkGroupAssignment method
	networkGroupAssignment := CreateNetworkGroupAssignment(NetworkGroupAssignmentRequestModel{
		GroupUID:   plan.GroupID.ValueString(),
		NetworkUID: plan.NetworkID.ValueString(),
	})

	// Update the state to match the plan
	plan.ID = types.StringValue(networkGroupAssignment.UID)
	plan.GroupID = types.StringValue(networkGroupAssignment.GroupUID)
	plan.NetworkID = types.StringValue(networkGroupAssignment.NetworkUID)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *networkGroupAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state networkGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client getNetworkGroupAssignment method
	request := r.client.ConfigurationAPI.
		GetUxiV1alpha1NetworkGroupAssignmentsGet(ctx).
		Uid(state.ID.ValueString())

	networkGroupAssignmentResponse, _, err := util.RetryFor429(request.Execute)
	if err != nil || len(networkGroupAssignmentResponse.Items) != 1 {
		resp.Diagnostics.AddError(
			"Error reading Network Group Assignment",
			"Could not retrieve Network Group Assignment, unexpected error: "+err.Error(),
		)
		return
	}
	networkGroupAssignment := networkGroupAssignmentResponse.Items[0]

	// Update state from client response
	state.GroupID = types.StringValue(networkGroupAssignment.Group.Id)
	state.NetworkID = types.StringValue(networkGroupAssignment.Network.Id)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *networkGroupAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan networkGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "updating a network group assignment is not supported")
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *networkGroupAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state networkGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing networkGroupAssignment using the plan_id
}

func (r *networkGroupAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var CreateNetworkGroupAssignment = func(request NetworkGroupAssignmentRequestModel) NetworkGroupAssignmentResponseModel {
	// TODO: Query the networkGroupAssignment using the client

	return NetworkGroupAssignmentResponseModel{
		UID:        "mock_uid",
		GroupUID:   "mock_group_uid",
		NetworkUID: "mock_network_uid",
	}
}
