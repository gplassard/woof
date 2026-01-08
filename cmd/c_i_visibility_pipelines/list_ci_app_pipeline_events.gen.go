package c_i_visibility_pipelines

import (
	"fmt"
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
Documentation: https://docs.datadoghq.com/api/latest/c-i-visibility-pipelines/#list-ci-app-pipeline-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppPipelineEventsResponse
		var err error

		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCIAppPipelineEvents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-ci-app-pipeline-events")

		fmt.Println(cmdutil.FormatJSON(res, "c_i_app_pipeline_event"))
	},
}

func init() {

	Cmd.AddCommand(ListCIAppPipelineEventsCmd)
}
