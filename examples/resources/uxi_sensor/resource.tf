resource "uxi_sensor" "my_sensor" {
  name         = "name"
  address_note = "address_note"
  notes        = "notes"
  pcap_mode    = "pcap_mode"

  lifecycle {
    prevent_destroy = true
  }
}
