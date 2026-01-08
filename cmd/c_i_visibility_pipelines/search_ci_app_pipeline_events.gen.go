package c_i_visibility_pipelines

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchCIAppPipelineEventsCmd = &cobra.Command{
	Use: "search-ci-app-pipeline-events",

	Short: "Search pipelines events",
	Long: `Search pipelines events
Documentation: https://docs.datadoghq.com/api/latest/c-i-visibility-pipelines/#search-ci-app-pipeline-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppPipelineEventsResponse
		var err error

		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchCIAppPipelineEvents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to search-ci-app-pipeline-events")

		fmt.Println(cmdutil.FormatJSON(res, "c_i_app_pipeline_event"))
	},
}

func init() {

	Cmd.AddCommand(SearchCIAppPipelineEventsCmd)
}
