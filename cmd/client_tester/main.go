package main

import (
	"aruba-uxi/configration_api_client"
	"fmt"
)

func main() {
	client := configuration_api_client.UxiConfigurationApiClient{
		Token: "Token",
		Host:  "https://www.google.com",
	}
	sensors, err := client.GetSensors()

	fmt.Println(sensors)
	if err != nil {
		panic(err)
	}
}
