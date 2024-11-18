package resource_test

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/provider"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/mocked/util"
	"github.com/stretchr/testify/assert"

	"github.com/h2non/gock"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestServiceTestResource(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Creating a service_test is not allowed
			{
				Config: provider.ProviderConfig + `
					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}`,

				ExpectError: regexp.MustCompile(
					`(?s)creating a service_test is not supported; service_tests can only be\s*imported`,
				),
			},
			// Importing a service_test
			{
				PreConfig: func() {
					util.MockGetServiceTest(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("id", ""),
							},
						),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"uxi_service_test.my_service_test",
						"name",
						"name",
					),
					resource.TestCheckResourceAttr("uxi_service_test.my_service_test", "id", "id"),
				),
			},
			// ImportState testing
			{
				PreConfig: func() {
					util.MockGetServiceTest(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("id", ""),
							},
						),
						1,
					)
				},
				ResourceName:      "uxi_service_test.my_service_test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Updating a service_test is not allowed
			{
				PreConfig: func() {
					util.MockGetServiceTest(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("id", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + `
				resource "uxi_service_test" "my_service_test" {
					name = "updated_name"
				}`,
				ExpectError: regexp.MustCompile(
					`(?s)updating a service_test is not supported; service_tests can only be updated\s*through the dashboard`,
				),
			},
			// Deleting a service_test is not allowed
			{
				PreConfig: func() {
					util.MockGetServiceTest(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("id", ""),
							},
						),
						1,
					)
				},
				Config: provider.ProviderConfig + ``,
				ExpectError: regexp.MustCompile(
					`(?s)deleting a service_test is not supported; service_tests can only removed from\s*state`,
				),
			},
			// Remove service_test from state
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_service_test.my_service_test

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestServiceTestResourceTooManyRequestsHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()
	var mockTooManyRequests *gock.Response

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Read
			{
				PreConfig: func() {
					mockTooManyRequests = gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/service-tests").
						Reply(http.StatusTooManyRequests).
						SetHeaders(util.RateLimitingHeaders)
					util.MockGetServiceTest(
						"id",
						util.GeneratePaginatedResponse(
							[]map[string]interface{}{
								util.GenerateServiceTestResponse("id", ""),
							},
						),
						2,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "id"
					}`,

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("uxi_service_test.my_service_test", "id", "id"),
					func(s *terraform.State) error {
						assert.Equal(t, mockTooManyRequests.Mock.Request().Counter, 0)
						return nil
					},
				),
			},
			// Cleanup
			{
				Config: provider.ProviderConfig + `
					removed {
						from = uxi_service_test.my_service_test

						lifecycle {
							destroy = false
						}
					}`,
			},
		},
	})

	mockOAuth.Mock.Disable()
}

func TestServiceTestResourceHttpErrorHandling(t *testing.T) {
	defer gock.Off()
	mockOAuth := util.MockOAuth()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: provider.TestAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			// we required terraform 1.7.0 and above for the `removed` block
			tfversion.RequireAbove(tfversion.Version1_7_0),
		},
		Steps: []resource.TestStep{
			// Read not found
			{
				PreConfig: func() {
					util.MockGetServiceTest(
						"id",
						util.GeneratePaginatedResponse([]map[string]interface{}{}),
						1,
					)
				},
				Config: provider.ProviderConfig + `
					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "id"
					}`,
				ExpectError: regexp.MustCompile(`Error: Cannot import non-existent remote object`),
			},
			// Read HTTP error
			{
				PreConfig: func() {
					gock.New("https://test.api.capenetworks.com").
						Get("/networking-uxi/v1alpha1/service-tests").
						Reply(http.StatusInternalServerError).
						JSON(map[string]interface{}{
							"httpStatusCode": http.StatusInternalServerError,
							"errorCode":      "HPE_GL_ERROR_INTERNAL_SERVER_ERROR",
							"message":        "Current request cannot be processed due to unknown issue",
							"debugId":        "12312-123123-123123-1231212",
						})
				},
				Config: provider.ProviderConfig + `
					resource "uxi_service_test" "my_service_test" {
						name = "name"
					}

					import {
						to = uxi_service_test.my_service_test
						id = "id"
					}`,
				ExpectError: regexp.MustCompile(
					`(?s)Current request cannot be processed due to unknown issue\s*DebugID: 12312-123123-123123-1231212`,
				),
			},
		},
	})

	mockOAuth.Mock.Disable()
}
