package provider

import (
	"context"
	"net/http"
	"os"

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

const tokenURLDefault = "https://sso.common.cloud.hpe.com/as/token.oauth2"

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &uxiConfigurationProvider{
			version: version,
		}
	}
}

type uxiProviderModel struct {
	Host         types.String `tfsdk:"host"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
	TokenURL     types.String `tfsdk:"token_url"`
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
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"host":          schema.StringAttribute{Optional: true},
		"client_id":     schema.StringAttribute{Optional: true},
		"client_secret": schema.StringAttribute{Optional: true, Sensitive: true},
		"token_url":     schema.StringAttribute{Optional: true},
	}}
}

func (p *uxiConfigurationProvider) Configure(
	ctx context.Context,
	req provider.ConfigureRequest,
	resp *provider.ConfigureResponse,
) {
	var config uxiProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("UXI_HOST")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	tokenURL := os.Getenv("TOKEN_URL")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.ClientID.IsNull() {
		clientID = config.ClientID.ValueString()
	}

	if !config.ClientSecret.IsNull() {
		clientSecret = config.ClientSecret.ValueString()
	}

	if !config.TokenURL.IsNull() {
		tokenURL = config.TokenURL.ValueString()
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Host",
			"The provider cannot initialize as there is a missing or empty value for the UXI API host. "+
				"Set the host value in the configuration or use the UXI_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
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

	if tokenURL == "" {
		tokenURL = tokenURLDefault
	}

	if resp.Diagnostics.HasError() {
		return
	}

	uxiConfiguration := config_api_client.NewConfiguration()
	uxiConfiguration.Host = host
	uxiConfiguration.Scheme = "https"
	uxiConfiguration.HTTPClient = getHttpClient(clientID, clientSecret, tokenURL)
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
