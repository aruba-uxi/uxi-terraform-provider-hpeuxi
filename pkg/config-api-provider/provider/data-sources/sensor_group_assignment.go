package datasources

import (
	"context"

	config_api_client "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/util"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
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

func (d *sensorGroupAssignmentDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sensor_group_assignment"
}

func (d *sensorGroupAssignmentDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"sensor_id": schema.StringAttribute{
				Computed: true,
			},
			"group_id": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"sensor_group_assignment_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *sensorGroupAssignmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state sensorGroupAssignmentDataSourceModel

	// Read configuration from request
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		GetUxiV1alpha1SensorGroupAssignmentsGet(ctx).
		Uid(state.Filter.SensorGroupAssignmentID)
	sensorGroupAssignmentResponse, _, err := util.RetryFor429(request.Execute)

	if err != nil || len(sensorGroupAssignmentResponse.Items) != 1 {
		resp.Diagnostics.AddError(
			"Error reading Sensor Group Assignment",
			"Could not retrieve Sensor Group Assignment, unexpected error: "+err.Error(),
		)
		return
	}

	sensorGroupAssignment := sensorGroupAssignmentResponse.Items[0]
	state.ID = types.StringValue(sensorGroupAssignment.Id)
	state.SensorID = types.StringValue(sensorGroupAssignment.Sensor.Id)
	state.GroupID = types.StringValue(sensorGroupAssignment.Group.Id)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *sensorGroupAssignmentDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
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
