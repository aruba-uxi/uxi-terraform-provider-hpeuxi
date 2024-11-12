package util

import (
	"context"
	"strconv"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/aruba-uxi/terraform-provider-hpeuxi/test/live/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
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
	t st.Fatalf,
	serviceTest config_api_client.ServiceTestsListItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(
			"data.uxi_service_test.my_service_test",
			"id",
			config.ServiceTestUid,
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
