/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	config_api_client "github.com/aruba-uxi/terraform-provider-hpeuxi/pkg/config-api-client"
)

func GenerateErrorSummary(actionName string, entityName string) string {
	return fmt.Sprintf("Error performing %s on %s", actionName, entityName)
}

func RaiseForStatus(response *http.Response, err error) (bool, string) {
	if err != nil {
		var detail string

		var uErr *url.Error
		var apiErr *config_api_client.GenericOpenAPIError

		switch {
		case errors.As(err, &uErr):
			detail = handleURLError(uErr)
		case errors.As(err, &apiErr):
			detail = handleJSONError(response, apiErr)
		default:
			detail = "Unexpected error: " + err.Error()
		}

		return true, detail
	}

	return false, ""
}

func handleJSONError(
	response *http.Response,
	apiErr *config_api_client.GenericOpenAPIError,
) string {
	data, err := parseJSONResponse(response)
	if err != nil {
		return err.Error()
	}

	message := buildJSONErrorMsg(data, apiErr)
	debugID := buildJSONDebugID(data)
	parts := []string{message, debugID}

	return strings.Join(parts, "\n")
}

func parseJSONResponse(response *http.Response) (map[string]any, error) {
	var data map[string]interface{}

	err := json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return map[string]any{}, fmt.Errorf(
			"Unexpected error: there was an error decoding the API response body for "+
				"%d status code response.",
			response.StatusCode,
		)
	}

	return data, nil
}

func buildJSONErrorMsg(data map[string]any, apiErr *config_api_client.GenericOpenAPIError) string {
	message, found := data["message"]
	if !found {
		return "Unexpected error: " + apiErr.Error()
	}

	messageStr, castOK := message.(string)
	if !castOK {
		return "Unexpected error: " + apiErr.Error()
	}

	return messageStr
}

func buildJSONDebugID(data map[string]any) string {
	debugID, found := data["debugId"]
	if !found {
		return ""
	}

	debugIDStr, castOK := debugID.(string)
	if !castOK {
		return ""
	}

	return "DebugID: " + debugIDStr
}

func handleURLError(uErr *url.Error) string {
	switch {
	case uErr.Timeout():
		return "Error: Request timed out. Please check your network."
	case uErr.Temporary():
		return "Error: Temporary network error. Please try again later."
	default:
		return fmt.Sprintf("URL Error: %v\n", uErr)
	}
}
