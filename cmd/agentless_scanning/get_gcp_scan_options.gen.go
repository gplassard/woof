package agentless_scanning

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetGcpScanOptionsCmd = &cobra.Command{
	Use:   "get-gcp-scan-options [project_id]",
	
	Short: "Get GCP scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.GetGcpScanOptions(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-gcp-scan-options")

		cmdutil.PrintJSON(res, "gcp_scan_options")
	},
}

func init() {
	Cmd.AddCommand(GetGcpScanOptionsCmd)
}
