package audit

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAuditLogsCmd = &cobra.Command{
	Use:     "list-audit-logs",
	Aliases: []string{"list-logs"},
	Short:   "Get a list of Audit Logs events",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAuditApi(client.NewAPIClient())
		res, _, err := api.ListAuditLogs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-audit-logs")

		cmd.Println(cmdutil.FormatJSON(res, "audit"))
	},
}

func init() {
	Cmd.AddCommand(ListAuditLogsCmd)
}
