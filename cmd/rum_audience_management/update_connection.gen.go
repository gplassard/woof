package rum_audience_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateConnectionCmd = &cobra.Command{
	Use: "update-connection [entity]",

	Short: "Update connection",
	Long: `Update connection
Documentation: https://docs.datadoghq.com/api/latest/rum-audience-management/#update-connection`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.UpdateConnectionRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.UpdateConnection(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-connection")

	},
}

func init() {

	UpdateConnectionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateConnectionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateConnectionCmd)
}
