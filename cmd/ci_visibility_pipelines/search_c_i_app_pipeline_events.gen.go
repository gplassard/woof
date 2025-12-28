package ci_visibility_pipelines

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchCIAppPipelineEventsCmd = &cobra.Command{
	Use:   "search-c-i-app-pipeline-events",
	Short: "Search pipelines events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.SearchCIAppPipelineEvents(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchCIAppPipelineEventsOptionalParameters())
		if err != nil {
			log.Fatalf("failed to search-c-i-app-pipeline-events: %v", err)
		}

		cmdutil.PrintJSON(res, "cipipeline")
	},
}

func init() {
	Cmd.AddCommand(SearchCIAppPipelineEventsCmd)
}
