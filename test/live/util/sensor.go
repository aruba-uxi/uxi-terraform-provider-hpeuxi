package util

import (
	"context"
)

type sensorProperties struct {
	Id           string
	Name         string
	Notes        *string
	AddressNotes *string
	PcapMode     *string
}

func GetSensorProperties(id string) sensorProperties {
	result, _, err := Client.ConfigurationAPI.
		SensorsGet(context.Background()).
		Id(id).
		Execute()
	if err != nil {
		panic(err)
	}
	if len(result.Items) != 1 {
		panic("sensor with id `" + id + "` could not be found")
	}
	sensor := result.Items[0]
	// Read these in, as they may not be always constant with the acceptance test
	// customer
	return sensorProperties{
		Id:           sensor.Id,
		Name:         sensor.Name,
		Notes:        sensor.Notes.Get(),
		AddressNotes: sensor.AddressNote.Get(),
		PcapMode:     sensor.PcapMode.Get(),
	}
}
