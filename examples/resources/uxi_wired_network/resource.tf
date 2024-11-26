resource "uxi_wired_network" "my_wired_network" {
  name         = "name"

  # Deleting of wired networks is not supported yet
  lifecycle {
    prevent_destroy = true
  }
}
