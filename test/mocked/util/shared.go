/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package util

import (
	"net/http"

	"github.com/h2non/gock"
)

const (
	mockToken  = "mock_token"
	MockDomain = "test.api.capenetworks.com"
	MockUxiUrl = "https://" + MockDomain
)

var RateLimitingHeaders = map[string]string{
	"X-RateLimit-Limit":     "100",
	"X-RateLimit-Remaining": "0",
	"X-RateLimit-Reset":     "0.01",
}

func GeneratePaginatedResponse(items []map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"items": items,
		"next":  nil,
		"count": len(items),
	}
}

func MockOAuth() *gock.Response {
	return gock.New("https://sso.common.cloud.hpe.com").
		Post("/as/token.oauth2").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		Persist().
		Reply(http.StatusOK).
		JSON(map[string]interface{}{
			"access_token": mockToken,
			"token_type":   "bearer",
			"expires_in":   3600,
		})
}
