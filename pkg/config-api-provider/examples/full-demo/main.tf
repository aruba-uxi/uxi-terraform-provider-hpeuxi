terraform {
  required_providers {
    uxi = {
      source = "registry.terraform.io/arubauxi/configuration"
    }
  }
}

provider "uxi" {
    host		  = "test.api.capenetworks.com"
    client_id     = "client_id"
    client_secret = "some_random_secret"
    token_url     = "https://test.sso.common.cloud.hpe.com/as/token.oauth2"
}

resource "uxi_group" "my_group" {
  name            = "parent"
}

resource "uxi_group" "my_group_2" {
  name            = "child"
  parent_group_id = uxi_group.my_group.id
}

// Sensor Resource
/*
To remove:
removed {
    from = uxi_sensor.my_sensor

    lifecycle {
        destroy = false
    }
}
*/

resource "uxi_sensor" "my_sensor" {
  name         = "name"
  address_note = "address_note"
  notes        = "notes"
  pcap_mode    = "pcap_mode"
}

import {
    to = uxi_sensor.my_sensor
    id = "uid"
}


// Agent Resource
/*
To remove:
removed {
    from = uxi_agent.my_agent

    lifecycle {
        destroy = false
    }
}
*/
resource "uxi_agent" "my_agent" {
  name         = "name"
  notes        = "notes"
  pcap_mode    = "pcap_mode"
}

import {
    to = uxi_agent.my_agent
    id = "uid"
}


// Wireless Network Resource
/*
To remove:
removed {
    from = uxi_wireless_network.my_wireless_network

    lifecycle {
        destroy = false
    }
}
*/
resource "uxi_wireless_network" "my_wireless_network" {
    name = "name"
}

import {
    to = uxi_wireless_network.my_wireless_network
    id = "uid"
}

// Wired Network Resource
/*
To remove:
removed {
    from = uxi_wired_network.my_wired_network

    lifecycle {
        destroy = false
    }
}
*/
resource "uxi_wired_network" "my_wired_network" {
    name = "name"
}

import {
    to = uxi_wired_network.my_wired_network
    id = "uid"
}
// Service Test Resource
/*
To remove:
removed {
    from = uxi_service_test.my_service_test

    lifecycle {
        destroy = false
    }
}
*/
resource "uxi_service_test" "my_service_test" {
    name = "name"
}

import {
    to = uxi_service_test.my_service_test
    id = "uid"
}

// Sensor Group Assignment
resource "uxi_sensor_group_assignment" "my_uxi_sensor_group_assignment" {
    sensor_id = uxi_sensor.my_sensor.id
    group_id = uxi_group.my_group.id
}

// Agent Group Assignment
resource "uxi_agent_group_assignment" "my_uxi_agent_group_assignment" {
    agent_id = uxi_agent.my_agent.id
    group_id = uxi_group.my_group.id
}

// Network Group Assignment
resource "uxi_network_group_assignment" "my_uxi_network_group_assignment" {
    network_id = uxi_agent.my_agent.id
    group_id = uxi_group.my_group.id
}

# output "group" {
#   value = uxi_group.group
# }
