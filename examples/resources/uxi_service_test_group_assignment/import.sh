# Import service test group assignment using its ID
terraform import uxi_service_test_group_assignment.my_service_test_group_assignment <my_service_test_group_assignment_id>

# Import service test group assignment using its ID with an import block
import {
    to = uxi_service_test_group_assignment.my_service_test_group_assignment
    id = "<my_service_test_group_assignment_id>"
}
