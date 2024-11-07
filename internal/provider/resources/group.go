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

const groupNotFoundErrorString = "not found"

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &groupResource{}
	_ resource.ResourceWithConfigure = &groupResource{}
)

type groupResourceModel struct {
	ID            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	ParentGroupId types.String `tfsdk:"parent_group_id"`
}

func NewGroupResource() resource.Resource {
	return &groupResource{}
}

type groupResource struct {
	client *config_api_client.APIClient
}

func (r *groupResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (r *groupResource) Schema(
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
			"parent_group_id": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					// UXI business logic does not permit moving of groups
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *groupResource) Configure(
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

func (r *groupResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	// Retrieve values from plan
	var plan groupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	groups_post_request := config_api_client.NewGroupsPostRequest(plan.Name.ValueString())
	if !plan.ParentGroupId.IsUnknown() && !plan.ParentGroupId.IsNull() {
		groups_post_request.SetParentId(plan.ParentGroupId.ValueString())
	}
	request := r.client.ConfigurationAPI.
		GroupsPost(ctx).
		GroupsPostRequest(*groups_post_request)
	group, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(util.GenerateErrorSummary("create", "uxi_group"), errorDetail)
		return
	}

	// Update the state to match the plan (replace with response from client)
	plan.ID = types.StringValue(group.Id)
	plan.Name = types.StringValue(group.Name)
	// only update parent if not attached to root node (else leave it as null)
	parentGroup, _ := r.getGroup(ctx, group.Parent.Id)
	if parentGroup != nil && !util.IsRoot(*parentGroup) {
		plan.ParentGroupId = types.StringValue(group.Parent.Id)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *groupResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	// Get current state
	var state groupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	group, errorDetail := r.getGroup(ctx, state.ID.ValueString())

	errorSummary := util.GenerateErrorSummary("read", "uxi_group")

	if errorDetail != nil {
		if *errorDetail == groupNotFoundErrorString {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(errorSummary, *errorDetail)
		return
	}

	if util.IsRoot(*group) {
		resp.Diagnostics.AddError(errorSummary, "The root group cannot be used as a resource")
		return
	}

	// Update state from client response
	state.ID = types.StringValue(group.Id)
	state.Name = types.StringValue(group.Name)
	// only update parent if not attached to root node (else leave it as null)
	parentGroup, _ := r.getGroup(ctx, group.Parent.Get().Id)
	if parentGroup != nil && !util.IsRoot(*parentGroup) {
		state.ParentGroupId = types.StringValue(group.Parent.Get().Id)
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *groupResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	// Retrieve values from plan
	var plan, state groupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	patchRequest := config_api_client.NewGroupsPatchRequest(plan.Name.ValueString())
	request := r.client.ConfigurationAPI.
		GroupsPatch(ctx, plan.ID.ValueString()).
		GroupsPatchRequest(*patchRequest)
	group, response, err := util.RetryFor429(request.Execute)

	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(util.GenerateErrorSummary("update", "uxi_group"), errorDetail)
		return
	}

	// Update the state to match the plan (replace with response from client)
	plan.ID = types.StringValue(group.Id)
	plan.Name = types.StringValue(group.Name)
	// only update parent if not attached to root node (else leave it as null)
	parentGroup, _ := r.getGroup(ctx, group.Parent.Id)
	if parentGroup != nil && !util.IsRoot(*parentGroup) {
		state.ParentGroupId = types.StringValue(group.Parent.Id)
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *groupResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	// Retrieve values from state
	var state groupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing group using the plan_id
	request := r.client.ConfigurationAPI.GroupsDelete(ctx, state.ID.ValueString())

	_, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(util.GenerateErrorSummary("delete", "uxi_group"), errorDetail)
		return
	}
}

func (r *groupResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *groupResource) getGroup(
	ctx context.Context,
	id string,
) (*config_api_client.GroupsGetItem, *string) {
	request := r.client.ConfigurationAPI.GroupsGet(ctx).Id(id)
	groupResponse, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		return nil, &errorDetail
	}

	if len(groupResponse.Items) != 1 {
		notFound := groupNotFoundErrorString
		return nil, &notFound
	}

	return &groupResponse.Items[0], nil
}
