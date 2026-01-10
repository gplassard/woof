package events

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListEventsCmd = &cobra.Command{
	Use:     "list-events",
	Aliases: []string{"list"},
	Short:   "Get a list of events",
	Long: `Get a list of events
Documentation: https://docs.datadoghq.com/api/latest/events/#list-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.EventsListResponse
		var err error

		optionalParams := datadogV2.NewListEventsOptionalParameters()

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

		api := datadogV2.NewEventsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListEvents(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-events")

		fmt.Println(cmdutil.FormatJSON(res, "event"))
	},
}

func init() {

	ListEventsCmd.Flags().String("filter-query", "", "Search query following events syntax.")

	ListEventsCmd.Flags().String("filter-from", "", "Minimum timestamp for requested events, in milliseconds.")

	ListEventsCmd.Flags().String("filter-to", "", "Maximum timestamp for requested events, in milliseconds.")

	ListEventsCmd.Flags().String("sort", "", "Order of events in results.")

	ListEventsCmd.Flags().String("page-cursor", "", "List following results with a cursor provided in the previous query.")

	ListEventsCmd.Flags().Int64("page-limit", 0, "Maximum number of events in the response.")

	Cmd.AddCommand(ListEventsCmd)
}
