// This file describes a potential layout for a manifest defined by the customer
// This file will be used as reference when building the terraform datasources and resources

# Note:
# general data source | some sort of root_group_query
# investigate implementation of import block -> conflict of user defined fields
# assessment to see how difficult it would be for a user get the uid of various resources
# revisit the flow of a user onboarding existing resources -> moving from data source to resource

data "uxi_group" "root_group" {
  filter {
    some_sort_of_filter = ""
  }
}

resource "uxi_group" "group" {
  name = "group_name"                    # required, mutable
  parent_uid = data.group.root_group.uid # required, immutable -> recreate on update
}

# Sensor
resource "uxi_sensor" "sensor" {
  name = "name"                  # mutable
  address_note = "address_note"  # mutable
  notes = "note"                 # mutable
  pcap_mode = "light"            # mutable
}

resource "uxi_sensor_group_assignment" "sensor_group_assignment" {
  sensor_id = uxi_sensor.sensor.id  # required, immutable -> recreate on update
  group_id = uxi_group.group.id     # required, immutable -> recreate on update
}

# Agent
resource "uxi_agent" "agent_sensor" {
  name = "name"                 # mutable
  address_note = "address_note" # mutable
  notes = "note"                # mutable
  pcap_mode = "light"           # mutable
}

resource "uxi_agent_group_assignment" "agent_group_assignment" {
  agent_id = uxi_agent.agent_sensor.id  # required, immutable -> recreate on update
  group_id = uxi_group.group.id         # required, immutable -> recreate on update
}

# wireless networks
data "uxi_wireless_network" "wireless_network" {
  uid = ""
}

resource "uxi_network_group_assignment" "wireless_network_group_assignment" {
  network_id = uxi_wireless_network.wireless_network.id  # required, immutable -> recreate on update
  group_id = uxi_group.group.id                          # required, immutable -> recreate on update
}

# wired networks
data "uxi_wired_network" "wired_network" {
  uid = ""
}

resource "uxi_network_group_assignment" "wired_network_group_assignment" {
  network_id = uxi_wired_network.wired_network.id # required, immutable -> recreate on update
  group_id = uxi_group.group.id                   # required, immutable -> recreate on update
}

# service-test
data "uxi_service_test" "service_test" {
  uid = ""
}

resource "uxi_service_test_group_assignment" "service_test_group_assignment" {
  service_test_id = uxi_service_test.service_test.id  # required, immutable -> recreate on update
  group_id = uxi_group.group.id                       # required, immutable -> recreate on update
}


# ----------------------------------
import {
  to = "asfasdf"
  id = uid
}

resource "uxi_sensor_group_assignment" ""

# terraform refresh-state
