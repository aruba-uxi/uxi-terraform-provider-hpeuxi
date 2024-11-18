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
	_ datasource.DataSource              = &agentGroupAssignmentDataSource{}
	_ datasource.DataSourceWithConfigure = &agentGroupAssignmentDataSource{}
)

func NewAgentGroupAssignmentDataSource() datasource.DataSource {
	return &agentGroupAssignmentDataSource{}
}

type agentGroupAssignmentDataSource struct {
	client *config_api_client.APIClient
}

type agentGroupAssignmentDataSourceModel struct {
	ID      types.String `tfsdk:"id"`
	AgentID types.String `tfsdk:"agent_id"`
	GroupID types.String `tfsdk:"group_id"`
	Filter  struct {
		AgentGroupAssignmentID string `tfsdk:"agent_group_assignment_id"`
	} `tfsdk:"filter"`
}

func (d *agentGroupAssignmentDataSource) Metadata(
	_ context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_agent_group_assignment"
}

func (d *agentGroupAssignmentDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"agent_id": schema.StringAttribute{
				Computed: true,
			},
			"group_id": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"agent_group_assignment_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *agentGroupAssignmentDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var state agentGroupAssignmentDataSourceModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		AgentGroupAssignmentsGet(ctx).
		Id(state.Filter.AgentGroupAssignmentID)
	agentGroupAssignmentResponse, response, err := util.RetryForTooManyRequests(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_agent_group_assignment")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(agentGroupAssignmentResponse.Items) != 1 {
		resp.Diagnostics.AddError(errorSummary, "Could not find specified data source")
		return
	}

	agentGroupAssignment := agentGroupAssignmentResponse.Items[0]
	state.ID = types.StringValue(agentGroupAssignment.Id)
	state.AgentID = types.StringValue(agentGroupAssignment.Agent.Id)
	state.GroupID = types.StringValue(agentGroupAssignment.Group.Id)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *agentGroupAssignmentDataSource) Configure(
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
			"Data Source type: Agent Group Assignment. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
