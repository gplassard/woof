package action_connection

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateActionConnectionCmd = &cobra.Command{
	Use:   "update-action-connection [connection_id]",
	Aliases: []string{ "update", },
	Short: "Update an existing Action Connection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err := api.UpdateActionConnection(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateActionConnectionRequest{})
		cmdutil.HandleError(err, "failed to update-action-connection")

		cmdutil.PrintJSON(res, "action_connection")
	},
}

func init() {
	Cmd.AddCommand(UpdateActionConnectionCmd)
}
