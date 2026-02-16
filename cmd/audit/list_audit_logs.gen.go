package audit

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAuditLogsCmd = &cobra.Command{
	Use:     "list-audit-logs",
	Aliases: []string{"list-logs"},
	Short:   "Get a list of Audit Logs events",
	Long: `Get a list of Audit Logs events
Documentation: https://docs.datadoghq.com/api/latest/audit/#list-audit-logs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AuditLogsEventsResponse
		var err error

		optionalParams := datadogV2.NewListAuditLogsOptionalParameters()

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

		api := datadogV2.NewAuditApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAuditLogs(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-audit-logs")

		fmt.Println(cmdutil.FormatJSON(res, "audit"))
	},
}

func init() {

	ListAuditLogsCmd.Flags().String("filter-query", "", "Search query following Audit Logs syntax.")

	ListAuditLogsCmd.Flags().String("filter-from", "", "Minimum timestamp for requested events.")

	ListAuditLogsCmd.Flags().String("filter-to", "", "Maximum timestamp for requested events.")

	ListAuditLogsCmd.Flags().String("sort", "", "Order of events in results.")

	ListAuditLogsCmd.Flags().String("page-cursor", "", "List following results with a cursor provided in the previous query.")

	ListAuditLogsCmd.Flags().Int64("page-limit", 0, "Maximum number of events in the response.")

	Cmd.AddCommand(ListAuditLogsCmd)
}
