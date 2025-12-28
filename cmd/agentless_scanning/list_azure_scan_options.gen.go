package agentless_scanning

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAzureScanOptionsCmd = &cobra.Command{
	Use:   "list-azure-scan-options",
	
	Short: "List Azure scan options",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.ListAzureScanOptions(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-azure-scan-options: %v", err)
		}

		cmdutil.PrintJSON(res, "azure_scan_options")
	},
}

func init() {
	Cmd.AddCommand(ListAzureScanOptionsCmd)
}
