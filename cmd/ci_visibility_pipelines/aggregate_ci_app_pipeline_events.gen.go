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
	Long: `Aggregate pipelines events
Documentation: https://docs.datadoghq.com/api/latest/ci-visibility-pipelines/#aggregate-ci-app-pipeline-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppPipelinesAnalyticsAggregateResponse
		var err error

		var body datadogV2.CIAppPipelinesAggregateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err = api.AggregateCIAppPipelineEvents(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to aggregate-ci-app-pipeline-events")

		cmd.Println(cmdutil.FormatJSON(res, "ci_visibility_pipelines"))
	},
}

func init() {

	AggregateCIAppPipelineEventsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AggregateCIAppPipelineEventsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AggregateCIAppPipelineEventsCmd)
}
