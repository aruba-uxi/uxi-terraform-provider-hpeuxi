/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package datasources

import (
	"context"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &sensorGroupAssignmentDataSource{}
	_ datasource.DataSourceWithConfigure = &sensorGroupAssignmentDataSource{}
)

func NewSensorGroupAssignmentDataSource() datasource.DataSource {
	return &sensorGroupAssignmentDataSource{}
}

type sensorGroupAssignmentDataSource struct {
	client *config_api_client.APIClient
}

type sensorGroupAssignmentDataSourceModel struct {
	ID       types.String `tfsdk:"id"`
	SensorID types.String `tfsdk:"sensor_id"`
	GroupID  types.String `tfsdk:"group_id"`
	Filter   struct {
		SensorGroupAssignmentID string `tfsdk:"sensor_group_assignment_id"`
	} `tfsdk:"filter"`
}

func (d *sensorGroupAssignmentDataSource) Metadata(
	_ context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_sensor_group_assignment"
}

func (d *sensorGroupAssignmentDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The identifier of the network group assignment.",
			},
			"sensor_id": schema.StringAttribute{
				Computed:    true,
				Description: "The identifier of the assigned sensor.",
			},
			"group_id": schema.StringAttribute{
				Computed:    true,
				Description: "The identifier of the assigned group.",
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"sensor_group_assignment_id": schema.StringAttribute{
						Required:    true,
						Description: "The identifier of the sensor group assignment.",
					},
				},
				Description: "The filter used to filter the specific sensor group assignment.",
			},
		},
	}
}

func (d *sensorGroupAssignmentDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var state sensorGroupAssignmentDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		SensorGroupAssignmentsGet(ctx).
		Id(state.Filter.SensorGroupAssignmentID)
	sensorGroupAssignmentResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_sensor_group_assignment")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(sensorGroupAssignmentResponse.Items) != 1 {
		resp.Diagnostics.AddError(errorSummary, "Could not find specified data source")
		resp.State.RemoveResource(ctx)
		return
	}

	sensorGroupAssignment := sensorGroupAssignmentResponse.Items[0]
	state.ID = types.StringValue(sensorGroupAssignment.Id)
	state.SensorID = types.StringValue(sensorGroupAssignment.Sensor.Id)
	state.GroupID = types.StringValue(sensorGroupAssignment.Group.Id)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *sensorGroupAssignmentDataSource) Configure(
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
			"Data Source type: Sensor Group Assignment. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
