# Retrieve data for a network group assignment
data "hpeuxi_network_group_assignment" "my_network_group_assignment" {
  filter = {
    id = "<my_network_group_assignment_id>"
  }
}
