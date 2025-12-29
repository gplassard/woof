package ci_visibility_pipelines

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCIAppPipelineEventCmd = &cobra.Command{
	Use:   "create-ci-app-pipeline-event",
	
	Short: "Send pipeline event",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.CreateCIAppPipelineEvent(client.NewContext(apiKey, appKey, site), datadogV2.CIAppCreatePipelineEventRequest{})
		cmdutil.HandleError(err, "failed to create-ci-app-pipeline-event")

		cmd.Println(cmdutil.FormatJSON(res, "ci_visibility_pipelines"))
	},
}

func init() {
	Cmd.AddCommand(CreateCIAppPipelineEventCmd)
}
