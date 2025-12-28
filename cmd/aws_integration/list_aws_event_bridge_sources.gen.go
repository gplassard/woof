package aws_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAWSEventBridgeSourcesCmd = &cobra.Command{
	Use:   "list-aws-event-bridge-sources",
	
	Short: "Get all Amazon EventBridge sources",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListAWSEventBridgeSources(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-aws-event-bridge-sources")

		cmdutil.PrintJSON(res, "event_bridge")
	},
}

func init() {
	Cmd.AddCommand(ListAWSEventBridgeSourcesCmd)
}
