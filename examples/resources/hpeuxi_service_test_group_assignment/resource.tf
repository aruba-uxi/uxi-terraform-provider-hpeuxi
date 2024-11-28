resource "hpeuxi_service_test_group_assignment" "my_service_test_group_assignment" {
    service_test_id = hpeuxi_service_test.my_service_test.id
    group_id = hpeuxi_group.my_group.id
}
