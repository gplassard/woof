package agentless_scanning

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAwsScanOptionsCmd = &cobra.Command{
	Use:   "createawsscanoptions",
	Short: "Create AWS scan options",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.CreateAwsScanOptions(client.NewContext(apiKey, appKey, site), datadogV2.AwsScanOptionsCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createawsscanoptions: %v", err)
		}

		cmdutil.PrintJSON(res, "agentless_scanning")
	},
}

func init() {
	Cmd.AddCommand(CreateAwsScanOptionsCmd)
}
