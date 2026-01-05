package agentless_scanning

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateAwsScanOptionsCmd = &cobra.Command{
	Use: "update-aws-scan-options [account_id]",

	Short: "Update AWS scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		_, err := api.UpdateAwsScanOptions(client.NewContext(apiKey, appKey, site), args[0], datadogV2.AwsScanOptionsUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-aws-scan-options")

	},
}

func init() {
	Cmd.AddCommand(UpdateAwsScanOptionsCmd)
}
