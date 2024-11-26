resource "uxi_service_test" "my_service_test" {
  name         = "name"

  # Deleting of service tests is not supported
  lifecycle {
    prevent_destroy = true
  }
}
