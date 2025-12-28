package action_connection

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAppKeyRegistrationsCmd = &cobra.Command{
	Use:   "list_app_key_registrations",
	Short: "List App Key Registrations",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err := api.ListAppKeyRegistrations(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_app_key_registrations: %v", err)
		}

		cmdutil.PrintJSON(res, "app_key_registration")
	},
}

func init() {
	Cmd.AddCommand(ListAppKeyRegistrationsCmd)
}
