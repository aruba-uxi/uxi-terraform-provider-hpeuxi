/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

func ConditionalProperty(property string, value *string) string {
	if value == nil {
		return ""
	}
	return property + `= "` + *value + `"`
}
