/*
Copyright 2024 Hewlett Packard Enterprise Development LP.
*/

package config

import "os"

const (
	MaxRetriesForTooManyRequests = 10
	TokenURL                     = "https://sso.common.cloud.hpe.com/as/token.oauth2"
	UXIDefaultHost               = "api.capenetworks.com"
)

var (
	Host         string
	ClientID     string
	ClientSecret string
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func InitializeConfig() {
	Host = getEnv("UXI_HOST", UXIDefaultHost)
	ClientID = os.Getenv("GREENLAKE_UXI_CLIENT_ID")
	ClientSecret = os.Getenv("GREENLAKE_UXI_CLIENT_SECRET")
}
