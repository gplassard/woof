package logs

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLogsGetCmd = &cobra.Command{
	Use:     "list-logs-get",
	Aliases: []string{"list-get"},
	Short:   "Search logs (GET)",
	Long: `Search logs (GET)
Documentation: https://docs.datadoghq.com/api/latest/logs/#list-logs-get`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LogsListResponse
		var err error

		optionalParams := datadogV2.NewListLogsGetOptionalParameters()

		if cmd.Flags().Changed("filter-query") {
			val, _ := cmd.Flags().GetString("filter-query")
			optionalParams.WithFilterQuery(val)
		}

		if cmd.Flags().Changed("filter-indexes") {
			val, _ := cmd.Flags().GetString("filter-indexes")
			optionalParams.WithFilterIndexes(val)
		}

		if cmd.Flags().Changed("filter-from") {
			val, _ := cmd.Flags().GetString("filter-from")
			optionalParams.WithFilterFrom(val)
		}

		if cmd.Flags().Changed("filter-to") {
			val, _ := cmd.Flags().GetString("filter-to")
			optionalParams.WithFilterTo(val)
		}

		if cmd.Flags().Changed("filter-storage-tier") {
			val, _ := cmd.Flags().GetString("filter-storage-tier")
			optionalParams.WithFilterStorageTier(val)
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

		api := datadogV2.NewLogsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListLogsGet(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-logs-get")

		cmd.Println(cmdutil.FormatJSON(res, "log"))
	},
}

func init() {

	ListLogsGetCmd.Flags().String("filter-query", "", "Search query following logs syntax.")

	ListLogsGetCmd.Flags().String("filter-indexes", "", "For customers with multiple indexes, the indexes to search. Defaults to '*' which means all indexes")

	ListLogsGetCmd.Flags().String("filter-from", "", "Minimum timestamp for requested logs.")

	ListLogsGetCmd.Flags().String("filter-to", "", "Maximum timestamp for requested logs.")

	ListLogsGetCmd.Flags().String("filter-storage-tier", "", "Specifies the storage type to be used")

	ListLogsGetCmd.Flags().String("sort", "", "Order of logs in results.")

	ListLogsGetCmd.Flags().String("page-cursor", "", "List following results with a cursor provided in the previous query.")

	ListLogsGetCmd.Flags().Int64("page-limit", 0, "Maximum number of logs in the response.")

	Cmd.AddCommand(ListLogsGetCmd)
}
