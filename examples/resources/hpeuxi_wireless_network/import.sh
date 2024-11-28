# Import a wireless network using its ID
terraform import hpeuxi_wireless_network.my_wireless_network <my_wireless_network_id>

# Import a wireless network using its ID with an import block
import {
    to = hpeuxi_wireless_network.my_wireless_network
    id = "<my_wireless_network_id>"
}
