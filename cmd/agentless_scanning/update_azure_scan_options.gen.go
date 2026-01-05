package agentless_scanning

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateAzureScanOptionsCmd = &cobra.Command{
	Use: "update-azure-scan-options [subscription_id]",

	Short: "Update Azure scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.UpdateAzureScanOptions(client.NewContext(apiKey, appKey, site), args[0], datadogV2.AzureScanOptionsInputUpdate{})
		cmdutil.HandleError(err, "failed to update-azure-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "azure_scan_options"))
	},
}

func init() {
	Cmd.AddCommand(UpdateAzureScanOptionsCmd)
}
