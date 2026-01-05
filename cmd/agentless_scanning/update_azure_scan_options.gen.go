package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateAzureScanOptionsCmd = &cobra.Command{
	Use: "update-azure-scan-options [subscription_id] [payload]",

	Short: "Update Azure scan options",
	Long: `Update Azure scan options
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#update-azure-scan-options`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AzureScanOptions
		var err error

		var body datadogV2.AzureScanOptionsInputUpdate
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err = api.UpdateAzureScanOptions(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-azure-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "azure_scan_options"))
	},
}

func init() {
	Cmd.AddCommand(UpdateAzureScanOptionsCmd)
}
