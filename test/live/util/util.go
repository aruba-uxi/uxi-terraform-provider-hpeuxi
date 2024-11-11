package util

import (
	"context"
	"os"

	"github.com/aruba-uxi/terraform-provider-configuration-api/pkg/config-api-client"
	resources_util "github.com/aruba-uxi/terraform-provider-configuration/internal/provider/util"
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

func GetRoot() *config_api_client.GroupsGetItem {
	groups, _, _ := Client.ConfigurationAPI.GroupsGet(context.Background()).Execute()
	for _, group := range groups.Items {
		if resources_util.IsRoot(group) {
			return &group
		}
	}
	return nil
}

func GetGroupByName(name string) *config_api_client.GroupsGetItem {
	groups, _, _ := Client.ConfigurationAPI.GroupsGet(context.Background()).Execute()
	for _, group := range groups.Items {
		if group.Name == name {
			return &group
		}
	}
	return nil
}

func ConditionalProperty(property string, value *string) string {
	if value == nil {
		return ""
	}
	return property + `= "` + *value + `"`
}

var Client = NewClient()
