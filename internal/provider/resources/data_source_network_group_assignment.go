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
	_ datasource.DataSource              = &networkGroupAssignmentDataSource{}
	_ datasource.DataSourceWithConfigure = &networkGroupAssignmentDataSource{}
)

func NewNetworkGroupAssignmentDataSource() datasource.DataSource {
	return &networkGroupAssignmentDataSource{}
}

type networkGroupAssignmentDataSource struct {
	client *config_api_client.APIClient
}

type networkGroupAssignmentDataSourceModel struct {
	ID        types.String `tfsdk:"id"`
	NetworkID types.String `tfsdk:"network_id"`
	GroupID   types.String `tfsdk:"group_id"`
	Filter    struct {
		ID string `tfsdk:"id"`
	} `tfsdk:"filter"`
}

func (d *networkGroupAssignmentDataSource) Metadata(
	_ context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_network_group_assignment"
}

func (d *networkGroupAssignmentDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Retrieves a specific network group assignment.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The identifier of the network group assignment.",
				Computed:    true,
			},
			"network_id": schema.StringAttribute{
				Description: "The identifier of the assigned wired or wireless network.",
				Computed:    true,
			},
			"group_id": schema.StringAttribute{
				Description: "The identifier of the assigned group.",
				Computed:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Description: "The filter used to filter the specific network group assignment.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The identifier of the network group assignment.",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *networkGroupAssignmentDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var state networkGroupAssignmentDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		NetworkGroupAssignmentsGet(ctx).
		Id(state.Filter.ID)
	networkGroupAssignmentResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	defer response.Body.Close()
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_network_group_assignment")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)

		return
	}

	if len(networkGroupAssignmentResponse.Items) != 1 {
		resp.Diagnostics.AddError(errorSummary, "Could not find specified data source")
		resp.State.RemoveResource(ctx)

		return
	}

	networkGroupAssignment := networkGroupAssignmentResponse.Items[0]
	state.ID = types.StringValue(networkGroupAssignment.Id)
	state.NetworkID = types.StringValue(networkGroupAssignment.Network.Id)
	state.GroupID = types.StringValue(networkGroupAssignment.Group.Id)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *networkGroupAssignmentDataSource) Configure(
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
			"Data Source type: Network Group Assignment. Please report this issue to the provider developers.",
		)

		return
	}

	d.client = client
}
