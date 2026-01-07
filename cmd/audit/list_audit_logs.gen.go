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

		api := datadogV2.NewAuditApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAuditLogs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-audit-logs")

		fmt.Println(cmdutil.FormatJSON(res, "audit"))
	},
}

func init() {

	Cmd.AddCommand(ListAuditLogsCmd)
}
