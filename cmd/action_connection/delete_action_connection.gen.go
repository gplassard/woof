package action_connection

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteActionConnectionCmd = &cobra.Command{
	Use:     "delete-action-connection [connection_id]",
	Aliases: []string{"delete"},
	Short:   "Delete an existing Action Connection",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		_, err := api.DeleteActionConnection(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-action-connection")

	},
}

func init() {
	Cmd.AddCommand(DeleteActionConnectionCmd)
}
