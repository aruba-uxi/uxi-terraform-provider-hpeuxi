/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package shared

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	AgentGroupAssignmentType       = "networking-uxi/agent-group-assignment"
	AgentType                      = "networking-uxi/agent"
	GroupType                      = "networking-uxi/group"
	NetworkGroupAssignmentType     = "networking-uxi/network-group-assignment"
	SensorGroupAssignmentType      = "networking-uxi/sensor-group-assignment"
	SensorType                     = "networking-uxi/sensor"
	ServiceTestGroupAssignmentType = "networking-uxi/service-test-group-assignment"
	ServiceTestType                = "networking-uxi/service-test"
	WiredNetworkType               = "networking-uxi/wired-network"
	WirelessNetworkType            = "networking-uxi/wireless-network"

	AgentGroupAssignmentPath       = "/networking-uxi/v1alpha1/agent-group-assignments"
	AgentPath                      = "/networking-uxi/v1alpha1/agents"
	GroupPath                      = "/networking-uxi/v1alpha1/groups"
	NetworkGroupAssignmentPath     = "/networking-uxi/v1alpha1/network-group-assignments"
	SensorGroupAssignmentPath      = "/networking-uxi/v1alpha1/sensor-group-assignments"
	SensorPath                     = "/networking-uxi/v1alpha1/sensors"
	ServiceTestGroupAssignmentPath = "/networking-uxi/v1alpha1/service-test-group-assignments"
	ServiceTestPath                = "/networking-uxi/v1alpha1/service-tests"
	WiredNetworkPath               = "/networking-uxi/v1alpha1/wired-networks"
	WirelessNetworkPath            = "/networking-uxi/v1alpha1/wireless-networks"
)

func TestOptionalValue(
	t *testing.T,
	tfResource string,
	tfKey string,
	property *string,
) resource.TestCheckFunc {
	if property == nil {
		return resource.TestCheckNoResourceAttr(tfResource, tfKey)
	}

	return resource.TestCheckResourceAttrPtr(tfResource, tfKey, property)
}
