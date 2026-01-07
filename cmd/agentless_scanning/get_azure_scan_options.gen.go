package agentless_scanning

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAzureScanOptionsCmd = &cobra.Command{
	Use: "get-azure-scan-options [subscription_id]",

	Short: "Get Azure scan options",
	Long: `Get Azure scan options
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#get-azure-scan-options`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AzureScanOptions
		var err error

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAzureScanOptions(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-azure-scan-options")

		fmt.Println(cmdutil.FormatJSON(res, "azure_scan_options"))
	},
}

func init() {

	Cmd.AddCommand(GetAzureScanOptionsCmd)
}
