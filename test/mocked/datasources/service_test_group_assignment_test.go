package data_source_test

import (
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/nbio/st"
)

func TestServiceTestGroupAssignmentDataSource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					util.MockGetServiceTestGroupAssignment(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestGroupAssignmentResponse("id", ""),
							},
						),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						filter = {
							service_test_group_assignment_id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.uxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						"id",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_service_test_group_assignment.my_service_test_group_assignment",
						"group_id",
						"group_id",
					),
					resource.TestCheckResourceAttr(
						"data.uxi_service_test_group_assignment.my_service_test_group_assignment",
						"service_test_id",
						"service_test_id",
					),
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestServiceTestGroupAssignmentDataSource429Handling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mock429 *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				PreConfig: func() {
					mock429 = gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/service-test-group-assignments").
						Reply(429).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetServiceTestGroupAssignment(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestGroupAssignmentResponse("id", ""),
							},
						),
						3,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						filter = {
							service_test_group_assignment_id = "id"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.uxi_service_test_group_assignment.my_service_test_group_assignment",
						"id",
						"id",
					),
					func(s *terraform.State) error {
						st.Assert(t, mock429.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
func TestServiceTestGroupAssignmentDataSourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/service-test-group-assignments").
						Reply(500).
						JSON(map[string]interface{}{
							"httpStatusCode": 500,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					data "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						filter = {
							service_test_group_assignment_id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
			{
				PreConfig: func() {
					util.MockGetServiceTestGroupAssignment(
						"id",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					data "uxi_service_test_group_assignment" "my_service_test_group_assignment" {
						filter = {
							service_test_group_assignment_id = "id"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Could not find specified data source`),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
