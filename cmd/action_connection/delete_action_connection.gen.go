package action_connection

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteActionConnectionCmd = &cobra.Command{
	Use:     "delete-action-connection [connection_id]",
	Aliases: []string{"delete"},
	Short:   "Delete an existing Action Connection",
	Long: `Delete an existing Action Connection
Documentation: https://docs.datadoghq.com/api/latest/action-connection/#delete-action-connection`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteActionConnection(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-action-connection")

	},
}

func init() {

	Cmd.AddCommand(DeleteActionConnectionCmd)
}
