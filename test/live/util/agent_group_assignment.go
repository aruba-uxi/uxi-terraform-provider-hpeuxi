package util

import (
	"context"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
)

func GetAgentGroupAssignment(id string) config_api_client.AgentGroupAssignmentsItem {
	result, _, err := Client.ConfigurationAPI.
		AgentGroupAssignmentsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("agent_group_assignment with id `" + id + "` could not be found")
	}
	return result.Items[0]
}

func CheckStateAgainstAgentGroupAssignment(
	t st.Fatalf,
	entity string,
	agentGroupAssignment config_api_client.AgentGroupAssignmentsItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", agentGroupAssignment.Id),
		resource.TestCheckResourceAttr(entity, "group_id", agentGroupAssignment.Group.Id),
		resource.TestCheckResourceAttr(entity, "agent_id", agentGroupAssignment.Agent.Id),
	)
}
