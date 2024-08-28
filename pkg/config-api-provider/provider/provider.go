package provider

import (
	"context"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/terraform-provider-configuration/provider/resources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &uxiConfigurationProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &uxiConfigurationProvider{
			version: version,
		}
	}
}

type uxiConfigurationProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *uxiConfigurationProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "uxi"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *uxiConfigurationProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// TODO: Obtain a greenlake access token
// Configure prepares a Configuration API client for data sources and resources.
func (p *uxiConfigurationProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Init
	// initialise client
}

// DataSources defines the data sources implemented in the provider.
func (p *uxiConfigurationProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *uxiConfigurationProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewAgentResource,
		resources.NewGroupResource,
		resources.NewSensorResource,
		resources.NewWiredNetworkResource,
		resources.NewWirelessNetworkResource,
		resources.NewServiceTestResource,
		resources.NewSensorGroupAssignmentResource,
	}
}
