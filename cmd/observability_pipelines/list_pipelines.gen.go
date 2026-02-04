package observability_pipelines

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListPipelinesCmd = &cobra.Command{
	Use: "list-pipelines",

	Short: "List pipelines",
	Long: `List pipelines
Documentation: https://docs.datadoghq.com/api/latest/observability-pipelines/#list-pipelines`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListPipelinesResponse
		var err error

		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListPipelines(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-pipelines")

		fmt.Println(cmdutil.FormatJSON(res, "pipeline"))
	},
}

func init() {

	Cmd.AddCommand(ListPipelinesCmd)
}
