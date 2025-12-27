package agentless_scanning

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateGcpScanOptionsCmd = &cobra.Command{
	Use:   "creategcpscanoptions",
	Short: "Create GCP scan options",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.CreateGcpScanOptions(client.NewContext(apiKey, appKey, site), datadogV2.GcpScanOptions{})
		if err != nil {
			log.Fatalf("failed to creategcpscanoptions: %v", err)
		}

		cmdutil.PrintJSON(res, "agentless_scanning")
	},
}

func init() {
	Cmd.AddCommand(CreateGcpScanOptionsCmd)
}
