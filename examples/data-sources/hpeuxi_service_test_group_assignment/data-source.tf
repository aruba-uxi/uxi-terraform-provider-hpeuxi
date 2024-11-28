# Retrieve data for a service test group assignment
data "hpeuxi_service_test_group_assignment" "my_service_test_group_assignment" {
  filter = {
    id = "<my_service_test_group_assignment_id>"
  }
}
