package resources

import (
	"context"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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

type GroupResponseModel struct {
	UID       string  `json:"uid"`
	Name      string  `json:"name"`
	ParentUid *string `json:"parent_uid"`
	Path      string  `json:"path"`
}

type GroupCreateRequestModel struct {
	Name      string
	ParentUid string
}

type GroupUpdateRequestModel struct {
	Name string
}

func NewGroupResource() resource.Resource {
	return &groupResource{}
}

type groupResource struct {
	client *config_api_client.APIClient
}

func (r *groupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (r *groupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					// UXI business logic does not permit moving of groups
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *groupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *groupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan groupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	groups_post_request := config_api_client.NewGroupsPostRequest(plan.ParentGroupId.ValueString(), plan.Name.ValueString())
	group, _, err := r.client.ConfigurationAPI.GroupsPostConfigurationAppV1GroupsPost(context.Background()).GroupsPostRequest(*groups_post_request).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating group",
			"Could not create group, unexpected error: "+err.Error(),
		)
		return
	}

	// Update the state to match the plan (replace with response from client)
	plan.ID = types.StringValue(group.Uid)
	plan.Name = types.StringValue(group.Name)
	plan.ParentGroupId = types.StringValue(group.ParentUid)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *groupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state groupResourceModel
	diags := req.State.Get(ctx, &state)
	if state.ID.ValueString() == GetRootGroupUID() {
		diags.AddError("operation not supported", "the root node cannot be used as a resource")
	}
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	groupResponse, _, err := r.client.ConfigurationAPI.
		GroupsGetConfigurationAppV1GroupsGet(context.Background()).
		Uid(state.ID.ValueString()).
		Execute()

	if err != nil || len(groupResponse.Groups) != 1 {
		resp.Diagnostics.AddError(
			"Error reading Group",
			"Could not retrieve Group, unexpected error: "+err.Error(),
		)
		return
	}

	group := groupResponse.Groups[0]

	// Update state from client response
	state.Name = types.StringValue(group.Name)
	state.ParentGroupId = types.StringValue(*group.ParentUid.Get())

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *groupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan groupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client updateGroup method
	group := UpdateGroup(GroupUpdateRequestModel{
		Name: plan.Name.ValueString(),
	})

	// Update the state to match the plan (replace with response from client)
	plan.Name = types.StringValue(group.Name)
	plan.ParentGroupId = types.StringValue(*group.ParentUid)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *groupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state groupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing group using the plan_id
}

func (r *groupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var GetGroup = func(uid string) GroupResponseModel {
	// TODO: Query the group using the client

	parent_uid := "mock_parent_uid"

	return GroupResponseModel{
		UID:       uid,
		Name:      "mock_name",
		ParentUid: &parent_uid,
		Path:      "mock_path",
	}
}

var UpdateGroup = func(request GroupUpdateRequestModel) GroupResponseModel {
	// TODO: Query the group using the client

	parent_uid := "mock_parent_uid"

	return GroupResponseModel{
		UID:       "mock_uid",
		Name:      "mock_name",
		ParentUid: &parent_uid,
		Path:      "mock_path",
	}
}

var GetRootGroupUID = func() string {
	// Get root node here using the client
	return "root_group_uid"
}
