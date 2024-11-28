resource "hpeuxi_wireless_network" "my_wireless_network" {
  name         = "name"

  # Deleting of wireless networks is not supported yet
  lifecycle {
    prevent_destroy = true
  }
}
