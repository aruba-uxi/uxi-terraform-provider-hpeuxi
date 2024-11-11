package util

import (
	"context"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func GetGroupByName(name string) *config_api_client.GroupsGetItem {
	groups, _, _ := Client.ConfigurationAPI.GroupsGet(context.Background()).Execute()
	for _, group := range groups.Items {
		if group.Name == name {
			return &group
		}
	}
	return nil
}
