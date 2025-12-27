package agentless_scanning

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAwsOnDemandTaskCmd = &cobra.Command{
	Use:   "create_aws_on_demand_task",
	Short: "Create AWS on demand task",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.CreateAwsOnDemandTask(client.NewContext(apiKey, appKey, site), datadogV2.AwsOnDemandCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_aws_on_demand_task: %v", err)
		}

		cmdutil.PrintJSON(res, "agentless_scanning")
	},
}

func init() {
	Cmd.AddCommand(CreateAwsOnDemandTaskCmd)
}
