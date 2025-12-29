package ci_visibility_pipelines

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCIAppPipelineEventsCmd = &cobra.Command{
	Use:   "list-ci-app-pipeline-events",
	
	Short: "Get a list of pipelines events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.ListCIAppPipelineEvents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-ci-app-pipeline-events")

		cmd.Println(cmdutil.FormatJSON(res, "cipipeline"))
	},
}

func init() {
	Cmd.AddCommand(ListCIAppPipelineEventsCmd)
}
