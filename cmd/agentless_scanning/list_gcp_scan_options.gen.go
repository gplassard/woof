package agentless_scanning

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListGcpScanOptionsCmd = &cobra.Command{
	Use:   "list-gcp-scan-options",
	
	Short: "List GCP scan options",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.ListGcpScanOptions(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-gcp-scan-options: %v", err)
		}

		cmdutil.PrintJSON(res, "gcp_scan_options")
	},
}

func init() {
	Cmd.AddCommand(ListGcpScanOptionsCmd)
}
