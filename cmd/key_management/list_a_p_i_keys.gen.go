package key_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAPIKeysCmd = &cobra.Command{
	Use:   "list-a-p-i-keys",
	Short: "Get all API keys",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.ListAPIKeys(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-a-p-i-keys: %v", err)
		}

		cmdutil.PrintJSON(res, "api_keys")
	},
}

func init() {
	Cmd.AddCommand(ListAPIKeysCmd)
}
