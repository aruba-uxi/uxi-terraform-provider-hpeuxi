# Import network group assignment using its ID
terraform import hpeuxi_network_group_assignment.my_network_group_assignment <my_network_group_assignment_id>

# Import network group assignment using its ID with an import block
import {
    to = hpeuxi_network_group_assignment.my_network_group_assignment
    id = "<my_network_group_assignment_id>"
}
