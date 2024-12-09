/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

func ConvertToStrings[T ~string](input []T) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = string(v)
	}

	return result
}
