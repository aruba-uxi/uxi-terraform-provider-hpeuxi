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
	_ resource.Resource              = &networkGroupAssignmentResource{}
	_ resource.ResourceWithConfigure = &networkGroupAssignmentResource{}
)

type networkGroupAssignmentResourceModel struct {
	ID        types.String `tfsdk:"id"`
	NetworkID types.String `tfsdk:"network_id"`
	GroupID   types.String `tfsdk:"group_id"`
}

func NewNetworkGroupAssignmentResource() resource.Resource {
	return &networkGroupAssignmentResource{}
}

type networkGroupAssignmentResource struct {
	client *config_api_client.APIClient
}

func (r *networkGroupAssignmentResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_network_group_assignment"
}

func (r *networkGroupAssignmentResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Manages a network group assignment." +
			"\n\nNote: it is recommended to use a `hpeuxi_group` **resource** `id` as the `group_id`. " +
			"This will help maintain dependencies between resources. This is useful when a " +
			"destructive action is performed on an ancestor of the assigned group.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the network group assignment",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				Description: "The identifier of the network to be assigned. " +
					"Use `hpeuxi_wired_network` resource id field; " +
					"`hpeuxi_wired_network` datasource id field; " +
					"`hpeuxi_wireless_network` resource id field;" +
					"`hpeuxi_wireless_network` datasource id field " +
					"wired network id; " +
					"or wireless network id here.",
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"group_id": schema.StringAttribute{
				Description: "The identifier of the group to be assigned to. " +
					"Use `hpeuxi_group` resource id field (recommended); " +
					"`hpeuxi_group` datasource id field or " +
					"group id here.",
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *networkGroupAssignmentResource) Configure(
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
			"Resource type: Network Group Assignment. Please report this issue to the provider developers.",
		)

		return
	}

	r.client = client
}

func (r *networkGroupAssignmentResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan networkGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	postRequest := config_api_client.NewNetworkGroupAssignmentPostRequest(
		plan.GroupID.ValueString(),
		plan.NetworkID.ValueString(),
	)
	request := r.client.ConfigurationAPI.
		NetworkGroupAssignmentPost(ctx).
		NetworkGroupAssignmentPostRequest(*postRequest)
	networkGroupAssignment, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	if errorPresent {
		resp.Diagnostics.AddError(
			util.GenerateErrorSummary("create", "hpeuxi_network_group_assignment"),
			errorDetail,
		)

		return
	}

	defer response.Body.Close()

	plan.ID = types.StringValue(networkGroupAssignment.Id)
	plan.GroupID = types.StringValue(networkGroupAssignment.Group.Id)
	plan.NetworkID = types.StringValue(networkGroupAssignment.Network.Id)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *networkGroupAssignmentResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state networkGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		NetworkGroupAssignmentsGet(ctx).
		Id(state.ID.ValueString())
	networkGroupAssignmentResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	errorSummary := util.GenerateErrorSummary("read", "hpeuxi_network_group_assignment")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)

		return
	}

	defer response.Body.Close()

	if len(networkGroupAssignmentResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)

		return
	}
	networkGroupAssignment := networkGroupAssignmentResponse.Items[0]

	state.ID = types.StringValue(networkGroupAssignment.Id)
	state.GroupID = types.StringValue(networkGroupAssignment.Group.Id)
	state.NetworkID = types.StringValue(networkGroupAssignment.Network.Id)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *networkGroupAssignmentResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan networkGroupAssignmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"updating a network group assignment is not supported",
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *networkGroupAssignmentResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state networkGroupAssignmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		NetworkGroupAssignmentDelete(ctx, state.ID.ValueString())
	_, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	if errorPresent {
		if response != nil && response.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)

			return
		}
		resp.Diagnostics.AddError(
			util.GenerateErrorSummary("delete", "hpeuxi_network_group_assignment"),
			errorDetail,
		)

		return
	}

	defer response.Body.Close()
}

func (r *networkGroupAssignmentResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
