resource "uxi_sensor" "my_sensor" {
  name         = "name"
  address_note = "Example 3rd Floor"
  notes        = "notes"
  pcap_mode    = "light"

  # Deleting a sensor is not supported
  lifecycle {
    prevent_destroy = true
  }
}
