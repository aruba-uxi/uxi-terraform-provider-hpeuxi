/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	configuration "github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/config"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/resources"
	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var _ provider.Provider = &hpeuxiConfigurationProvider{}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &hpeuxiConfigurationProvider{
			version: version,
		}
	}
}

type hpeuxiProviderModel struct {
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
}

type hpeuxiConfigurationProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *hpeuxiConfigurationProvider) Metadata(
	_ context.Context,
	_ provider.MetadataRequest,
	resp *provider.MetadataResponse,
) {
	resp.TypeName = "hpeuxi"
	resp.Version = p.version
}

func (p *hpeuxiConfigurationProvider) Schema(
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
					"is exported as the `HPEUXI_CLIENT_ID` environment variable instead.",
				Optional: true,
			},
			"client_secret": schema.StringAttribute{
				Description: "The Client Secret as obtained from HPE GreenLake API client credentials. " +
					"It is recommended that this configuration is left blank and the Client Secret " +
					"is exported as the `HPEUXI_CLIENT_SECRET` environment variable instead.",
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *hpeuxiConfigurationProvider) Configure(
	ctx context.Context,
	req provider.ConfigureRequest,
	resp *provider.ConfigureResponse,
) {
	var (
		config                 hpeuxiProviderModel
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
				"Set the client_id value in the provider configuration or use the HPEUXI_CLIENT_ID "+
				"environment variable (recommended). If either is already set, ensure the value is not empty.",
		)
	}

	if clientSecret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Missing Client Secret",
			"The provider cannot initialize as there is a missing or empty value for the Client Secret. "+
				"Set the client_secret value in the provider configuration or use the HPEUXI_CLIENT_SECRET "+
				"environment variable (recommended). If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	hpeuxiConfiguration := config_api_client.NewConfiguration()
	hpeuxiConfiguration.Host = configuration.Host
	hpeuxiConfiguration.Scheme = "https"
	hpeuxiConfiguration.HTTPClient = getHTTPClient(clientID, clientSecret, configuration.TokenURL)
	uxiClient := config_api_client.NewAPIClient(hpeuxiConfiguration)

	resp.DataSourceData = uxiClient
	resp.ResourceData = uxiClient
}

func (p *hpeuxiConfigurationProvider) DataSources(
	_ context.Context,
) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		resources.NewAgentDataSource,
		resources.NewAgentGroupAssignmentDataSource,
		resources.NewGroupDataSource,
		resources.NewNetworkGroupAssignmentDataSource,
		resources.NewSensorDataSource,
		resources.NewSensorGroupAssignmentDataSource,
		resources.NewServiceTestDataSource,
		resources.NewServiceTestGroupAssignmentDataSource,
		resources.NewWiredNetworkDataSource,
		resources.NewWirelessNetworkDataSource,
	}
}

func (p *hpeuxiConfigurationProvider) Resources(_ context.Context) []func() resource.Resource {
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

func getHTTPClient(clientID string, clientSecret string, tokenURL string) *http.Client {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		AuthStyle:    oauth2.AuthStyleInParams,
	}

	return config.Client(context.Background())
}
