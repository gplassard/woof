package ci_visibility_pipelines

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCIAppPipelineEventsCmd = &cobra.Command{
	Use:   "listciapppipelineevents",
	Short: "Get a list of pipelines events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.ListCIAppPipelineEvents(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listciapppipelineevents: %v", err)
		}

		cmdutil.PrintJSON(res, "ci_visibility_pipelines")
	},
}

func init() {
	Cmd.AddCommand(ListCIAppPipelineEventsCmd)
}
