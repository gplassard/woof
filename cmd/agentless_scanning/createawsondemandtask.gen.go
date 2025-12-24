package agentless_scanning

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAwsOnDemandTaskCmd = &cobra.Command{
	Use:   "createawsondemandtask",
	Short: "Create AWS on demand task",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.CreateAwsOnDemandTask(client.NewContext(apiKey, appKey, site), datadogV2.AwsOnDemandCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createawsondemandtask: %v", err)
		}

		cmdutil.PrintJSON(res, "agentless_scanning")
	},
}

func init() {
	Cmd.AddCommand(CreateAwsOnDemandTaskCmd)
}
