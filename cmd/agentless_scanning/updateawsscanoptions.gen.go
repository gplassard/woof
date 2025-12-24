package agentless_scanning

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateAwsScanOptionsCmd = &cobra.Command{
	Use:   "updateawsscanoptions [account_id]",
	Short: "Update AWS scan options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		_, err := api.UpdateAwsScanOptions(client.NewContext(apiKey, appKey, site), args[0], datadogV2.AwsScanOptionsUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updateawsscanoptions: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(UpdateAwsScanOptionsCmd)
}
