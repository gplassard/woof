package audit

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchAuditLogsCmd = &cobra.Command{
	Use:   "searchauditlogs",
	Short: "Search Audit Logs events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/audit/events/search")
		fmt.Println("OperationID: SearchAuditLogs")
	},
}

func init() {
	Cmd.AddCommand(SearchAuditLogsCmd)
}
