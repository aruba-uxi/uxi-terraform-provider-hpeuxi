package datasources

import (
	"context"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &rootGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &rootGroupDataSource{}
)

func NewRootGroupDataSource() datasource.DataSource {
	return &rootGroupDataSource{}
}

type rootGroupDataSource struct{}

type rootGroupDataSourceModel struct {
	ID types.String `tfsdk:"id"`
}

func (d *rootGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_root_group"
}

func (d *rootGroupDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

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

func (d *rootGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}
}

// TODO: Switch this to use the Client Model when that becomes available
var GetRootGroup = func() resources.GroupResponseModel {
	return resources.GroupResponseModel{
		UID:       "mock_uid",
		Name:      "root",
		ParentUid: nil,
		Path:      "mock_uid",
	}
}
