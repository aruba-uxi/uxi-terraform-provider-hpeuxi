/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package shared

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func CheckStateAgainstServiceTest(
	t *testing.T,
	entity string,
	serviceTest config_api_client.ServiceTestsGetItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", serviceTest.Id),
		resource.TestCheckResourceAttr(entity, "category", serviceTest.Category),
		resource.TestCheckResourceAttr(entity, "name", serviceTest.Name),
		TestOptionalValue(t, entity, "target", serviceTest.Target.Get()),
		resource.TestCheckResourceAttr(entity, "template", serviceTest.Template),
		resource.TestCheckResourceAttr(
			entity,
			"is_enabled",
			strconv.FormatBool(serviceTest.IsEnabled),
		),
	)
}
