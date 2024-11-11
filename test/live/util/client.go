package util

import (
	"context"
	"os"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var CLIENT_ID = os.Getenv("UXI_CLIENT_ID")
var CLIENT_SECRET = os.Getenv("UXI_CLIENT_SECRET")

const HOST = "api.staging.capedev.io"
const TOKEN_URL = "https://sso.common.cloud.hpe.com/as/token.oauth2"

func NewClient() *config_api_client.APIClient {
	config := &clientcredentials.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		TokenURL:     TOKEN_URL,
		AuthStyle:    oauth2.AuthStyleInParams,
	}

	// Create a context and fetch a token
	uxiConfiguration := config_api_client.NewConfiguration()
	uxiConfiguration.Host = HOST
	uxiConfiguration.Scheme = "https"
	uxiConfiguration.HTTPClient = config.Client(context.Background())

	return config_api_client.NewAPIClient(uxiConfiguration)
}

var Client = NewClient()
