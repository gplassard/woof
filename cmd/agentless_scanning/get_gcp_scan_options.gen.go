package agentless_scanning

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetGcpScanOptionsCmd = &cobra.Command{
	Use: "get-gcp-scan-options [project_id]",

	Short: "Get GCP scan options",
	Long: `Get GCP scan options
Documentation: https://docs.datadoghq.com/api/latest/agentless-scanning/#get-gcp-scan-options`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GcpScanOptions
		var err error

		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetGcpScanOptions(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-gcp-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_scan_options"))
	},
}

func init() {

	Cmd.AddCommand(GetGcpScanOptionsCmd)
}
