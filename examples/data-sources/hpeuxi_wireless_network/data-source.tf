# Retrieve data for a wireless network
data "hpeuxi_wireless_network" "my_wireless_network" {
  filter = {
    id = "<my_wireless_network_id>"
  }
}
