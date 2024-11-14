package util

import (
	"context"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/nbio/st"
)

func GetSensorGroupAssignment(id string) config_api_client.SensorGroupAssignmentsItem {
	result, _, err := Client.ConfigurationAPI.
		SensorGroupAssignmentsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("sensor_group_assignment with id `" + id + "` could not be found")
	}
	return result.Items[0]
}

func CheckStateAgainstSensorGroupAssignment(
	t st.Fatalf,
	entity string,
	sensorGroupAssignment config_api_client.SensorGroupAssignmentsItem,
) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(entity, "id", sensorGroupAssignment.Id),
		resource.TestCheckResourceAttr(entity, "group_id", sensorGroupAssignment.Group.Id),
		resource.TestCheckResourceAttr(entity, "sensor_id", sensorGroupAssignment.Sensor.Id),
	)
}
