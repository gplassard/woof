package agentless_scanning

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAwsOnDemandTasksCmd = &cobra.Command{
	Use:   "list_aws_on_demand_tasks",
	Short: "List AWS on demand tasks",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAgentlessScanningApi(client.NewAPIClient())
		res, _, err := api.ListAwsOnDemandTasks(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_aws_on_demand_tasks: %v", err)
		}

		cmdutil.PrintJSON(res, "aws_resource")
	},
}

func init() {
	Cmd.AddCommand(ListAwsOnDemandTasksCmd)
}
