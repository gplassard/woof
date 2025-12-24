package action_connection

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAppKeyRegistrationsCmd = &cobra.Command{
	Use:   "listappkeyregistrations",
	Short: "List App Key Registrations",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err := api.ListAppKeyRegistrations(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listappkeyregistrations: %v", err)
		}

		cmdutil.PrintJSON(res, "action_connection")
	},
}

func init() {
	Cmd.AddCommand(ListAppKeyRegistrationsCmd)
}
