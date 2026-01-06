package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteAWSEventBridgeSourceCmd = &cobra.Command{
	Use: "delete-aws-event-bridge-source",

	Short: "Delete an Amazon EventBridge source",
	Long: `Delete an Amazon EventBridge source
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#delete-aws-event-bridge-source`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSEventBridgeDeleteResponse
		var err error

		var body datadogV2.AWSEventBridgeDeleteRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.DeleteAWSEventBridgeSource(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to delete-aws-event-bridge-source")

		cmd.Println(cmdutil.FormatJSON(res, "event_bridge"))
	},
}

func init() {

	DeleteAWSEventBridgeSourceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DeleteAWSEventBridgeSourceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DeleteAWSEventBridgeSourceCmd)
}
