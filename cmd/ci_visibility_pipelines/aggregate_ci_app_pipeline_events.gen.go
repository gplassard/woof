package ci_visibility_pipelines

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AggregateCIAppPipelineEventsCmd = &cobra.Command{
	Use: "aggregate-ci-app-pipeline-events",

	Short: "Aggregate pipelines events",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.AggregateCIAppPipelineEvents(client.NewContext(apiKey, appKey, site), datadogV2.CIAppPipelinesAggregateRequest{})
		cmdutil.HandleError(err, "failed to aggregate-ci-app-pipeline-events")

		cmd.Println(cmdutil.FormatJSON(res, "ci_visibility_pipelines"))
	},
}

func init() {
	Cmd.AddCommand(AggregateCIAppPipelineEventsCmd)
}
