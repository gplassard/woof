package action_connection

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateActionConnectionCmd = &cobra.Command{
	Use:   "create_action_connection",
	Short: "Create a new Action Connection",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err := api.CreateActionConnection(client.NewContext(apiKey, appKey, site), datadogV2.CreateActionConnectionRequest{})
		if err != nil {
			log.Fatalf("failed to create_action_connection: %v", err)
		}

		cmdutil.PrintJSON(res, "action_connection")
	},
}

func init() {
	Cmd.AddCommand(CreateActionConnectionCmd)
}
