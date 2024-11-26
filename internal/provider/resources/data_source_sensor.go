/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/util"
	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
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
	Filter             struct {
		ID types.String `tfsdk:"id"`
	} `tfsdk:"filter"`
}

func (d *sensorDataSource) Metadata(
	_ context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_sensor"
}

func (d *sensorDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Retrieves a specific sensor.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the sensor.",
				Computed:    true,
			},
			"serial": schema.StringAttribute{
				Description: "The serial number of the sensor.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the sensor.",
				Computed:    true,
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
				Description: "The address note of the sensor.",
				Computed:    true,
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
				Description: "The notes of the sensor.",
				Computed:    true,
			},
			"pcap_mode": schema.StringAttribute{
				Description: "The packet capture mode of the sensor.",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Description: "The filter used to filter the specific sensor.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The identifier of the sensor group assignment.",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *sensorDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var state sensorDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		SensorsGet(ctx).
		Id(state.Filter.ID.ValueString())

	sensorResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	// defer response.Body.Close()
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_sensor")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)

		return
	}

	if len(sensorResponse.Items) != 1 {
		resp.Diagnostics.AddError(errorSummary, "Could not find specified data source")
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
	state.PcapMode = types.StringPointerValue(sensor.PcapMode.Get())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *sensorDataSource) Configure(
	_ context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
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
