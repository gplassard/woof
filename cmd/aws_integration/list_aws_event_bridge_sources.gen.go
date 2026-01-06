package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAWSEventBridgeSourcesCmd = &cobra.Command{
	Use: "list-aws-event-bridge-sources",

	Short: "Get all Amazon EventBridge sources",
	Long: `Get all Amazon EventBridge sources
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#list-aws-event-bridge-sources`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSEventBridgeListResponse
		var err error

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAWSEventBridgeSources(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-aws-event-bridge-sources")

		cmd.Println(cmdutil.FormatJSON(res, "event_bridge"))
	},
}

func init() {

	Cmd.AddCommand(ListAWSEventBridgeSourcesCmd)
}
