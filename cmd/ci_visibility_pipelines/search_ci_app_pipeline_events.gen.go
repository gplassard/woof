package ci_visibility_pipelines

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SearchCIAppPipelineEventsCmd = &cobra.Command{
	Use: "search-ci-app-pipeline-events [payload]",

	Short: "Search pipelines events",
	Long: `Search pipelines events
Documentation: https://docs.datadoghq.com/api/latest/ci-visibility-pipelines/#search-ci-app-pipeline-events`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppPipelineEventsResponse
		var err error

		var body datadogV2.SearchCIAppPipelineEventsOptionalParameters
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err = api.SearchCIAppPipelineEvents(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-ci-app-pipeline-events")

		cmd.Println(cmdutil.FormatJSON(res, "cipipeline"))
	},
}

func init() {
	Cmd.AddCommand(SearchCIAppPipelineEventsCmd)
}
