/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

import (
	"context"

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
	_ resource.Resource              = &sensorResource{}
	_ resource.ResourceWithConfigure = &sensorResource{}
)

type sensorResourceModel struct {
	ID                 types.String  `tfsdk:"id"`
	Serial             types.String  `tfsdk:"serial"`
	Name               types.String  `tfsdk:"name"`
	ModelNumber        types.String  `tfsdk:"model_number"`
	WifiMacAddress     types.String  `tfsdk:"wifi_mac_address"`
	EthernetMacAddress types.String  `tfsdk:"ethernet_mac_address"`
	AddressNote        types.String  `tfsdk:"address_note"`
	Longitude          types.Float32 `tfsdk:"longitude"`
	Latitude           types.Float32 `tfsdk:"latitude"`
	Notes              types.String  `tfsdk:"notes"`
	PcapMode           types.String  `tfsdk:"pcap_mode"`
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
		Description: "Manages a sensor.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the sensor.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial": schema.StringAttribute{
				Description: "The serial number of the sensor.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the sensor.",
				Required:    true,
			},
			"model_number": schema.StringAttribute{
				Description: "The model number of the sensor.",
				Computed:    true,
			},
			"wifi_mac_address": schema.StringAttribute{
				Description: "The wifi mac address of the sensor.",
				Computed:    true,
			},
			"ethernet_mac_address": schema.StringAttribute{
				Description: "The ethernet mac address of the sensor.",
				Computed:    true,
			},
			"address_note": schema.StringAttribute{
				Description: "The address notes of the sensor.",
				Optional:    true,
				// computed because goes from nil -> "" when sensor becomes configured
				Computed: true,
			},
			"longitude": schema.Float32Attribute{
				Description: "The geolocation longitude of the sensor.",
				Computed:    true,
			},
			"latitude": schema.Float32Attribute{
				Description: "The geolocation latitude of the sensor.",
				Computed:    true,
			},
			"notes": schema.StringAttribute{
				Description: "The address notes of the sensor.",
				Optional:    true,
				// computed because goes from nil -> "" when sensor becomes configured
				Computed: true,
			},
			"pcap_mode": schema.StringAttribute{
				Description: "The packet capture mode of the agent.",
				Optional:    true,
				// computed because goes from nil -> "light" when sensor becomes configured
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
	sensorResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)
	errorSummary := util.GenerateErrorSummary("read", "hpeuxi_sensor")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)

		return
	}

	defer response.Body.Close()

	if len(sensorResponse.Items) != 1 {
		resp.State.RemoveResource(ctx)

		return
	}
	sensor := sensorResponse.Items[0]

	state.ID = types.StringValue(sensor.Id)
	state.Name = types.StringValue(sensor.Name)
	state.Serial = types.StringValue(sensor.Serial)
	state.ModelNumber = types.StringValue(sensor.ModelNumber)
	state.WifiMacAddress = types.StringPointerValue(sensor.WifiMacAddress.Get())
	state.EthernetMacAddress = types.StringPointerValue(sensor.EthernetMacAddress.Get())
	state.AddressNote = types.StringPointerValue(sensor.AddressNote.Get())
	state.Latitude = types.Float32PointerValue(sensor.Latitude.Get())
	state.Longitude = types.Float32PointerValue(sensor.Longitude.Get())
	state.Notes = types.StringPointerValue(sensor.Notes.Get())
	state.PcapMode = types.StringPointerValue((*string)(sensor.GetPcapMode().Ptr()))

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

	errorSummary := util.GenerateErrorSummary("update", "hpeuxi_sensor")
	patchRequest := config_api_client.NewSensorPatchRequest()
	patchRequest.Name = plan.Name.ValueStringPointer()
	patchRequest.AddressNote = plan.AddressNote.ValueStringPointer()
	patchRequest.Notes = plan.Notes.ValueStringPointer()
	plannedPcapMode := plan.PcapMode.ValueStringPointer()
	if !plan.PcapMode.IsUnknown() && plannedPcapMode != nil {
		pcapMode, err := config_api_client.NewSensorPcapModeFromValue(*plannedPcapMode)
		if err != nil {
			resp.Diagnostics.AddError(errorSummary, err.Error())

			return
		}
		patchRequest.PcapMode = pcapMode
	}

	request := r.client.ConfigurationAPI.
		SensorPatch(ctx, plan.ID.ValueString()).
		SensorPatchRequest(*patchRequest)
	sensor, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)

		return
	}

	defer response.Body.Close()

	plan.ID = types.StringValue(sensor.Id)
	plan.Name = types.StringValue(sensor.Name)
	plan.Serial = types.StringValue(sensor.Serial)
	plan.ModelNumber = types.StringValue(sensor.ModelNumber)
	plan.WifiMacAddress = types.StringPointerValue(sensor.WifiMacAddress.Get())
	plan.EthernetMacAddress = types.StringPointerValue(sensor.EthernetMacAddress.Get())
	plan.AddressNote = types.StringPointerValue(sensor.AddressNote.Get())
	plan.Latitude = types.Float32PointerValue(sensor.Latitude.Get())
	plan.Longitude = types.Float32PointerValue(sensor.Longitude.Get())
	plan.Notes = types.StringPointerValue(sensor.Notes.Get())
	plan.PcapMode = types.StringPointerValue((*string)(sensor.GetPcapMode().Ptr()))

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
