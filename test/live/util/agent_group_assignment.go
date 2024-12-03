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

func GetAgentGroupAssignment(id string) *config_api_client.AgentGroupAssignmentsItem {
	result, response, err := Client.ConfigurationAPI.
		AgentGroupAssignmentsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if len(result.Items) != 1 {
		return nil
	}

	return &result.Items[0]
}

func CheckStateAgainstAgentGroupAssignment(
	t *testing.T,
	entity string,
	agentGroupAssignment config_api_client.AgentGroupAssignmentsItem,
) resource.TestCheckFunc {
	t.Helper()

	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", agentGroupAssignment.Id),
		resource.TestCheckResourceAttr(entity, "group_id", agentGroupAssignment.Group.Id),
		resource.TestCheckResourceAttr(entity, "agent_id", agentGroupAssignment.Agent.Id),
	)
}
