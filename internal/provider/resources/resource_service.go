/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

import (
	"context"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &serviceTestResource{}
	_ resource.ResourceWithConfigure = &serviceTestResource{}
)

type serviceTestResourceModel struct {
	ID        types.String `tfsdk:"id"`
	Category  types.String `tfsdk:"category"`
	Name      types.String `tfsdk:"name"`
	Target    types.String `tfsdk:"target"`
	Template  types.String `tfsdk:"template"`
	IsEnabled types.Bool   `tfsdk:"is_enabled"`
}

func NewServiceTestResource() resource.Resource {
	return &serviceTestResource{}
}

type serviceTestResource struct {
	client *config_api_client.APIClient
}

func (r *serviceTestResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_service_test"
}

func (r *serviceTestResource) Schema(
	_ context.Context,
	_ resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Manages a service test.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the service test.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"category": schema.StringAttribute{
				Description: "The category of the service test.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the service test.",
				Required:    true,
			},
			"target": schema.StringAttribute{
				Description: "The target of the service test.",
				Computed:    true,
			},
			"template": schema.StringAttribute{
				Description: "The template of the service test.",
				Computed:    true,
			},
			"is_enabled": schema.BoolAttribute{
				Description: "Whether the service test is enabled or not.",
				Computed:    true,
			},
		},
	}
}

func (r *serviceTestResource) Configure(
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

func (r *serviceTestResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	var plan serviceTestResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"creating a service_test is not supported; service_tests can only be imported",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *serviceTestResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	var state serviceTestResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		ServiceTestsGet(ctx).
		Id(state.Id.ValueString())
	sensorResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_service_test")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(sensorResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)
		return
	}
	serviceTest := sensorResponse.Items[0]

	state.ID = types.StringValue(serviceTest.Id)
	state.Category = types.StringValue(serviceTest.Category)
	state.Name = types.StringValue(serviceTest.Name)
	state.Target = types.StringPointerValue(serviceTest.Target.Get())
	state.Template = types.StringValue(serviceTest.Template)
	state.IsEnabled = types.BoolValue(serviceTest.IsEnabled)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serviceTestResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	var plan serviceTestResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"updating a service_test is not supported; service_tests can only be updated through the dashboard",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *serviceTestResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	var state serviceTestResourceModel
	diags := req.State.Get(ctx, &state)
	diags.AddError(
		"operation not supported",
		"deleting a service_test is not supported; service_tests can only removed from state",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *serviceTestResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
