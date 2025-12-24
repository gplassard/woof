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

var listUserOrganizations = &cobra.Command{
	Use:   "organizations [user_id]",
	Short: "Get a user organization",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userID := args[0]

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

		res, _, err := usersApi.ListUserOrganizations(ctx, userID)
		if err != nil {
			log.Fatalf("failed to list user organizations: %v", err)
		}

		// Filter data array to only include items with "type": "users"
		// and keep relationships attributes
		b, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("failed to marshal user organizations for filtering: %v", err)
		}
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			log.Fatalf("failed to unmarshal user organizations for filtering: %v", err)
		}

		if data, ok := m["data"].([]interface{}); ok {
			var filteredData []interface{}
			for _, item := range data {
				if obj, ok := item.(map[string]interface{}); ok {
					if t, ok := obj["type"].(string); ok && t == "users" {
						filteredData = append(filteredData, obj)
					}
				}
			}
			m["data"] = filteredData
		}

		// Remove included array if it exists as it may contain non-user objects
		delete(m, "included")

		s, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			log.Fatalf("failed to marshal user organizations: %v", err)
		}
		fmt.Println(string(s))
	},
}
