package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var DeleteAWSEventBridgeSourceCmd = &cobra.Command{
	Use: "delete-aws-event-bridge-source [payload]",

	Short: "Delete an Amazon EventBridge source",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSEventBridgeDeleteResponse
		var err error

		var body datadogV2.AWSEventBridgeDeleteRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err = api.DeleteAWSEventBridgeSource(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to delete-aws-event-bridge-source")

		cmd.Println(cmdutil.FormatJSON(res, "event_bridge"))
	},
}

func init() {
	Cmd.AddCommand(DeleteAWSEventBridgeSourceCmd)
}
