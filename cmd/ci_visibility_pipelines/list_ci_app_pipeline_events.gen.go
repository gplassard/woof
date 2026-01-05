package ci_visibility_pipelines

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCIAppPipelineEventsCmd = &cobra.Command{
	Use: "list-ci-app-pipeline-events",

	Short: "Get a list of pipelines events",
	Long: `Get a list of pipelines events
Documentation: https://docs.datadoghq.com/api/latest/ci-visibility-pipelines/#list-ci-app-pipeline-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppPipelineEventsResponse
		var err error

		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err = api.ListCIAppPipelineEvents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-ci-app-pipeline-events")

		cmd.Println(cmdutil.FormatJSON(res, "cipipeline"))
	},
}

func init() {

	Cmd.AddCommand(ListCIAppPipelineEventsCmd)
}
