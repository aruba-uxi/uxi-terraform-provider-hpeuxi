package provider

import (
	"github.com/aruba-uxi/terraform-provider-configuration/internal/provider"
	"github.com/aruba-uxi/terraform-provider-configuration/test/live/util"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	ProviderConfig = `provider "uxi" {
		host		  = "api.staging.capedev.io"
		client_id     = "` + util.CLIENT_ID + `"
		client_secret = "` + util.CLIENT_SECRET + `"
		token_url     = "https://sso.common.cloud.hpe.com/as/token.oauth2"
	}`
	// TestAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"uxi": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)
