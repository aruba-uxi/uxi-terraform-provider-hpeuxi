# Import an agent group assignment using its ID
terraform import hpeuxi_agent_group_assignment.my_agent_group_assignment <my_agent_group_assignment_id>

# Import an agent group assignment using its ID with an import block
import {
    to = hpeuxi_agent_group_assignment.my_agent_group_assignment
    id = "<my_agent_group_assignment_id>"
}
