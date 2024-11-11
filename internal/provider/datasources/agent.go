package datasources

import (
	"context"

	config_api_client "github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-configuration/internal/provider/util"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &agentDataSource{}
	_ datasource.DataSourceWithConfigure = &agentDataSource{}
)

func NewAgentDataSource() datasource.DataSource {
	return &agentDataSource{}
}

type agentDataSource struct {
	client *config_api_client.APIClient
}

type agentDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Serial             types.String `tfsdk:"serial"`
	Name               types.String `tfsdk:"name"`
	ModelNumber        types.String `tfsdk:"model_number"`
	WifiMacAddress     types.String `tfsdk:"wifi_mac_address"`
	EthernetMacAddress types.String `tfsdk:"ethernet_mac_address"`
	Notes              types.String `tfsdk:"notes"`
	PcapMode           types.String `tfsdk:"pcap_mode"`
	Filter             struct {
		AgentID types.String `tfsdk:"agent_id"`
	} `tfsdk:"filter"`
}

func (d *agentDataSource) Metadata(
	_ context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_agent"
}

func (d *agentDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
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
			"notes": schema.StringAttribute{
				Computed: true,
			},
			"pcap_mode": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"agent_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *agentDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var state agentDataSourceModel

	// Read configuration from request
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		AgentsGet(ctx).
		Id(state.Filter.AgentID.ValueString())

	agentResponse, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_agent")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(agentResponse.Items) != 1 {
		resp.Diagnostics.AddError(errorSummary, "Could not find specified data source")
		return
	}

	agent := agentResponse.Items[0]

	state.Id = types.StringValue(agent.Id)
	state.Name = types.StringValue(agent.Name)
	state.Serial = types.StringValue(agent.Serial)
	state.ModelNumber = types.StringPointerValue(agent.ModelNumber.Get())
	state.WifiMacAddress = types.StringPointerValue(agent.WifiMacAddress.Get())
	state.EthernetMacAddress = types.StringPointerValue(agent.EthernetMacAddress.Get())
	state.Notes = types.StringPointerValue(agent.Notes.Get())
	state.PcapMode = types.StringPointerValue(agent.PcapMode.Get())

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *agentDataSource) Configure(
	_ context.Context,
	req datasource.ConfigureRequest,
	resp *datasource.ConfigureResponse,
) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Data Source type: Agent. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
