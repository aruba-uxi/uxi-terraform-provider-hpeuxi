/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"context"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func GetServiceTest(id string) config_api_client.ServiceTestsListItem {
	result, _, err := Client.ConfigurationAPI.
		ServiceTestsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("service_test with id `" + id + "` could not be found")
	}

	return result.Items[0]
}
