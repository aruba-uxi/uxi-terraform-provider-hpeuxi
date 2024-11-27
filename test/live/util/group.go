/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"context"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func GetGroupByName(name string) *config_api_client.GroupsGetItem {
	groups, response, err := Client.ConfigurationAPI.GroupsGet(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	for _, group := range groups.Items {
		if group.Name == name {
			return &group
		}
	}

	return nil
}
