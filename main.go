package main

import (
	"context"
	"log"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func main() {
	log.Println("Starting datadog-api-client")
	ctx := context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_APP_KEY"),
			},
		},
	)
	ctx = context.WithValue(ctx,
		datadog.ContextServerVariables,
		map[string]string{
			"site": "datadoghq.eu",
		})

	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	usersApi := datadogV2.NewUsersApi(apiClient)

	res, _, err := usersApi.ListUsers(ctx)
	if err != nil {
		log.Fatalf("failed to list users: %v", err)
	}
	log.Printf("users: %v", res)
}
