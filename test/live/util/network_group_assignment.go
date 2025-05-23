/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func GetNetworkGroupAssignment(id string) config_api_client.NetworkGroupAssignmentsGetItem {
	result, response, err := Client.ConfigurationAPI.
		NetworkGroupAssignmentsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if len(result.Items) != 1 {
		panic("network_group_assignment with id `" + id + "` could not be found")
	}

	return result.Items[0]
}

func CheckStateAgainstNetworkGroupAssignment(
	t *testing.T,
	entity string,
	networkGroupAssignment config_api_client.NetworkGroupAssignmentsGetItem,
) resource.TestCheckFunc {
	t.Helper()

	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", networkGroupAssignment.Id),
		resource.TestCheckResourceAttr(entity, "group_id", networkGroupAssignment.Group.Id),
		resource.TestCheckResourceAttr(entity, "network_id", networkGroupAssignment.Network.Id),
	)
}
