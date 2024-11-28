# Retrieve data for an agent group assignment
data "hpeuxi_agent_group_assignment" "my_agent_group_assignment" {
  filter = {
    id = "<my_agent_group_assignment_id>"
  }
}
