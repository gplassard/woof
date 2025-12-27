package audit

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchAuditLogsCmd = &cobra.Command{
	Use:   "search_audit_logs",
	Short: "Search Audit Logs events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAuditApi(client.NewAPIClient())
		res, _, err := api.SearchAuditLogs(client.NewContext(apiKey, appKey, site), *datadogV2.NewSearchAuditLogsOptionalParameters())
		if err != nil {
			log.Fatalf("failed to search_audit_logs: %v", err)
		}

		cmdutil.PrintJSON(res, "audit")
	},
}

func init() {
	Cmd.AddCommand(SearchAuditLogsCmd)
}
