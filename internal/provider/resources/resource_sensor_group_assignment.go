/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

import (
	"context"
	"net/http"

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
	_ resource.Resource              = &sensorGroupAssignmentResource{}
	_ resource.ResourceWithConfigure = &sensorGroupAssignmentResource{}
)

type sensorGroupAssignmentResourceModel struct {
	ID       types.String `tfsdk:"id"`
	SensorID types.String `tfsdk:"sensor_id"`
	GroupID  types.String `tfsdk:"group_id"`
}

func NewSensorGroupAssignmentResource() resource.Resource {
	return &sensorGroupAssignmentResource{}
}

type sensorGroupAssignmentResource struct {
	client *config_api_client.APIClient
}

func (r *sensorGroupAssignmentResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_sensor_group_assignment"
}

func (r *sensorGroupAssignmentResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Manages a sensor group assignment.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the sensor group assignment",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"sensor_id": schema.StringAttribute{
				Description: "The identifier of the sensor to be assigned. " +
					"Use sensor id; " +
					"`uxi_sensor` resource id field or " +
					"`uxi_sensor` datasource id field here.",
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"group_id": schema.StringAttribute{
				Description: "The identifier of the group to be assigned to. " +
					"Use group id; " +
					"`uxi_group` resource id field or " +
					"`uxi_group` datasource id field here.",
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *sensorGroupAssignmentResource) Configure(
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
			"Resource type: Sensor Group Assignment. Please report this issue to the provider developers.",
		)

		return
	}

	r.client = client
}

func (r *sensorGroupAssignmentResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan sensorGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postRequest := config_api_client.NewSensorGroupAssignmentsPostRequest(
		plan.GroupID.ValueString(),
		plan.SensorID.ValueString(),
	)
	request := r.client.ConfigurationAPI.
		SensorGroupAssignmentsPost(ctx).
		SensorGroupAssignmentsPostRequest(*postRequest)
	sensorGroupAssignment, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	if errorPresent {
		resp.Diagnostics.AddError(
			util.GenerateErrorSummary("create", "uxi_sensor_group_assignment"),
			errorDetail,
		)

		return
	}

	defer response.Body.Close()

	plan.ID = types.StringValue(sensorGroupAssignment.Id)
	plan.GroupID = types.StringValue(sensorGroupAssignment.Group.Id)
	plan.SensorID = types.StringValue(sensorGroupAssignment.Sensor.Id)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorGroupAssignmentResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state sensorGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		SensorGroupAssignmentsGet(ctx).
		Id(state.ID.ValueString())
	sensorGroupAssignmentResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	errorSummary := util.GenerateErrorSummary("read", "uxi_sensor_group_assignment")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)

		return
	}

	defer response.Body.Close()

	if len(sensorGroupAssignmentResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)

		return
	}

	sensorGroupAssignment := sensorGroupAssignmentResponse.Items[0]

	state.ID = types.StringValue(sensorGroupAssignment.Id)
	state.GroupID = types.StringValue(sensorGroupAssignment.Group.Id)
	state.SensorID = types.StringValue(sensorGroupAssignment.Sensor.Id)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorGroupAssignmentResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan sensorGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "updating a sensor group assignment is not supported")
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorGroupAssignmentResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state sensorGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		SensorGroupAssignmentsDelete(ctx, state.ID.ValueString())
	_, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	if errorPresent {
		if response != nil && response.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)

			return
		}
		resp.Diagnostics.AddError(
			util.GenerateErrorSummary("delete", "uxi_sensor_group_assignment"),
			errorDetail,
		)

		return
	}

	defer response.Body.Close()
}

func (r *sensorGroupAssignmentResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
