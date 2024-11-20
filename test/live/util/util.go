/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"strconv"
)

func ConditionalProperty(property string, value *string) string {
	if value == nil {
		return ""
	}
	return property + `= "` + *value + `"`
}

func Int32PtrToStringPtr(value *int32) *string {
	if value == nil {
		return nil
	}
	result := strconv.Itoa(int(*value))
	return &result
}
