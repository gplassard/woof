package aws_integration

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAWSEventBridgeSourceCmd = &cobra.Command{
	Use:   "create-aws-event-bridge-source",
	
	Short: "Create an Amazon EventBridge source",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateAWSEventBridgeSource(client.NewContext(apiKey, appKey, site), datadogV2.AWSEventBridgeCreateRequest{})
		cmdutil.HandleError(err, "failed to create-aws-event-bridge-source")

		cmd.Println(cmdutil.FormatJSON(res, "event_bridge"))
	},
}

func init() {
	Cmd.AddCommand(CreateAWSEventBridgeSourceCmd)
}
