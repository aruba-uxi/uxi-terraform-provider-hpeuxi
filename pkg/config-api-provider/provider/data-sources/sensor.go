package datasources

import (
	"context"

	config_api_client "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/util"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &sensorDataSource{}
	_ datasource.DataSourceWithConfigure = &sensorDataSource{}
)

func NewSensorDataSource() datasource.DataSource {
	return &sensorDataSource{}
}

type sensorDataSource struct {
	client *config_api_client.APIClient
}

type sensorDataSourceModel struct {
	Id                 types.String  `tfsdk:"id"`
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
	Filter             struct {
		SensorID types.String `tfsdk:"sensor_id"`
	} `tfsdk:"filter"`
}

func (d *sensorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sensor"
}

func (d *sensorDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"serial": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"model_number": schema.StringAttribute{
				Computed: true,
			},
			"wifi_mac_address": schema.StringAttribute{
				Computed: true,
			},
			"ethernet_mac_address": schema.StringAttribute{
				Computed: true,
			},
			"address_note": schema.StringAttribute{
				Computed: true,
			},
			"longitude": schema.Float32Attribute{
				Computed: true,
			},
			"latitude": schema.Float32Attribute{
				Computed: true,
			},
			"notes": schema.StringAttribute{
				Computed: true,
			},
			"pcap_mode": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"sensor_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *sensorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state sensorDataSourceModel

	// Read configuration from request
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		GetNetworkingUxiV1alpha1SensorsGet(ctx).
		Id(state.Filter.SensorID.ValueString())

	sensorResponse, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_sensor")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(sensorResponse.Items) != 1 {
		resp.Diagnostics.AddError(errorSummary, "Could not find specified data source")
		return
	}

	sensor := sensorResponse.Items[0]

	state.Id = types.StringValue(sensor.Id)
	state.Name = types.StringValue(sensor.Name)
	state.ModelNumber = types.StringPointerValue(sensor.ModelNumber.Get())
	state.WifiMacAddress = types.StringPointerValue(sensor.WifiMacAddress.Get())
	state.EthernetMacAddress = types.StringPointerValue(sensor.EthernetMacAddress.Get())
	state.AddressNote = types.StringPointerValue(sensor.AddressNote.Get())
	state.Longitude = types.Float32PointerValue(sensor.Longitude.Get())
	state.Latitude = types.Float32PointerValue(sensor.Latitude.Get())
	state.Notes = types.StringPointerValue(sensor.Notes.Get())
	state.PcapMode = types.StringPointerValue(sensor.PcapMode.Get())

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *sensorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Data Source type: Sensor. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
