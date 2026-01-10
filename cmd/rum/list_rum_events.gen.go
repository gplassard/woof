package rum

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRUMEventsCmd = &cobra.Command{
	Use:     "list-rum-events",
	Aliases: []string{"list-events"},
	Short:   "Get a list of RUM events",
	Long: `Get a list of RUM events
Documentation: https://docs.datadoghq.com/api/latest/rum/#list-rum-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RUMEventsResponse
		var err error

		optionalParams := datadogV2.NewListRUMEventsOptionalParameters()

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

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListRUMEvents(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-rum-events")

		fmt.Println(cmdutil.FormatJSON(res, "rum"))
	},
}

func init() {

	ListRUMEventsCmd.Flags().String("filter-query", "", "Search query following RUM syntax.")

	ListRUMEventsCmd.Flags().String("filter-from", "", "Minimum timestamp for requested events.")

	ListRUMEventsCmd.Flags().String("filter-to", "", "Maximum timestamp for requested events.")

	ListRUMEventsCmd.Flags().String("sort", "", "Order of events in results.")

	ListRUMEventsCmd.Flags().String("page-cursor", "", "List following results with a cursor provided in the previous query.")

	ListRUMEventsCmd.Flags().Int64("page-limit", 0, "Maximum number of events in the response.")

	Cmd.AddCommand(ListRUMEventsCmd)
}
