package key_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListApplicationKeysCmd = &cobra.Command{
	Use:   "list-application-keys",
	
	Short: "Get all application keys",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.ListApplicationKeys(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-application-keys: %v", err)
		}

		cmdutil.PrintJSON(res, "application_keys")
	},
}

func init() {
	Cmd.AddCommand(ListApplicationKeysCmd)
}
