package spans

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSpansGetCmd = &cobra.Command{
	Use:     "list-spans-get",
	Aliases: []string{"list-get"},
	Short:   "Get a list of spans",
	Long: `Get a list of spans
Documentation: https://docs.datadoghq.com/api/latest/spans/#list-spans-get`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SpansListResponse
		var err error

		optionalParams := datadogV2.NewListSpansGetOptionalParameters()

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

		api := datadogV2.NewSpansApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListSpansGet(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-spans-get")

		cmd.Println(cmdutil.FormatJSON(res, "spans"))
	},
}

func init() {

	ListSpansGetCmd.Flags().String("filter-query", "", "Search query following spans syntax.")

	ListSpansGetCmd.Flags().String("filter-from", "", "Minimum timestamp for requested spans. Supports date-time ISO8601, date math, and regular timestamps (milliseconds).")

	ListSpansGetCmd.Flags().String("filter-to", "", "Maximum timestamp for requested spans. Supports date-time ISO8601, date math, and regular timestamps (milliseconds).")

	ListSpansGetCmd.Flags().String("sort", "", "Order of spans in results.")

	ListSpansGetCmd.Flags().String("page-cursor", "", "List following results with a cursor provided in the previous query.")

	ListSpansGetCmd.Flags().Int64("page-limit", 0, "Maximum number of spans in the response.")

	Cmd.AddCommand(ListSpansGetCmd)
}
