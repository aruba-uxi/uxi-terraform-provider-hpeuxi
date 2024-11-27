resource "uxi_network_group_assignment" "my_network_group_assignment" {
    network_id = uxi_wired_network.my_network.id
    group_id = uxi_group.my_group.id
}
