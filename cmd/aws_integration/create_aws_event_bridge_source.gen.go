package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAWSEventBridgeSourceCmd = &cobra.Command{
	Use: "create-aws-event-bridge-source",

	Short: "Create an Amazon EventBridge source",
	Long: `Create an Amazon EventBridge source
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#create-aws-event-bridge-source`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSEventBridgeCreateResponse
		var err error

		var body datadogV2.AWSEventBridgeCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateAWSEventBridgeSource(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-aws-event-bridge-source")

		cmd.Println(cmdutil.FormatJSON(res, "event_bridge"))
	},
}

func init() {

	CreateAWSEventBridgeSourceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateAWSEventBridgeSourceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateAWSEventBridgeSourceCmd)
}
