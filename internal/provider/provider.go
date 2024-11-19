/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package provider

import (
	"context"
	"net/http"

	configuration "github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/datasources"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/resources"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	_ provider.Provider = &uxiConfigurationProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &uxiConfigurationProvider{
			version: version,
		}
	}
}

type uxiProviderModel struct {
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
}

type uxiConfigurationProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *uxiConfigurationProvider) Metadata(
	_ context.Context,
	_ provider.MetadataRequest,
	resp *provider.MetadataResponse,
) {
	resp.TypeName = "uxi"
	resp.Version = p.version
}

func (p *uxiConfigurationProvider) Schema(
	_ context.Context,
	_ provider.SchemaRequest,
	resp *provider.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Interact with HPE Aruba Network UXI Configuration.",
		Attributes: map[string]schema.Attribute{
			"client_id": schema.StringAttribute{
				Description: "The Client ID as obtained from HPE GreenLake API client credentials. " +
					"It is recommended that this configuration is left blank and the Client ID " +
					"is exported as the GREENLAKE_UXI_CLIENT_ID environment variable instead.",
				Optional: true,
			},
			"client_secret": schema.StringAttribute{
				Description: "The Client Secret as obtained from HPE GreenLake API client credentials. " +
					"It is recommended that this configuration is left blank and the Client Secret " +
					"is exported as the GREENLAKE_UXI_CLIENT_SECRET environment variable instead.",
				Optional:  true,
				Sensitive: true,
			},
		}}
}

func (p *uxiConfigurationProvider) Configure(
	ctx context.Context,
	req provider.ConfigureRequest,
	resp *provider.ConfigureResponse,
) {
	var (
		config                 uxiProviderModel
		clientID, clientSecret string
	)
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configuration.InitializeConfig()

	if !config.ClientID.IsNull() {
		clientID = config.ClientID.ValueString()
	} else {
		clientID = configuration.ClientID
	}

	if !config.ClientSecret.IsNull() {
		clientSecret = config.ClientSecret.ValueString()
	} else {
		clientSecret = configuration.ClientSecret
	}

	if clientID == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Missing Client ID",
			"The provider cannot initialize as there is a missing or empty value for the Client ID. "+
				"Set the Client ID value in the configuration or use the CLIENT_ID environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if clientSecret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Missing Client Secret",
			"The provider cannot initialize as there is a missing or empty value for the Client Secret. "+
				"Set the Client Secret value in the configuration or use the CLIENT_SECRET environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	uxiConfiguration := config_api_client.NewConfiguration()
	uxiConfiguration.Host = configuration.Host
	uxiConfiguration.Scheme = "https"
	uxiConfiguration.HTTPClient = getHttpClient(clientID, clientSecret, configuration.TokenURL)
	uxiClient := config_api_client.NewAPIClient(uxiConfiguration)

	resp.DataSourceData = uxiClient
	resp.ResourceData = uxiClient
}

func (p *uxiConfigurationProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewAgentDataSource,
		datasources.NewAgentGroupAssignmentDataSource,
		datasources.NewGroupDataSource,
		datasources.NewNetworkGroupAssignmentDataSource,
		datasources.NewSensorDataSource,
		datasources.NewSensorGroupAssignmentDataSource,
		datasources.NewServiceTestDataSource,
		datasources.NewServiceTestGroupAssignmentDataSource,
		datasources.NewWiredNetworkDataSource,
		datasources.NewWirelessNetworkDataSource,
	}
}

func (p *uxiConfigurationProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewAgentGroupAssignmentResource,
		resources.NewAgentResource,
		resources.NewGroupResource,
		resources.NewNetworkGroupAssignmentResource,
		resources.NewSensorGroupAssignmentResource,
		resources.NewSensorResource,
		resources.NewServiceTestGroupAssignmentResource,
		resources.NewServiceTestResource,
		resources.NewWiredNetworkResource,
		resources.NewWirelessNetworkResource,
	}
}

func getHttpClient(clientID string, clientSecret string, tokenURL string) *http.Client {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		AuthStyle:    oauth2.AuthStyleInParams,
	}

	return config.Client(context.Background())
}
