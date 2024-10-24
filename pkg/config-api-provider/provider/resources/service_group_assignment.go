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

func (r *serviceTestGroupAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
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
		PostUxiV1alpha1ServiceTestGroupAssignmentsPost(ctx).
		ServiceTestGroupAssignmentsPostRequest(*postRequest)
	serviceTestGroupAssignment, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(util.GenerateErrorSummary("delete", "uxi_service_test_group_assignment"), errorDetail)
		return
	}

	// Update the state to match the plan
	plan.ID = types.StringValue(serviceTestGroupAssignment.Id)
	plan.GroupID = types.StringValue(serviceTestGroupAssignment.Group.Id)
	plan.ServiceTestID = types.StringValue(serviceTestGroupAssignment.ServiceTest.Id)

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
	state.GroupID = types.StringValue(serviceTestGroupAssignment.Group.Id)
	state.ServiceTestID = types.StringValue(serviceTestGroupAssignment.ServiceTest.Id)

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

	r.client.ConfigurationAPI.
		DeleteServiceTestGroupAssignmentUxiV1alpha1ServiceTestGroupAssignmentsIdDelete(ctx, state.ID.ValueString()).
		Execute()
}

func (r *serviceTestGroupAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var GetServiceTestGroupAssignment = func(uid string) config_api_client.ServiceTestGroupAssignmentsPostResponse {
	// TODO: Query the serviceTestGroupAssignment using the client
	resourceType := "uxi/service-test-group-assignment"
	return config_api_client.ServiceTestGroupAssignmentsPostResponse{
		Id:          uid,
		Group:       *config_api_client.NewGroup("mock_group_uid"),
		ServiceTest: *config_api_client.NewServiceTest("mock_serviceTest_uid"),
		Type:        &resourceType,
	}
}
