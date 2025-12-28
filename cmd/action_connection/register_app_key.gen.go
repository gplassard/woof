package action_connection

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var RegisterAppKeyCmd = &cobra.Command{
	Use:   "register-app-key [app_key_id]",
	
	Short: "Register a new App Key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err := api.RegisterAppKey(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to register-app-key")

		cmdutil.PrintJSON(res, "app_key_registration")
	},
}

func init() {
	Cmd.AddCommand(RegisterAppKeyCmd)
}
