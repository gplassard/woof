package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAzureScanOptionsCmd = &cobra.Command{
	Use: "list-azure-scan-options",

	Short: "List Azure scan options",
	Long: `List Azure scan options
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#list-azure-scan-options`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AzureScanOptionsArray
		var err error

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAzureScanOptions(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-azure-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "azure_scan_options"))
	},
}

func init() {

	Cmd.AddCommand(ListAzureScanOptionsCmd)
}
