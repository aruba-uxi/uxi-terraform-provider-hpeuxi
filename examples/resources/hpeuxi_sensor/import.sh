# Import sensor using its ID
terraform import hpeuxi_sensor.my_sensor <my_sensor_id>

# Import sensor using its ID with an import block
import {
    to = hpeuxi_sensor.my_sensor
    id = "<my_sensor_id>"
}
