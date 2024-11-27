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
	_ resource.Resource              = &serviceTestGroupAssignmentResource{}
	_ resource.ResourceWithConfigure = &serviceTestGroupAssignmentResource{}
)

type serviceTestGroupAssignmentResourceModel struct {
	ID            types.String `tfsdk:"id"`
	ServiceTestID types.String `tfsdk:"service_test_id"`
	GroupID       types.String `tfsdk:"group_id"`
}

func NewServiceTestGroupAssignmentResource() resource.Resource {
	return &serviceTestGroupAssignmentResource{}
}

type serviceTestGroupAssignmentResource struct {
	client *config_api_client.APIClient
}

func (r *serviceTestGroupAssignmentResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_service_test_group_assignment"
}

func (r *serviceTestGroupAssignmentResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Manages a service test group assignment.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the service test group assignment",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"service_test_id": schema.StringAttribute{
				Description: "The identifier of the service test to be assigned. " +
					"Use service test id; " +
					"`uxi_service_test` resource id field or " +
					"`uxi_service_test` datasource id field here.",
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

func (r *serviceTestGroupAssignmentResource) Configure(
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
			"Resource type: Service Test Group Assignment. Please report this issue to the provider developers.",
		)
		return
	}

	r.client = client
}

func (r *serviceTestGroupAssignmentResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan serviceTestGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postRequest := config_api_client.NewServiceTestGroupAssignmentsPostRequest(
		plan.GroupID.ValueString(),
		plan.ServiceTestID.ValueString(),
	)
	request := r.client.ConfigurationAPI.
		ServiceTestGroupAssignmentsPost(ctx).
		ServiceTestGroupAssignmentsPostRequest(*postRequest)
	serviceTestGroupAssignment, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(
			util.GenerateErrorSummary("create", "uxi_service_test_group_assignment"),
			errorDetail,
		)
		return
	}

	plan.ID = types.StringValue(serviceTestGroupAssignment.Id)
	plan.GroupID = types.StringValue(serviceTestGroupAssignment.Group.Id)
	plan.ServiceTestID = types.StringValue(serviceTestGroupAssignment.ServiceTest.Id)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serviceTestGroupAssignmentResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state serviceTestGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		ServiceTestGroupAssignmentsGet(ctx).
		Id(state.ID.ValueString())
	serviceTestGroupAssignmentResponse, response, err := util.RetryForTooManyRequests(
		request.Execute,
	)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_service_test_group_assignment")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(serviceTestGroupAssignmentResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)
		return
	}
	serviceTestGroupAssignment := serviceTestGroupAssignmentResponse.Items[0]

	state.ID = types.StringValue(serviceTestGroupAssignment.Id)
	state.GroupID = types.StringValue(serviceTestGroupAssignment.Group.Id)
	state.ServiceTestID = types.StringValue(serviceTestGroupAssignment.ServiceTest.Id)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serviceTestGroupAssignmentResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan serviceTestGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"updating a service_test group assignment is not supported",
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serviceTestGroupAssignmentResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state serviceTestGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		ServiceTestGroupAssignmentsDelete(ctx, state.ID.ValueString())

	_, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		if response != nil && response.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			util.GenerateErrorSummary("delete", "uxi_service_test_group_assignment"),
			errorDetail,
		)
		return
	}
}

func (r *serviceTestGroupAssignmentResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
