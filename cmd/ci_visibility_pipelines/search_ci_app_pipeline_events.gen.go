package ci_visibility_pipelines

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchCIAppPipelineEventsCmd = &cobra.Command{
	Use: "search-ci-app-pipeline-events",

	Short: "Search pipelines events",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.SearchCIAppPipelineEvents(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchCIAppPipelineEventsOptionalParameters())
		cmdutil.HandleError(err, "failed to search-ci-app-pipeline-events")

		cmd.Println(cmdutil.FormatJSON(res, "cipipeline"))
	},
}

func init() {
	Cmd.AddCommand(SearchCIAppPipelineEventsCmd)
}
