package aws_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteAWSEventBridgeSourceCmd = &cobra.Command{
	Use:   "delete-a-w-s-event-bridge-source",
	
	Short: "Delete an Amazon EventBridge source",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.DeleteAWSEventBridgeSource(client.NewContext(apiKey, appKey, site), datadogV2.AWSEventBridgeDeleteRequest{})
		if err != nil {
			log.Fatalf("failed to delete-a-w-s-event-bridge-source: %v", err)
		}

		cmdutil.PrintJSON(res, "event_bridge")
	},
}

func init() {
	Cmd.AddCommand(DeleteAWSEventBridgeSourceCmd)
}
