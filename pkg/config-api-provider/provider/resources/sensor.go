package resources

import (
	"context"

	"time"

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
	LastUpdated types.String `tfsdk:"last_updated"`
}

type sensorResponseModel struct {
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

func NewSensorResource() resource.Resource {
	return &sensorResource{}
}

type sensorResource struct{}

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
			"last_updated": schema.StringAttribute{
				Computed: true,
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
}

func (r *sensorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan sensorResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "creating a sensor resource is not supported; sensors can only be imported")
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

	// TODO: Call client create-sensor method
	// We are mocking the response of the client for this early stage of development
	response := GetSensor()

	// Update state from client response
	state.Name = types.StringValue(response.Name)
	state.AddressNote = types.StringValue(response.AddressNote)
	state.Notes = types.StringValue(response.Notes)
	state.PCapMode = types.StringValue(response.PCapMode)
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

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

	// TODO: Call client update-sensor method

	// Update the state to match the plan (replace with response from client)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

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
	diags.AddError("operation not supported", "deleting a sensor resource is not supported; sensors can only removed from state")
	resp.Diagnostics.Append(diags...)
}

func (r *sensorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Get the sensor using the configuration-api client
func GetSensor() sensorResponseModel {
	// TODO: Query the sensor using the client

	return sensorResponseModel{
		UID:                "temporary_uid",
		Serial:             "temporary_serial",
		Name:               "temporary_name",
		ModelNumber:        "temporary_model_number",
		WifiMacAddress:     "temporary_wifi_mac_address",
		EthernetMacAddress: "temporary_ethernet_mac_address",
		AddressNote:        "temporary_address_note",
		Longitude:          "temporary_longitude",
		Latitude:           "temporary_latitude",
		Notes:              "temporary_notes",
		PCapMode:           "temporary_pcap_mode",
	}
}
