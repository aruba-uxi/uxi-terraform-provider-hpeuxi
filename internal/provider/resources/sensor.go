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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &sensorResource{}
	_ resource.ResourceWithConfigure = &sensorResource{}
)

type sensorResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	AddressNote types.String `tfsdk:"address_note"`
	Notes       types.String `tfsdk:"notes"`
	PCapMode    types.String `tfsdk:"pcap_mode"`
}

func NewSensorResource() resource.Resource {
	return &sensorResource{}
}

type sensorResource struct {
	client *config_api_client.APIClient
}

func (r *sensorResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_sensor"
}

func (r *sensorResource) Schema(
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
			"address_note": schema.StringAttribute{
				Optional: true,
			},
			"notes": schema.StringAttribute{
				Optional: true,
			},
			"pcap_mode": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *sensorResource) Configure(
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

func (r *sensorResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	// Retrieve values from plan
	var plan sensorResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError(
		"operation not supported",
		"creating a sensor is not supported; sensors can only be imported",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *sensorResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	// Get current state
	var state sensorResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		SensorsGet(ctx).
		Id(state.ID.ValueString())
	sensorResponse, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_sensor")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(sensorResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)
		return
	}
	sensor := sensorResponse.Items[0]

	// Update state from client response
	state.ID = types.StringValue(sensor.Id)
	state.Name = types.StringValue(sensor.Name)
	state.AddressNote = types.StringPointerValue(sensor.AddressNote.Get())
	state.Notes = types.StringPointerValue(sensor.Notes.Get())
	state.PCapMode = types.StringPointerValue(sensor.PcapMode.Get())

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	// Retrieve values from plan
	var plan sensorResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	patchRequest := config_api_client.NewSensorsPatchRequest()
	if !plan.AddressNote.IsUnknown() {
		patchRequest.AddressNote = *config_api_client.NewNullableString(plan.AddressNote.ValueStringPointer())
	}
	if !plan.Notes.IsUnknown() {
		patchRequest.Notes = *config_api_client.NewNullableString(plan.Notes.ValueStringPointer())
	}
	if !plan.PCapMode.IsUnknown() {
		patchRequest.PcapMode = plan.PCapMode.ValueStringPointer()
	}
	request := r.client.ConfigurationAPI.
		SensorsPatch(ctx, plan.ID.ValueString()).
		SensorsPatchRequest(*patchRequest)
	sensor, response, err := util.RetryFor429(request.Execute)

	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(util.GenerateErrorSummary("update", "uxi_sensor"), errorDetail)
		return
	}

	// Update the state to match the plan (replace with response from client)
	plan.ID = types.StringValue(sensor.Id)
	plan.Name = types.StringValue(sensor.Name)
	if sensor.AddressNote.Get() != nil {
		plan.AddressNote = types.StringValue(*sensor.AddressNote.Get())
	}
	if sensor.Notes.Get() != nil {
		plan.Notes = types.StringValue(*sensor.Notes.Get())
	}
	if sensor.PcapMode.Get() != nil {
		plan.PCapMode = types.StringValue(*sensor.PcapMode.Get())
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	// Retrieve values from state
	var state sensorResourceModel
	diags := req.State.Get(ctx, &state)
	diags.AddError(
		"operation not supported",
		"deleting a sensor is not supported; sensors can only removed from state",
	)
	resp.Diagnostics.Append(diags...)
}

func (r *sensorResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
