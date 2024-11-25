/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package config

// These constants are from the customer:
// Configuration-API Acceptance Testing (844457745a1111ef880836000a52e73e)
// And therefore the client_id and client_secret used for the acceptance tests must match this
// customer.

// This agent is permanently on the customer
const AgentPermanentId = "abeac07d-3c28-31be-b9e7-30002c753ed4"

// An agent with this serial will get provisioned
const AgentCreateSerial = "56fb38331f19d278"

// The customer ID of the acceptance test customer
const CustomerId = "9c16d493-7649-40bc-975b-07422d227c0b"

// The root group ID
const GroupIdRoot = "07422d227c0b"

// This wired network is permanently on the customer
const (
	WiredNetworkId   = "ethernet-0ee5b46c2ef0"
	WiredNetworkName = "tf-provider-acceptance-tests-ethernet-0"
)

// This wireless network is permanently on the customer
const (
	WirelessNetworkId   = "ssid-bf704ff37dc0"
	WirelessNetworkName = "tf-provider-acceptance-tests-ssid-0"
)

// This service test is permanently on the customer
const (
	ServiceTestId   = "6f81e43d-76f1-4a15-aafe-4ce2371d918a"
	ServiceTestName = "tf-provider-acceptance-test-0"
)

// This sensor is permanently on the customer
const SensorId = "4b031caf-cea8-411d-8928-79f518163dae"

const DeviceGatewayHost = "https://device-gateway.staging.capedev.io"
