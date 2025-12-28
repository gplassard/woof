package key_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCurrentUserApplicationKeyCmd = &cobra.Command{
	Use:   "get-current-user-application-key [app_key_id]",
	
	Short: "Get one application key owned by current user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.GetCurrentUserApplicationKey(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-current-user-application-key")

		cmdutil.PrintJSON(res, "application_keys")
	},
}

func init() {
	Cmd.AddCommand(GetCurrentUserApplicationKeyCmd)
}
