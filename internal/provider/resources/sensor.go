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
				Computed: true,
			},
			"notes": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"pcap_mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}

func (r *sensorResource) Configure(
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

func (r *sensorResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
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

	state.ID = types.StringValue(sensor.Id)
	state.Name = types.StringValue(sensor.Name)
	state.AddressNote = types.StringPointerValue(sensor.AddressNote.Get())
	state.Notes = types.StringPointerValue(sensor.Notes.Get())
	state.PCapMode = types.StringPointerValue(sensor.PcapMode.Get())

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
	var plan sensorResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	patchRequest := config_api_client.NewSensorsPatchRequest()
	patchRequest.Name = plan.Name.ValueStringPointer()
	patchRequest.AddressNote = plan.AddressNote.ValueStringPointer()
	patchRequest.Notes = plan.Notes.ValueStringPointer()
	patchRequest.PcapMode = plan.PCapMode.ValueStringPointer()

	request := r.client.ConfigurationAPI.
		SensorsPatch(ctx, plan.ID.ValueString()).
		SensorsPatchRequest(*patchRequest)
	sensor, response, err := util.RetryFor429(request.Execute)

	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(util.GenerateErrorSummary("update", "uxi_sensor"), errorDetail)
		return
	}

	plan.ID = types.StringValue(sensor.Id)
	plan.Name = types.StringValue(sensor.Name)
	plan.AddressNote = types.StringPointerValue(sensor.AddressNote.Get())
	plan.Notes = types.StringPointerValue(sensor.Notes.Get())
	plan.PCapMode = types.StringPointerValue(sensor.PcapMode.Get())

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
