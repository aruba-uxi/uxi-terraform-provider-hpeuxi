# Import agent using its ID
terraform import hpeuxi_agent.my_agent <my_agent_id>

# Import agent using its ID with an import block
import {
    to = hpeuxi_agent.my_agent
    id = "<my_agent_id>"
}
