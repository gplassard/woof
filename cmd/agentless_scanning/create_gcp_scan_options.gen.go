package agentless_scanning

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateGcpScanOptionsCmd = &cobra.Command{
	Use:   "create-gcp-scan-options",
	
	Short: "Create GCP scan options",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.CreateGcpScanOptions(client.NewContext(apiKey, appKey, site), datadogV2.GcpScanOptions{})
		cmdutil.HandleError(err, "failed to create-gcp-scan-options")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_scan_options"))
	},
}

func init() {
	Cmd.AddCommand(CreateGcpScanOptionsCmd)
}
