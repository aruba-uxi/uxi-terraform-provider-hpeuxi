terraform {
  required_providers {
    uxi = {
      source = "registry.terraform.io/arubauxi/hpeuxi"
    }
  }
}

provider "uxi" {
    client_id     = "client_id"
    client_secret = "some_random_secret"
}


resource "uxi_group" "hq" {
  name            = "HQ"
}

resource "uxi_group" "building_a" {
  name            = "Building A"
  parent_group_id = uxi_group.hq.id
}

resource "uxi_group" "office_a_1" {
  name            = "Office A1"
  parent_group_id = uxi_group.building_a.id
}

resource "uxi_group" "office_a_2" {
  name            = "Office A2"
  parent_group_id = uxi_group.building_a.id
}

resource "uxi_group" "building_b" {
  name            = "Building B"
  parent_group_id = uxi_group.hq.id
}

resource "uxi_group" "office_b_1" {
  name            = "Office B1"
  parent_group_id = uxi_group.building_b.id
}

resource "uxi_group" "office_b_2" {
  name            = "Office B2"
  parent_group_id = uxi_group.building_b.id
}
