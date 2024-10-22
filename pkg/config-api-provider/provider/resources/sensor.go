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

// TODO: Switch this to use the Client Model when that becomes available
type SensorResponseModel struct {
	UID                string
	Serial             string
	Name               string
	ModelNumber        string
	WifiMacAddress     string
	EthernetMacAddress string
	AddressNote        string
	Longitude          string
	Latitude           string
	Notes              string
	PCapMode           string
}

// TODO: Switch this to use the Client Model when that becomes available
type SensorUpdateRequestModel struct {
	Name        string
	AddressNote string
	Notes       string
	PCapMode    string
}

func NewSensorResource() resource.Resource {
	return &sensorResource{}
}

type sensorResource struct {
	client *config_api_client.APIClient
}

func (r *sensorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sensor"
}

func (r *sensorResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *sensorResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *sensorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan sensorResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "creating a sensor is not supported; sensors can only be imported")
	resp.Diagnostics.Append(diags...)
}

func (r *sensorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state sensorResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := r.client.ConfigurationAPI.
		GetUxiV1alpha1SensorsGet(ctx).
		Id(state.ID.ValueString())
	sensorResponse, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_sensor")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(sensorResponse.Items) != 1 {
		resp.Diagnostics.AddError(errorSummary, "Could not find specified resource")
		return
	}
	sensor := sensorResponse.Items[0]

	// Update state from client response
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

func (r *sensorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan sensorResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing item
	response := UpdateSensor(SensorUpdateRequestModel{
		Name:        plan.Name.ValueString(),
		AddressNote: plan.AddressNote.ValueString(),
		Notes:       plan.Notes.ValueString(),
		PCapMode:    plan.PCapMode.ValueString(),
	})

	// Update resource state with updated items
	plan.Name = types.StringValue(response.Name)
	plan.AddressNote = types.StringValue(response.AddressNote)
	plan.Notes = types.StringValue(response.Notes)
	plan.PCapMode = types.StringValue(response.PCapMode)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *sensorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state sensorResourceModel
	diags := req.State.Get(ctx, &state)
	diags.AddError("operation not supported", "deleting a sensor is not supported; sensors can only removed from state")
	resp.Diagnostics.Append(diags...)
}

func (r *sensorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Update the sensor using the configuration-api client
var UpdateSensor = func(request SensorUpdateRequestModel) SensorResponseModel {
	// TODO: Query the sensor using the client

	return SensorResponseModel{
		UID:                "mock_uid",
		Serial:             "mock_serial",
		Name:               request.Name,
		ModelNumber:        "mock_model_number",
		WifiMacAddress:     "mock_wifi_mac_address",
		EthernetMacAddress: "mock_ethernet_mac_address",
		AddressNote:        request.AddressNote,
		Longitude:          "mock_longitude",
		Latitude:           "mock_latitude",
		Notes:              request.Notes,
		PCapMode:           request.PCapMode,
	}
}
