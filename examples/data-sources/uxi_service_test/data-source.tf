# Retrieve data for a service test
data "uxi_service_test" "my_service_test" {
  filter = {
    id = "<my_service_test_id>"
  }
}
