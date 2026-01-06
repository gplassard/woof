package action_connection

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateActionConnectionCmd = &cobra.Command{
	Use:     "update-action-connection [connection_id]",
	Aliases: []string{"update"},
	Short:   "Update an existing Action Connection",
	Long: `Update an existing Action Connection
Documentation: https://docs.datadoghq.com/api/latest/action-connection/#update-action-connection`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateActionConnectionResponse
		var err error

		var body datadogV2.UpdateActionConnectionRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateActionConnection(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-action-connection")

		cmd.Println(cmdutil.FormatJSON(res, "action_connection"))
	},
}

func init() {

	UpdateActionConnectionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateActionConnectionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateActionConnectionCmd)
}
