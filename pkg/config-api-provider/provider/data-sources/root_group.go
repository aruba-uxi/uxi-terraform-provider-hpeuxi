package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &rootGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &rootGroupDataSource{}
)

// NewRootGroupDataSource is a helper function to simplify the provider implementation.
func NewRootGroupDataSource() datasource.DataSource {
	return &rootGroupDataSource{}
}

// rootGroupDataSource is the data source implementation.
type rootGroupDataSource struct{}

// rootGroupDataSourceModel maps the data source schema data.
type rootGroupDataSourceModel struct {
	ID types.String `tfsdk:"id"`
}

// Metadata returns the data source type name.
func (d *rootGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_root_group"
}

// Schema defines the schema for the data source.
func (d *rootGroupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *rootGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state rootGroupDataSourceModel

	rootGroup := GetRootGroup()

	state.ID = types.StringValue(rootGroup.UID)

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *rootGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}
}

// TODO: Switch this to use the Client Model when that becomes available
type RootGroupResponseModel struct {
	UID       string
	Name      string
	ParentUid *string
	Path      string
}

var GetRootGroup = func() RootGroupResponseModel {
	return RootGroupResponseModel{
		UID:       "mock_uid",
		Name:      "root",
		ParentUid: nil,
		Path:      "mock_uid",
	}
}
