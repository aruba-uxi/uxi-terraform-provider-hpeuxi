/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func IsRoot(group config_api_client.GroupsGetItem) bool {
	_, set := group.Parent.Get().GetIdOk()
	return !set
}
