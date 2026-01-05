package rum_audience_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateConnectionCmd = &cobra.Command{
	Use: "create-connection [entity]",

	Short: "Create connection",
	Long: `Create connection
Documentation: https://docs.datadoghq.com/api/latest/rum-audience-management/#create-connection`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.CreateConnectionRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		_, err = api.CreateConnection(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-connection")

	},
}

func init() {

	CreateConnectionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateConnectionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateConnectionCmd)
}
