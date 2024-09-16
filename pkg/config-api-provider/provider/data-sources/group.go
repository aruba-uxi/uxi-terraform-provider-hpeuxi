package datasources

import (
	"context"

	config_api_client "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &groupDataSource{}
	_ datasource.DataSourceWithConfigure = &groupDataSource{}
)

func NewGroupDataSource() datasource.DataSource {
	return &groupDataSource{}
}

type groupDataSource struct {
	client *config_api_client.APIClient
}

type groupDataSourceModel struct {
	ID            types.String `tfsdk:"id"`
	Path          types.String `tfsdk:"path"`
	ParentGroupID types.String `tfsdk:"parent_group_id"`
	Name          types.String `tfsdk:"name"`
	Filter        struct {
		GroupID *string `tfsdk:"group_id"`
		IsRoot  *bool   `tfsdk:"is_root"`
	} `tfsdk:"filter"`
}

func (d *groupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (d *groupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"path": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"parent_group_id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"group_id": schema.StringAttribute{
						Optional: true,
					},
					"is_root": schema.BoolAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

func (d *groupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state groupDataSourceModel

	// Read configuration from request
	diags := req.Config.Get(ctx, &state)
	if state.Filter.GroupID == nil && (state.Filter.IsRoot == nil || !*state.Filter.IsRoot) {
		diags.AddError("invalid Group data source", "either filter.group_id must be set or 'filter.is_root = true' is required")
	} else if state.Filter.GroupID != nil && state.Filter.IsRoot != nil && *state.Filter.IsRoot {
		diags.AddError("invalid Group data source", "group_id and 'is_root = true' cannot both be set")
	}
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := d.client.ConfigurationAPI.GroupsGetConfigurationAppV1GroupsGet(context.Background())

	if state.Filter.IsRoot != nil && *state.Filter.IsRoot {
		request = request.Uid(*state.Filter.GroupID) // TODO: use root group filter here
	} else {
		request = request.Uid(*state.Filter.GroupID)
	}

	groupResponse, _, err := request.Execute()

	if err != nil || len(groupResponse.Groups) != 1 {
		resp.Diagnostics.AddError(
			"Error reading Group",
			"Could not retrieve Group, unexpected error: "+err.Error(),
		)
		return
	}

	group := groupResponse.Groups[0]
	state.ID = types.StringValue(group.Uid)
	state.Name = types.StringValue(group.Name)
	state.Path = types.StringValue(group.Path)
	if group.ParentUid.IsSet() {
		state.ParentGroupID = types.StringValue(*group.ParentUid.Get())
	}

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *groupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*config_api_client.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Data Source type: Group. Please report this issue to the provider developers.",
		)
		return
	}

	d.client = client
}
