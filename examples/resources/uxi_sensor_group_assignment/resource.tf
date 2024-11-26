terraform {
  required_providers {
    uxi = {
      source = "registry.terraform.io/arubauxi/hpeuxi"
    }
  }
}

provider "uxi" {

}

resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
  sensor_id = "<my_sensor_id>"
  group_id = "<my_group_id>"
}
