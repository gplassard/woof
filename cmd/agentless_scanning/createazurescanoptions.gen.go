package agentless_scanning

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAzureScanOptionsCmd = &cobra.Command{
	Use:   "createazurescanoptions",
	Short: "Create Azure scan options",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.CreateAzureScanOptions(client.NewContext(apiKey, appKey, site), datadogV2.AzureScanOptions{})
		if err != nil {
			log.Fatalf("failed to createazurescanoptions: %v", err)
		}

		cmdutil.PrintJSON(res, "agentless_scanning")
	},
}

func init() {
	Cmd.AddCommand(CreateAzureScanOptionsCmd)
}
