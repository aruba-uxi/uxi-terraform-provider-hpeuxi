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
	_ resource.Resource              = &serviceTestResource{}
	_ resource.ResourceWithConfigure = &serviceTestResource{}
)

type serviceTestResourceModel struct {
	ID    types.String `tfsdk:"id"`
	Title types.String `tfsdk:"title"`
}

// TODO: Switch this to use the Client Model when that becomes available
type ServiceTestResponseModel struct {
	Uid       string // <service_test_uid:str>,
	Category  string // <category:str>,
	Title     string // <title:str>,
	Target    string // Nullable<<target:str>>,
	Template  string // <template:str>,
	IsEnabled bool   // <is_enabled:bool>
}

func NewServiceTestResource() resource.Resource {
	return &serviceTestResource{}
}

type serviceTestResource struct{}

func (r *serviceTestResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_test"
}

func (r *serviceTestResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"title": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *serviceTestResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
}

func (r *serviceTestResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan serviceTestResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "creating a service_test is not supported; service_tests can only be imported")
	resp.Diagnostics.Append(diags...)
}

func (r *serviceTestResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state serviceTestResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	response := GetServiceTest()

	// Update state from client response
	state.Title = types.StringValue(response.Title)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serviceTestResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan serviceTestResourceModel
	diags := req.Plan.Get(ctx, &plan)
	diags.AddError("operation not supported", "updating a service_test is not supported; service_tests can only be updated through the dashboard")
	resp.Diagnostics.Append(diags...)
}

func (r *serviceTestResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state serviceTestResourceModel
	diags := req.State.Get(ctx, &state)
	diags.AddError("operation not supported", "deleting a service_test is not supported; service_tests can only removed from state")
	resp.Diagnostics.Append(diags...)
}

func (r *serviceTestResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Get the serviceTest using the configuration-api client
var GetServiceTest = func() ServiceTestResponseModel {
	// TODO: Query the serviceTest using the client

	return ServiceTestResponseModel{
		Uid:       "uid",
		Category:  "category",
		Title:     "title",
		Target:    "target",
		Template:  "template",
		IsEnabled: true,
	}
}
