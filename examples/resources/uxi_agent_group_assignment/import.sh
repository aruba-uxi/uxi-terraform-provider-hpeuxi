# Import an agent group assignment using its ID
terraform import uxi_agent_group_assignment.my_agent_group_assignment <my_agent_group_assignment_id>

# Import an agent group assignment using its ID with an import block
import {
    to = uxi_agent_group_assignment.my_agent_group_assignment
    id = "<my_agent_group_assignment_id>"
}
