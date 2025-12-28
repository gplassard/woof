package agentless_scanning

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAzureScanOptionsCmd = &cobra.Command{
	Use:   "get_azure_scan_options [subscription_id]",
	Short: "Get Azure scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.GetAzureScanOptions(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get_azure_scan_options: %v", err)
		}

		cmdutil.PrintJSON(res, "azure_scan_options")
	},
}

func init() {
	Cmd.AddCommand(GetAzureScanOptionsCmd)
}
