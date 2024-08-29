package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &serviceTestGroupAssignmentResource{}
	_ resource.ResourceWithConfigure = &serviceTestGroupAssignmentResource{}
)

type serviceTestGroupAssignmentResourceModel struct {
	ID            types.String `tfsdk:"id"`
	ServiceTestID types.String `tfsdk:"service_test_id"`
	GroupID       types.String `tfsdk:"group_id"`
}

type ServiceTestGroupAssignmentResponseModel struct {
	UID            string //  <assignment_uid>
	GroupUID       string //  <group_uid:str>,
	ServiceTestUID string //  <service_test_uid:str>
}

type ServiceTestGroupAssignmentRequestModel struct {
	GroupUID       string //  <group_uid:str>,
	ServiceTestUID string //  <service_test_uid:str>
}

func NewServiceTestGroupAssignmentResource() resource.Resource {
	return &serviceTestGroupAssignmentResource{}
}

type serviceTestGroupAssignmentResource struct{}

func (r *serviceTestGroupAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_test_group_assignment"
}

func (r *serviceTestGroupAssignmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"service_test_id": schema.StringAttribute{
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

func (r *serviceTestGroupAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
}

func (r *serviceTestGroupAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan serviceTestGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client createServiceTestGroupAssignment method
	serviceTestGroupAssignment := CreateServiceTestGroupAssignment(ServiceTestGroupAssignmentRequestModel{
		GroupUID:       plan.GroupID.ValueString(),
		ServiceTestUID: plan.ServiceTestID.ValueString(),
	})

	// Update the state to match the plan
	plan.ID = types.StringValue(serviceTestGroupAssignment.UID)
	plan.GroupID = types.StringValue(serviceTestGroupAssignment.GroupUID)
	plan.ServiceTestID = types.StringValue(serviceTestGroupAssignment.ServiceTestUID)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serviceTestGroupAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state serviceTestGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client getServiceTestGroupAssignment method
	serviceTestGroupAssignment := GetServiceTestGroupAssignment(state.ID.ValueString())

	// Update state from client response
	state.GroupID = types.StringValue(serviceTestGroupAssignment.GroupUID)
	state.ServiceTestID = types.StringValue(serviceTestGroupAssignment.ServiceTestUID)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serviceTestGroupAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan serviceTestGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "updating a service_test group assignment is not supported")
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serviceTestGroupAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state serviceTestGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing serviceTestGroupAssignment using the plan_id
}

func (r *serviceTestGroupAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var GetServiceTestGroupAssignment = func(uid string) ServiceTestGroupAssignmentResponseModel {
	// TODO: Query the serviceTestGroupAssignment using the client

	return ServiceTestGroupAssignmentResponseModel{
		UID:            uid,
		GroupUID:       "mock_group_uid",
		ServiceTestUID: "mock_serviceTest_uid",
	}
}

var CreateServiceTestGroupAssignment = func(request ServiceTestGroupAssignmentRequestModel) ServiceTestGroupAssignmentResponseModel {
	// TODO: Query the serviceTestGroupAssignment using the client

	return ServiceTestGroupAssignmentResponseModel{
		UID:            "mock_uid",
		GroupUID:       "mock_group_uid",
		ServiceTestUID: "mock_serviceTest_uid",
	}
}
