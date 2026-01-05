package ci_visibility_pipelines

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var AggregateCIAppPipelineEventsCmd = &cobra.Command{
	Use: "aggregate-ci-app-pipeline-events [payload]",

	Short: "Aggregate pipelines events",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppPipelinesAnalyticsAggregateResponse
		var err error

		var body datadogV2.CIAppPipelinesAggregateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err = api.AggregateCIAppPipelineEvents(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to aggregate-ci-app-pipeline-events")

		cmd.Println(cmdutil.FormatJSON(res, "ci_visibility_pipelines"))
	},
}

func init() {
	Cmd.AddCommand(AggregateCIAppPipelineEventsCmd)
}
