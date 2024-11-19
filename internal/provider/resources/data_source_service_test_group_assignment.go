/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package resources

import (
	"context"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/util"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &serviceTestGroupAssignmentDataSource{}
	_ datasource.DataSourceWithConfigure = &serviceTestGroupAssignmentDataSource{}
)

func NewServiceTestGroupAssignmentDataSource() datasource.DataSource {
	return &serviceTestGroupAssignmentDataSource{}
}

type serviceTestGroupAssignmentDataSource struct {
	client *config_api_client.APIClient
}

type serviceTestGroupAssignmentDataSourceModel struct {
	ID            types.String `tfsdk:"id"`
	ServiceTestID types.String `tfsdk:"service_test_id"`
	GroupID       types.String `tfsdk:"group_id"`
	Filter        struct {
		ServiceTestGroupAssignmentID string `tfsdk:"service_test_group_assignment_id"`
	} `tfsdk:"filter"`
}

func (d *serviceTestGroupAssignmentDataSource) Metadata(
	_ context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_service_test_group_assignment"
}

func (d *serviceTestGroupAssignmentDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"service_test_id": schema.StringAttribute{
				Computed: true,
			},
			"group_id": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"service_test_group_assignment_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *serviceTestGroupAssignmentDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var state serviceTestGroupAssignmentDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		ServiceTestGroupAssignmentsGet(ctx).
		Id(state.Filter.ServiceTestGroupAssignmentID)
	serviceTestGroupAssignmentResponse, response, err := util.RetryForTooManyRequests(
		request.Execute,
	)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_service_test_group_assignment")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(serviceTestGroupAssignmentResponse.Items) != 1 {
		resp.Diagnostics.AddError(errorSummary, "Could not find specified data source")
		resp.State.RemoveResource(ctx)
		return
	}

	serviceTestGroupAssignment := serviceTestGroupAssignmentResponse.Items[0]
	state.ID = types.StringValue(serviceTestGroupAssignment.Id)
	state.ServiceTestID = types.StringValue(serviceTestGroupAssignment.ServiceTest.Id)
	state.GroupID = types.StringValue(serviceTestGroupAssignment.Group.Id)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *serviceTestGroupAssignmentDataSource) Configure(
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
			"Data Source type: ServiceTest Group Assignment. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
