package aws_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteAWSEventBridgeSourceCmd = &cobra.Command{
	Use:   "delete-aws-event-bridge-source",
	
	Short: "Delete an Amazon EventBridge source",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.DeleteAWSEventBridgeSource(client.NewContext(apiKey, appKey, site), datadogV2.AWSEventBridgeDeleteRequest{})
		cmdutil.HandleError(err, "failed to delete-aws-event-bridge-source")

		cmdutil.PrintJSON(res, "event_bridge")
	},
}

func init() {
	Cmd.AddCommand(DeleteAWSEventBridgeSourceCmd)
}
