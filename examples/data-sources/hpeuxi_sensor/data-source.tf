# Retrieve data for a sensor
data "hpeuxi_sensor" "my_sensor" {
  filter = {
    id = "<my_sensor_id>"
  }
}
