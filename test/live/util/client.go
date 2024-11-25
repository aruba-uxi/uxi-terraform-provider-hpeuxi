/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"context"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/internal/provider/config"
	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func NewClient() *config_api_client.APIClient {
	config.InitializeConfig()
	clientConfig := &clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     config.TokenURL,
		AuthStyle:    oauth2.AuthStyleInParams,
	}

	// Create a context and fetch a token
	uxiConfiguration := config_api_client.NewConfiguration()
	uxiConfiguration.Host = config.Host
	uxiConfiguration.Scheme = "https"
	uxiConfiguration.HTTPClient = clientConfig.Client(context.Background())

	return config_api_client.NewAPIClient(uxiConfiguration)
}

var Client = NewClient()
