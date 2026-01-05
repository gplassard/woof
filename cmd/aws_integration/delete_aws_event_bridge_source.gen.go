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

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err := api.DeleteAWSEventBridgeSource(client.NewContext(apiKey, appKey, site), datadogV2.AWSEventBridgeDeleteRequest{})
		cmdutil.HandleError(err, "failed to delete-aws-event-bridge-source")

		cmd.Println(cmdutil.FormatJSON(res, "event_bridge"))
	},
}

func init() {
	Cmd.AddCommand(DeleteAWSEventBridgeSourceCmd)
}
