# Import a service test using its ID
terraform import uxi_service_test.my_service_test <my_service_test_id>

# Import a service test using its ID with an import block
import {
    to = uxi_service_test.my_service_test
    id = "<my_service_test_id>"
}
