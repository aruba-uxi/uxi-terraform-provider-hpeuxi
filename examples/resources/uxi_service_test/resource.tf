terraform {
  required_providers {
    uxi = {
      source = "registry.terraform.io/arubauxi/hpeuxi"
    }
  }
}

provider "uxi" {
}

resource "uxi_service_test" "my_service_test" {
  name         = "name"

  # Deleting of service tests is not supported
  lifecycle {
    prevent_destroy = true
  }
}
