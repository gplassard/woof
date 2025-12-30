package audit

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchAuditLogsCmd = &cobra.Command{
	Use:     "search-audit-logs",
	Aliases: []string{"search-logs"},
	Short:   "Search Audit Logs events",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAuditApi(client.NewAPIClient())
		res, _, err := api.SearchAuditLogs(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchAuditLogsOptionalParameters())
		cmdutil.HandleError(err, "failed to search-audit-logs")

		cmd.Println(cmdutil.FormatJSON(res, "audit"))
	},
}

func init() {
	Cmd.AddCommand(SearchAuditLogsCmd)
}
