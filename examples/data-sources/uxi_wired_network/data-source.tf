# Retrieve data for a wired network
data "uxi_wired_network" "my_wired_network" {
  filter = {
    id = "<my_wired_network_id>"
  }
}
