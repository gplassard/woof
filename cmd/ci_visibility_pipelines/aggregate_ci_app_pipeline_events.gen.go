package ci_visibility_pipelines

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AggregateCIAppPipelineEventsCmd = &cobra.Command{
	Use:   "aggregate-ci-app-pipeline-events",
	
	Short: "Aggregate pipelines events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.AggregateCIAppPipelineEvents(client.NewContext(apiKey, appKey, site), datadogV2.CIAppPipelinesAggregateRequest{})
		if err != nil {
			log.Fatalf("failed to aggregate-ci-app-pipeline-events: %v", err)
		}

		cmdutil.PrintJSON(res, "ci_visibility_pipelines")
	},
}

func init() {
	Cmd.AddCommand(AggregateCIAppPipelineEventsCmd)
}
