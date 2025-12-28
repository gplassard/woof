package ci_visibility_pipelines

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCIAppPipelineEventsCmd = &cobra.Command{
	Use:   "list_c_i_app_pipeline_events",
	Short: "Get a list of pipelines events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.ListCIAppPipelineEvents(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_c_i_app_pipeline_events: %v", err)
		}

		cmdutil.PrintJSON(res, "cipipeline")
	},
}

func init() {
	Cmd.AddCommand(ListCIAppPipelineEventsCmd)
}
