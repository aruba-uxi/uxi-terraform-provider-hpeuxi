resource "hpeuxi_agent_group_assignment" "my_agent_group_assignment" {
  agent_id = hpeuxi_agent.my_agent.id
  group_id = hpeuxi_group.my_group.id
}
