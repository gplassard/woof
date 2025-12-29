package action_connection

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAppKeyRegistrationsCmd = &cobra.Command{
	Use:   "list-app-key-registrations",
	
	Short: "List App Key Registrations",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err := api.ListAppKeyRegistrations(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-app-key-registrations")

		cmd.Println(cmdutil.FormatJSON(res, "app_key_registration"))
	},
}

func init() {
	Cmd.AddCommand(ListAppKeyRegistrationsCmd)
}
