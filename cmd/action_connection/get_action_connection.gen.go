package action_connection

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetActionConnectionCmd = &cobra.Command{
	Use:     "get-action-connection [connection_id]",
	Aliases: []string{"get"},
	Short:   "Get an existing Action Connection",
	Long: `Get an existing Action Connection
Documentation: https://docs.datadoghq.com/api/latest/action-connection/#get-action-connection`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetActionConnectionResponse
		var err error

		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		res, _, err = api.GetActionConnection(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-action-connection")

		cmd.Println(cmdutil.FormatJSON(res, "action_connection"))
	},
}

func init() {

	Cmd.AddCommand(GetActionConnectionCmd)
}
