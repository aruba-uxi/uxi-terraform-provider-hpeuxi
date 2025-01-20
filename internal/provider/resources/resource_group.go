/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

import (
	"context"
	"errors"
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

const groupNotFoundError = "not found"

var (
	_ resource.Resource              = &groupResource{}
	_ resource.ResourceWithConfigure = &groupResource{}
)

type groupResourceModel struct {
	ID            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	ParentGroupID types.String `tfsdk:"parent_group_id"`
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
		Description: "Manages a group. " +
			"\n\nNote: building a group hierarchy by using an `hpeuxi_group` **resource** `id` as " +
			"a child group's `parent_group_id` is recommended to maintain dependencies between " +
			"linked groups. This will help maintain accurate state if the user attempts to " +
			"change the parent of a non leaf group.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the group.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the group.",
				Required:    true,
			},
			"parent_group_id": schema.StringAttribute{
				Description: "The identifier of the parent of this group. " +
					"Use `hpeuxi_group` resource (recommended) or `data.hpeuxi_group` id for this attribute. " +
					"Alternatively leave blank to set group to highest level configurable node.",
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
	var plan groupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	groupsPostRequest := config_api_client.NewGroupPostRequest(plan.Name.ValueString())
	if !plan.ParentGroupID.IsUnknown() && !plan.ParentGroupID.IsNull() {
		groupsPostRequest.SetParentId(plan.ParentGroupID.ValueString())
	}
	request := r.client.ConfigurationAPI.
		GroupPost(ctx).
		GroupPostRequest(*groupsPostRequest)
	group, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	if errorPresent {
		resp.Diagnostics.AddError(util.GenerateErrorSummary("create", "hpeuxi_group"), errorDetail)

		return
	}

	defer response.Body.Close()

	plan.ID = types.StringValue(group.Id)
	plan.Name = types.StringValue(group.Name)

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
	var state groupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	group, errorDetail := r.getGroup(ctx, state.ID.ValueString())

	errorSummary := util.GenerateErrorSummary("read", "hpeuxi_group")

	if errorDetail != nil {
		if errorDetail.Error() == groupNotFoundError {
			resp.State.RemoveResource(ctx)

			return
		}
		resp.Diagnostics.AddError(errorSummary, errorDetail.Error())

		return
	}

	if util.IsRoot(*group) {
		resp.Diagnostics.AddError(errorSummary, "The root group cannot be used as a resource")

		return
	}

	state.ID = types.StringValue(group.Id)
	state.Name = types.StringValue(group.Name)
	parentGroup, _ := r.getGroup(ctx, group.Parent.Get().Id)
	if parentGroup != nil && !util.IsRoot(*parentGroup) {
		state.ParentGroupID = types.StringValue(group.Parent.Get().Id)
	}

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
	var plan groupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	patchRequest := config_api_client.NewGroupPatchRequest()
	patchRequest.Name = plan.Name.ValueStringPointer()
	request := r.client.ConfigurationAPI.
		GroupPatch(ctx, plan.ID.ValueString()).
		GroupPatchRequest(*patchRequest)
	group, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	if errorPresent {
		resp.Diagnostics.AddError(util.GenerateErrorSummary("update", "hpeuxi_group"), errorDetail)

		return
	}

	defer response.Body.Close()

	plan.ID = types.StringValue(group.Id)
	plan.Name = types.StringValue(group.Name)

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
	var state groupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.GroupDelete(ctx, state.ID.ValueString())

	response, err := util.RetryForTooManyRequestsNoReturn(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	if errorPresent {
		if response != nil && response.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)

			return
		}
		resp.Diagnostics.AddError(util.GenerateErrorSummary("delete", "hpeuxi_group"), errorDetail)

		return
	}

	defer response.Body.Close()
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
) (*config_api_client.GroupsGetItem, error) {
	request := r.client.ConfigurationAPI.GroupsGet(ctx).Id(id)

	groupResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	if errorPresent {
		return nil, errors.New(errorDetail)
	}
	if len(groupResponse.Items) != 1 {
		notFound := groupNotFoundError

		return nil, errors.New(notFound)
	}

	defer response.Body.Close()

	return &groupResponse.Items[0], nil
}
