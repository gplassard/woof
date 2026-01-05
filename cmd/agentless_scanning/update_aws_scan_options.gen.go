package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateAwsScanOptionsCmd = &cobra.Command{
	Use: "update-aws-scan-options [account_id] [payload]",

	Short: "Update AWS scan options",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.AwsScanOptionsUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		_, err = api.UpdateAwsScanOptions(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-aws-scan-options")

	},
}

func init() {
	Cmd.AddCommand(UpdateAwsScanOptionsCmd)
}
