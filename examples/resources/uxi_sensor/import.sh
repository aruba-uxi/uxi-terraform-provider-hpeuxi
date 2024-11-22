# Import sensor using its ID
terraform import uxi_sensor.my_sensor <my_sensor_id>

# Import sensor using its ID with an import block
import {
    to = uxi_sensor.my_sensor
    id = "uid"
}
