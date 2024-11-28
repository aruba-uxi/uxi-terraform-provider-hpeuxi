resource "hpeuxi_sensor_group_assignment" "my_sensor_group_assignment" {
    sensor_id = hpeuxi_sensor.my_sensor.id
    group_id = hpeuxi_group.my_group.id
}
