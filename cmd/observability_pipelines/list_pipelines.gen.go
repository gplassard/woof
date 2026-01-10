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

		optionalParams := datadogV2.NewListPipelinesOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListPipelines(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-pipelines")

		fmt.Println(cmdutil.FormatJSON(res, "observability_pipelines"))
	},
}

func init() {

	ListPipelinesCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListPipelinesCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	Cmd.AddCommand(ListPipelinesCmd)
}
