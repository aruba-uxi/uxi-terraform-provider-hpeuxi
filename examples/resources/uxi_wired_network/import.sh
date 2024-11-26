# Import a wired network using its ID
terraform import uxi_wired_network.my_wired_network <my_wired_network_id>

# Import a wired network using its ID with an import block
import {
    to = uxi_wired_network.my_wired_network
    id = "<my_wired_network_id>"
}
