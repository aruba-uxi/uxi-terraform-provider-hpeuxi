terraform {
  required_providers {
    uxi = {
      source = "registry.terraform.io/arubauxi/configuration"
    }
  }
}

provider "uxi" {}

resource "uxi_group" "my_group" {
  name       = "name"
  parent_uid = "parent_uid"
}

// Sensor Resource
/*
To import:
import {
    to = uxi_sensor.my_sensor
    id = "uid"
}

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

// Agent Resource
/*
To import:
import {
    to = uxi_agent.my_agent
    id = "uid"
}

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

// Wireless Network Resource
/*
To import:
import {
    to = uxi_wireless_network.my_wireless_network
    id = "uid"
}

To remove:
removed {
    from = uxi_wireless_network.my_wireless_network

    lifecycle {
        destroy = false
    }
}
*/
resource "uxi_wireless_network" "my_wireless_network" {
    alias = "alias"
}

// Wired Network Resource
/*
To import:
import {
    to = uxi_wired_network.my_wired_network
    id = "uid"
}

To remove:
removed {
    from = uxi_wired_network.my_wired_network

    lifecycle {
        destroy = false
    }
}
*/
resource "uxi_wired_network" "my_wired_network" {
    alias = "alias"
}

// Service Test Resource
/*
To import:
import {
    to = uxi_service_test.my_service_test
    id = "uid"
}

To remove:
removed {
    from = uxi_service_test.my_service_test

    lifecycle {
        destroy = false
    }
}
*/
resource "uxi_service_test" "my_service_test" {
    title = "title"
}


# output "group" {
#   value = uxi_group.group
# }
