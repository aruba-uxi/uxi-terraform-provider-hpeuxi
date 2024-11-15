/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"context"
	"strconv"
	"testing"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func GetServiceTest(id string) config_api_client.ServiceTestsListItem {
	result, _, err := Client.ConfigurationAPI.
		ServiceTestsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("service_test with id `" + id + "` could not be found")
	}
	return result.Items[0]
}

func CheckStateAgainstServiceTest(
	t *testing.T,
	serviceTest config_api_client.ServiceTestsListItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(
			"data.uxi_service_test.my_service_test",
			"id",
			config.ServiceTestId,
		),
		resource.TestCheckResourceAttr(
			"data.uxi_service_test.my_service_test",
			"category",
			serviceTest.Category,
		),
		resource.TestCheckResourceAttr(
			"data.uxi_service_test.my_service_test",
			"name",
			serviceTest.Name,
		),
		TestOptionalValue(
			t,
			"data.uxi_service_test.my_service_test",
			"target",
			serviceTest.Target.Get(),
		),
		resource.TestCheckResourceAttr(
			"data.uxi_service_test.my_service_test",
			"template",
			serviceTest.Template,
		),
		resource.TestCheckResourceAttr(
			"data.uxi_service_test.my_service_test",
			"is_enabled",
			strconv.FormatBool(serviceTest.IsEnabled),
		),
	)
}
