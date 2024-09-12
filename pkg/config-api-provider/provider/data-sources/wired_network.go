package datasources

import (
	"context"

	config_api_client "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &wiredNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &wiredNetworkDataSource{}
)

func NewWiredNetworkDataSource() datasource.DataSource {
	return &wiredNetworkDataSource{}
}

type wiredNetworkDataSource struct {
	client *config_api_client.APIClient
}

type wiredNetworkDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Filter struct {
		WiredNetworkID string `tfsdk:"wired_network_id"`
	} `tfsdk:"filter"`
}

func (d *wiredNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wired_network"
}

func (d *wiredNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"wired_network_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *wiredNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state wiredNetworkDataSourceModel

	// Read configuration from request
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	networkResponse, _, err := d.client.ConfigurationAPI.
		GetConfigurationAppV1WiredNetworksGet(context.Background()).
		Uid(state.Filter.WiredNetworkID).
		Execute()

	if err != nil || len(networkResponse.WiredNetworks) != 1 {
		resp.Diagnostics.AddError(
			"Error reading Wired Network",
			"Could not retrieve Wired Network, unexpected error: "+err.Error(),
		)
		return
	}

	network := networkResponse.WiredNetworks[0]
	state.ID = types.StringValue(network.Uid)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *wiredNetworkDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Data Source type: Wired Network. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
