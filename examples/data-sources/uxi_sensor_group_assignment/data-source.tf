# Retrieve data for a sensor group assignment
data "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
  filter = {
    id = "<my_sensor_group_assignment_id>"
  }
}
