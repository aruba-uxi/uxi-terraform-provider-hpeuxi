# Retrieve data for an agent
data "hpeuxi_agent" "my_agent" {
  filter = {
    id = "<my_agent_id>"
  }
}
