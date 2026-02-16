package ci_visibility_pipelines

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
Documentation: https://docs.datadoghq.com/api/latest/ci-visibility-pipelines/#list-ci-app-pipeline-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppPipelineEventsResponse
		var err error

		optionalParams := datadogV2.NewListCIAppPipelineEventsOptionalParameters()

		if cmd.Flags().Changed("filter-query") {
			val, _ := cmd.Flags().GetString("filter-query")
			optionalParams.WithFilterQuery(val)
		}

		if cmd.Flags().Changed("filter-from") {
			val, _ := cmd.Flags().GetString("filter-from")
			optionalParams.WithFilterFrom(val)
		}

		if cmd.Flags().Changed("filter-to") {
			val, _ := cmd.Flags().GetString("filter-to")
			optionalParams.WithFilterTo(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("page-cursor") {
			val, _ := cmd.Flags().GetString("page-cursor")
			optionalParams.WithPageCursor(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCIAppPipelineEvents(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-ci-app-pipeline-events")

		fmt.Println(cmdutil.FormatJSON(res, "cipipeline"))
	},
}

func init() {

	ListCIAppPipelineEventsCmd.Flags().String("filter-query", "", "Search query following log syntax.")

	ListCIAppPipelineEventsCmd.Flags().String("filter-from", "", "Minimum timestamp for requested events.")

	ListCIAppPipelineEventsCmd.Flags().String("filter-to", "", "Maximum timestamp for requested events.")

	ListCIAppPipelineEventsCmd.Flags().String("sort", "", "Order of events in results.")

	ListCIAppPipelineEventsCmd.Flags().String("page-cursor", "", "List following results with a cursor provided in the previous query.")

	ListCIAppPipelineEventsCmd.Flags().Int64("page-limit", 0, "Maximum number of events in the response.")

	Cmd.AddCommand(ListCIAppPipelineEventsCmd)
}
