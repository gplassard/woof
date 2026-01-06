package action_connection

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateActionConnectionCmd = &cobra.Command{
	Use:     "create-action-connection",
	Aliases: []string{"create"},
	Short:   "Create a new Action Connection",
	Long: `Create a new Action Connection
Documentation: https://docs.datadoghq.com/api/latest/action-connection/#create-action-connection`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateActionConnectionResponse
		var err error

		var body datadogV2.CreateActionConnectionRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateActionConnection(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-action-connection")

		cmd.Println(cmdutil.FormatJSON(res, "action_connection"))
	},
}

func init() {

	CreateActionConnectionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateActionConnectionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateActionConnectionCmd)
}
