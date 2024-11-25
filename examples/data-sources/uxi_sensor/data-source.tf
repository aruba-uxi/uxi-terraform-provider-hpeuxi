# Retrieve data for a sensor
data "uxi_sensor" "my_sensor" {
  filter = {
    id = "<my_sensor_id>"
  }
}
