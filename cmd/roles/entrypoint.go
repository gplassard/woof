package roles

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

var listRoles = &cobra.Command{
	Use:   "list",
	Short: "List Datadog roles",
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
		rolesApi := datadogV2.NewRolesApi(apiClient)

		res, _, err := rolesApi.ListRoles(ctx)
		if err != nil {
			log.Fatalf("failed to list roles: %v", err)
		}

		// Filter data array to only include items with "type": "user"
		// and keep relationships attributes
		b, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("failed to marshal roles for filtering: %v", err)
		}
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			log.Fatalf("failed to unmarshal roles for filtering: %v", err)
		}

		if data, ok := m["data"].([]interface{}); ok {
			var filteredData []interface{}
			for _, item := range data {
				if user, ok := item.(map[string]interface{}); ok {
					if t, ok := user["type"].(string); ok && t == "roles" {
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
			log.Fatalf("failed to marshal roles: %v", err)
		}
		fmt.Println(string(s))
	},
}
