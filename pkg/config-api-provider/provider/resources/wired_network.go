package resources

import (
	"context"

	// "github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &wiredNetworkResource{}
	_ resource.ResourceWithConfigure = &wiredNetworkResource{}
)

type wiredNetworkResourceModel struct {
	ID    types.String `tfsdk:"id"`
	Alias types.String `tfsdk:"alias"`
}

// TODO: Switch this to use the Client Model when that becomes available
type WiredNetworkResponseModel struct {
	Uid                  string // <network_uid:str>,
	Alias                string // <network_alias>,
	DatetimeCreated      string // <created_at:str(isoformat(with timezone?))>,
	DatetimeUpdated      string // <updated_at:str(isoformat(with timezone?))>,
	IpVersion            string // <ip_version:str>,
	Security             string // <security:str>,
	DnsLookupDomain      string // <dns_lookup_domain:str> | Nullable,
	DisableEdns          bool   // <disable_edns:bool>,
	UseDns64             bool   // <use_dns64:bool>,
	ExternalConnectivity bool   // <external_connectivity:bool>
	VlanId               int    // <vlan_id:int>
}

func NewWiredNetworkResource() resource.Resource {
	return &wiredNetworkResource{}
}

type wiredNetworkResource struct{}

func (r *wiredNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wired_network"
}

func (r *wiredNetworkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"alias": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *wiredNetworkResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
}

func (r *wiredNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan wiredNetworkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "creating a wired_network is not supported; wired_networks can only be imported")
	resp.Diagnostics.Append(diags...)
}

func (r *wiredNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state wiredNetworkResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response := GetWiredNetwork()

	// Update state from client response
	state.Alias = types.StringValue(response.Alias)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *wiredNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan wiredNetworkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "updating a wired_network is not supported; wired_networks can only be updated through the dashboard")
	resp.Diagnostics.Append(diags...)
}

func (r *wiredNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state wiredNetworkResourceModel
	diags := req.State.Get(ctx, &state)
	diags.AddError("operation not supported", "deleting a wired_network is not supported; wired_networks can only removed from state")
	resp.Diagnostics.Append(diags...)
}

func (r *wiredNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Get the wiredNetwork using the configuration-api client
var GetWiredNetwork = func() WiredNetworkResponseModel {
	// TODO: Query the wiredNetwork using the client

	return WiredNetworkResponseModel{
		Uid:                  "mock_uid",
		Alias:                "mock_alias",
		DatetimeCreated:      "mock_datetime_created",
		DatetimeUpdated:      "mock_datetime_updated",
		IpVersion:            "mock_ip_version",
		Security:             "mock_security",
		DnsLookupDomain:      "mock_dns_lookup_domain",
		DisableEdns:          false,
		UseDns64:             false,
		ExternalConnectivity: false,
		VlanId:               123,
	}
}
