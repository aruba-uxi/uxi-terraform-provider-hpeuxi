resource "hpeuxi_network_group_assignment" "my_network_group_assignment" {
    network_id = hpeuxi_wired_network.my_network.id
    group_id = hpeuxi_group.my_group.id
}
