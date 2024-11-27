resource "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
    service_test_id = uxi_service_test.my_service_test.id
    group_id = uxi_group.my_group.id
}
