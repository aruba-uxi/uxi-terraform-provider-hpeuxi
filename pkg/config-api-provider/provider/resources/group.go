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
	_ resource.Resource              = &groupResource{}
	_ resource.ResourceWithConfigure = &groupResource{}
)

type groupResourceModel struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	ParentUid types.String `tfsdk:"parent_uid"`
}

type GroupResponseModel struct {
	UID       string
	Name      string
	ParentUid string
	Path      string
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

type groupResource struct{}

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
			"parent_uid": schema.StringAttribute{
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
}

func (r *groupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan groupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client create-group method
	// We are mocking the response of the client for this early stage of development
	group := CreateGroup(GroupCreateRequestModel{
		Name:      plan.Name.ValueString(),
		ParentUid: plan.ParentUid.ValueString(),
	})

	// Update the state to match the plan (replace with response from client)
	plan.ID = types.StringValue(group.UID)
	plan.Name = types.StringValue(group.Name)
	plan.ParentUid = types.StringValue(group.ParentUid)

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
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Call client create-group method
	// We are mocking the response of the client for this early stage of development
	response := GetGroup()

	// Update state from client response
	state.Name = types.StringValue(response.Name)
	state.ParentUid = types.StringValue(response.ParentUid)

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
	plan.ParentUid = types.StringValue(group.ParentUid)

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

var GetGroup = func() GroupResponseModel {
	// TODO: Query the group using the client

	return GroupResponseModel{
		UID:       "mock_uid",
		Name:      "mock_name",
		ParentUid: "mock_parent_uid",
		Path:      "mock_path",
	}
}

var CreateGroup = func(request GroupCreateRequestModel) GroupResponseModel {
	// TODO: Query the group using the client

	return GroupResponseModel{
		UID:       "mock_uid",
		Name:      "mock_name",
		ParentUid: "mock_parent_uid",
		Path:      "mock_path",
	}
}

var UpdateGroup = func(request GroupUpdateRequestModel) GroupResponseModel {
	// TODO: Query the group using the client

	return GroupResponseModel{
		UID:       "mock_uid",
		Name:      "mock_name",
		ParentUid: "mock_parent_uid",
		Path:      "mock_path",
	}
}
