package users

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/spf13/cobra"
)

var listUsers = &cobra.Command{
	Use:   "list",
	Short: "List Datadog users",
	Run: func(cmd *cobra.Command, args []string) {
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

		// Filter data array to only include items with "type": "user"
		// and keep relationships attributes
		b, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("failed to marshal users for filtering: %v", err)
		}
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			log.Fatalf("failed to unmarshal users for filtering: %v", err)
		}

		if data, ok := m["data"].([]interface{}); ok {
			var filteredData []interface{}
			for _, item := range data {
				if user, ok := item.(map[string]interface{}); ok {
					if t, ok := user["type"].(string); ok && t == "users" {
						filteredData = append(filteredData, user)
					}
				}
			}
			m["data"] = filteredData
		}

		// Remove included array if it exists as it may contain non-user objects
		delete(m, "included")

		s, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			log.Fatalf("failed to marshal users: %v", err)
		}
		fmt.Println(string(s))
	},
}
