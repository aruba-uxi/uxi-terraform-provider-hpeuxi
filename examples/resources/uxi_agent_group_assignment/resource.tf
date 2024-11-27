resource "uxi_agent_group_assignment" "my_agent_group_assignment" {
    agent_id = uxi_agent.my_agent.id
    group_id = uxi_group.my_group.id
}
