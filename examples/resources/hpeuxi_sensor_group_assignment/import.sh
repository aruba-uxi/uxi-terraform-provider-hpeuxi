# Import sensor group assignment using its ID
terraform import hpeuxi_sensor_group_assignment.my_sensor_group_assignment <my_sensor_group_assignment_id>

# Import sensor group assignment using its ID with an import block
import {
    to = hpeuxi_sensor_group_assignment.my_sensor_group_assignment
    id = "<my_sensor_group_assignment_id>"
}
