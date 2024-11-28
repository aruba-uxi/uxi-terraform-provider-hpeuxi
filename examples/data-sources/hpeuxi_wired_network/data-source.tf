# Retrieve data for a wired network
data "hpeuxi_wired_network" "my_wired_network" {
  filter = {
    id = "<my_wired_network_id>"
  }
}
