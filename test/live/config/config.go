/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package config

import (
	"os"
)

var (
	AgentID              = os.Getenv("ACCEPTANCE_AGENT_ID")
	AgentProvisionSerial = os.Getenv("ACCEPTANCE_AGENT_PROVISION_SERIAL")
	CustomerID           = os.Getenv("ACCEPTANCE_CUSTOMER_ID")
	GroupIDRoot          = os.Getenv("ACCEPTANCE_GROUP_ID_ROOT")
	WiredNetworkID       = os.Getenv("ACCEPTANCE_WIRED_NETWORK_ID")
	WirelessNetworkID    = os.Getenv("ACCEPTANCE_WIRELESS_NETWORK_ID")
	ServiceTestID        = os.Getenv("ACCEPTANCE_SERVICE_TEST_ID")
	SensorID             = os.Getenv("ACCEPTANCE_SENSOR_ID")
	DeviceGatewayHost    = os.Getenv("ACCEPTANCE_AGENT_PROVISION_HOST")
)
