package audit

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var SearchAuditLogsCmd = &cobra.Command{
	Use:     "search-audit-logs [payload]",
	Aliases: []string{"search-logs"},
	Short:   "Search Audit Logs events",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.SearchAuditLogsOptionalParameters
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAuditApi(client.NewAPIClient())
		res, _, err := api.SearchAuditLogs(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-audit-logs")

		cmd.Println(cmdutil.FormatJSON(res, "audit"))
	},
}

func init() {
	Cmd.AddCommand(SearchAuditLogsCmd)
}
