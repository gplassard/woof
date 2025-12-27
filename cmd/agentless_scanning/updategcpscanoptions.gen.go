package agentless_scanning

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateGcpScanOptionsCmd = &cobra.Command{
	Use:   "updategcpscanoptions [project_id]",
	Short: "Update GCP scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.UpdateGcpScanOptions(client.NewContext(apiKey, appKey, site), args[0], datadogV2.GcpScanOptionsInputUpdate{})
		if err != nil {
			log.Fatalf("failed to updategcpscanoptions: %v", err)
		}

		cmdutil.PrintJSON(res, "agentless_scanning")
	},
}

func init() {
	Cmd.AddCommand(UpdateGcpScanOptionsCmd)
}
