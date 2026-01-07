package audit

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchAuditLogsCmd = &cobra.Command{
	Use:     "search-audit-logs",
	Aliases: []string{"search-logs"},
	Short:   "Search Audit Logs events",
	Long: `Search Audit Logs events
Documentation: https://docs.datadoghq.com/api/latest/audit/#search-audit-logs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AuditLogsEventsResponse
		var err error

		var body datadogV2.SearchAuditLogsOptionalParameters
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAuditApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchAuditLogs(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-audit-logs")

		fmt.Println(cmdutil.FormatJSON(res, "audit"))
	},
}

func init() {

	SearchAuditLogsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SearchAuditLogsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SearchAuditLogsCmd)
}
