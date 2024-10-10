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

type ResponseStatusCheck struct {
	Response *http.Response
	Err      error
}

func (r ResponseStatusCheck) RaiseForStatus() (bool, string) {
	if r.Err != nil {
		var detail string
		var data map[string]interface{}

		switch e := r.Err.(type) {
		case *url.Error:
			detail = r.handleURLError(e)
		default:
			if err := json.NewDecoder(r.Response.Body).Decode(&data); err != nil {
				detail = "Unexpected error: " + r.Err.Error()
			} else {
				detail = data["message"].(string) + "\nDebugID: " + data["debugId"].(string)
			}
		}

		return true, detail
	}
	return false, ""
}

func (r ResponseStatusCheck) handleURLError(uErr *url.Error) string {
	if uErr.Timeout() {
		return "Error: Request timed out. Please check your network."
	} else if uErr.Temporary() {
		return "Error: Temporary network error. Please try again later."
	} else {
		return fmt.Sprintf("URL Error: %v\n", uErr)
	}
}
