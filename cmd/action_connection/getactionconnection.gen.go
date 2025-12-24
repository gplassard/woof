package action_connection

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetActionConnectionCmd = &cobra.Command{
	Use:   "getactionconnection [connection_id]",
	Short: "Get an existing Action Connection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err := api.GetActionConnection(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getactionconnection: %v", err)
		}

		cmdutil.PrintJSON(res, "action_connection")
	},
}

func init() {
	Cmd.AddCommand(GetActionConnectionCmd)
}
