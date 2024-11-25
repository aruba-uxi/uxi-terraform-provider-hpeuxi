# Retrieve data for an agent
data "uxi_agent" "my_agent" {
  filter = {
    id = "<my_agent_id>"
  }
}
