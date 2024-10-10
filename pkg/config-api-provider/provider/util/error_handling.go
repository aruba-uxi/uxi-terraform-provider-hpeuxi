package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func GenerateErrorSummary(actionName string, entityName string) string {
	return fmt.Sprintf("Error performing %s on %s", actionName, entityName)
}

func RaiseForStatus(response *http.Response, err error) (bool, string) {
	if err != nil {
		var detail string
		var data map[string]interface{}

		switch e := err.(type) {
		case *url.Error:
			detail = handleURLError(e)
		default:
			if jsonDecodeErr := json.NewDecoder(response.Body).Decode(&data); jsonDecodeErr != nil {
				detail = "Unexpected error: " + jsonDecodeErr.Error()
			} else {
				detail = data["message"].(string) + "\nDebugID: " + data["debugId"].(string)
			}
		}

		return true, detail
	}
	return false, ""
}

func handleURLError(uErr *url.Error) string {
	if uErr.Timeout() {
		return "Error: Request timed out. Please check your network."
	} else if uErr.Temporary() {
		return "Error: Temporary network error. Please try again later."
	} else {
		return fmt.Sprintf("URL Error: %v\n", uErr)
	}
}
