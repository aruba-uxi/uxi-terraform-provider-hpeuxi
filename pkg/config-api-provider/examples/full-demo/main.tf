terraform {
  required_providers {
    uxi = {
      source = "registry.terraform.io/arubauxi/configuration"
    }
  }
}

provider "uxi" {}

# resource "uxi_group" "group" {
#   name       = "test_name"
#   parent_uid = "9999"
# }

// to import: terraform import uxi_sensor.my_sensor my_sensor_uid
// to remove: terraform state rm uxi_sensor.my_sensor
resource "uxi_sensor" "my_sensor" {
  name         = "name"
  address_note = "address_note"
  notes        = "notes"
  pcap_mode    = "pcap_mode"
}

// to import: terraform import uxi_agent.my_agent my_agent_uid
// to remove: terraform state rm uxi_agent.my_agent
resource "uxi_agent" "my_agent" {
  name         = "name"
  notes        = "notes"
  pcap_mode    = "pcap_mode"
}

# output "group" {
#   value = uxi_group.group
# }
