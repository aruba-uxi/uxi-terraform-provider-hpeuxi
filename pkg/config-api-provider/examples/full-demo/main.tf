terraform {
  required_providers {
    uxi = {
      source = "registry.terraform.io/arubauxi/configuration"
    }
  }
}

provider "uxi" {}

resource "uxi_group" "group" {
  name       = "test_name"
  parent_uid = "9999"
}

output "group" {
  value = uxi_group.group
}
