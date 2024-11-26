resource "uxi_sensor_group_assignment" "my_sensor_group_assignment" {
    sensor_id = uxi_sensor.my_sensor.id
    group_id = uxi_group.my_group.id
}
