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
	_ datasource.DataSource              = &serviceTestDataSource{}
	_ datasource.DataSourceWithConfigure = &serviceTestDataSource{}
)

func NewServiceTestDataSource() datasource.DataSource {
	return &serviceTestDataSource{}
}

type serviceTestDataSource struct {
	client *config_api_client.APIClient
}

type serviceTestDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Category  types.String `tfsdk:"category"`
	Name      types.String `tfsdk:"name"`
	Target    types.String `tfsdk:"target"`
	Template  types.String `tfsdk:"template"`
	IsEnabled types.String `tfsdk:"is_enabled"`
	Filter    struct {
		ServiceTestID types.String `tfsdk:"service_test_id"`
	} `tfsdk:"filter"`
}

func (d *serviceTestDataSource) Metadata(
	_ context.Context,
	req datasource.MetadataRequest,
	resp *datasource.MetadataResponse,
) {
	resp.TypeName = req.ProviderTypeName + "_service_test"
}

func (d *serviceTestDataSource) Schema(
	_ context.Context,
	_ datasource.SchemaRequest,
	resp *datasource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"category": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"target": schema.StringAttribute{
				Computed: true,
			},
			"template": schema.StringAttribute{
				Computed: true,
			},
			"is_enabled": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"service_test_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *serviceTestDataSource) Read(
	ctx context.Context,
	req datasource.ReadRequest,
	resp *datasource.ReadResponse,
) {
	var state serviceTestDataSourceModel

	// Read configuration from request
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.
		ServiceTestsGet(ctx).
		Id(state.Filter.ServiceTestID.ValueString())

	serviceTestResponse, response, err := util.RetryFor429(request.Execute)
	errorPresent, errorDetail := util.RaiseForStatus(response, err)

	errorSummary := util.GenerateErrorSummary("read", "uxi_service_test")

	if errorPresent {
		resp.Diagnostics.AddError(errorSummary, errorDetail)
		return
	}

	if len(serviceTestResponse.Items) != 1 {
		resp.Diagnostics.AddError(errorSummary, "Could not find specified data source")
		return
	}

	serviceTest := serviceTestResponse.Items[0]

	state.Id = types.StringValue(serviceTest.Id)
	state.Name = types.StringValue(serviceTest.Name)

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *serviceTestDataSource) Configure(
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
			"Data Source type: ServiceTest. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
