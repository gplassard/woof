package key_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCurrentUserApplicationKeysCmd = &cobra.Command{
	Use:   "list-current-user-application-keys",
	
	Short: "Get all application keys owned by current user",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.ListCurrentUserApplicationKeys(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-current-user-application-keys: %v", err)
		}

		cmdutil.PrintJSON(res, "application_keys")
	},
}

func init() {
	Cmd.AddCommand(ListCurrentUserApplicationKeysCmd)
}
