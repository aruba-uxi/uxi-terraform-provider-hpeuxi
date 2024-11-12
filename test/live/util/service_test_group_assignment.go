package util

import (
	"context"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
)

func GetServiceTestGroupAssignment(id string) config_api_client.ServiceTestGroupAssignmentsItem {
	result, _, err := Client.ConfigurationAPI.
		ServiceTestGroupAssignmentsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("service_test_group_assignment with id `" + id + "` could not be found")
	}
	return result.Items[0]
}

func CheckStateAgainstServiceTestGroupAssignment(
	t st.Fatalf,
	entity string,
	serviceTestGroupAssignment config_api_client.ServiceTestGroupAssignmentsItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", serviceTestGroupAssignment.Id),
		resource.TestCheckResourceAttr(entity, "group_id", serviceTestGroupAssignment.Group.Id),
		resource.TestCheckResourceAttr(
			entity,
			"service_test_id",
			serviceTestGroupAssignment.ServiceTest.Id,
		),
	)
}
