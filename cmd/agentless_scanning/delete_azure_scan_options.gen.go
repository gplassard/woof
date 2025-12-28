package agentless_scanning

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteAzureScanOptionsCmd = &cobra.Command{
	Use:   "delete-azure-scan-options [subscription_id]",
	
	Short: "Delete Azure scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		_, err := api.DeleteAzureScanOptions(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-azure-scan-options: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteAzureScanOptionsCmd)
}
