// NOTE: This module is just here for demonstration purposes

package main

import (
	"context"
	"fmt"

	"github.com/aruba-uxi/configuration-api-terraform-provider/pkg/config-api-client"
)

func main() {
	config := config_api_client.NewConfiguration()
	config.Host = "localhost:80"
	config.Scheme = "http"
	client := config_api_client.NewAPIClient(config)

	resp, httpResp, err := client.HealthAPI.GetLivezHealthLivezGet(context.Background()).Execute()

	fmt.Println(resp)
	fmt.Println(httpResp)
	fmt.Println(err)
}
