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
	_ resource.Resource              = &sensorGroupAssignmentResource{}
	_ resource.ResourceWithConfigure = &sensorGroupAssignmentResource{}
)

type sensorGroupAssignmentResourceModel struct {
	ID       types.String `tfsdk:"id"`
	SensorID types.String `tfsdk:"sensor_id"`
	GroupID  types.String `tfsdk:"group_id"`
}

type SensorGroupAssignmentResponseModel struct {
	UID       string //  <assignment_uid>
	GroupUID  string //  <group_uid:str>,
	SensorUID string //  <sensor_uid:str>
}

type SensorGroupAssignmentRequestModel struct {
	GroupUID  string //  <group_uid:str>,
	SensorUID string //  <sensor_uid:str>
}

func NewSensorGroupAssignmentResource() resource.Resource {
	return &sensorGroupAssignmentResource{}
}

type sensorGroupAssignmentResource struct{}

func (r *sensorGroupAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sensor_group_assignment"
}

func (r *sensorGroupAssignmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"sensor_id": schema.StringAttribute{
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

func (r *sensorGroupAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
}

func (r *sensorGroupAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan sensorGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client createSensorGroupAssignment method
	sensorGroupAssignment := CreateSensorGroupAssignment(SensorGroupAssignmentRequestModel{
		GroupUID:  plan.GroupID.ValueString(),
		SensorUID: plan.SensorID.ValueString(),
	})

	// Update the state to match the plan
	plan.ID = types.StringValue(sensorGroupAssignment.UID)
	plan.GroupID = types.StringValue(sensorGroupAssignment.GroupUID)
	plan.SensorID = types.StringValue(sensorGroupAssignment.SensorUID)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorGroupAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state sensorGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client getSensorGroupAssignment method
	sensorGroupAssignment := GetSensorGroupAssignment(state.ID.ValueString())

	// Update state from client response
	state.GroupID = types.StringValue(sensorGroupAssignment.GroupUID)
	state.SensorID = types.StringValue(sensorGroupAssignment.SensorUID)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorGroupAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan sensorGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "updating a sensor group assignment is not supported")
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorGroupAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state sensorGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing sensorGroupAssignment using the plan_id
}

func (r *sensorGroupAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var GetSensorGroupAssignment = func(uid string) SensorGroupAssignmentResponseModel {
	// TODO: Query the sensorGroupAssignment using the client

	return SensorGroupAssignmentResponseModel{
		UID:       uid,
		GroupUID:  "mock_group_uid",
		SensorUID: "mock_sensor_uid",
	}
}

var CreateSensorGroupAssignment = func(request SensorGroupAssignmentRequestModel) SensorGroupAssignmentResponseModel {
	// TODO: Query the sensorGroupAssignment using the client

	return SensorGroupAssignmentResponseModel{
		UID:       "mock_uid",
		GroupUID:  "mock_group_uid",
		SensorUID: "mock_sensor_uid",
	}
}
