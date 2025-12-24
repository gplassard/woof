package audit

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAuditLogsCmd = &cobra.Command{
	Use:   "listauditlogs",
	Short: "Get a list of Audit Logs events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/audit/events")
		fmt.Println("OperationID: ListAuditLogs")
	},
}

func init() {
	Cmd.AddCommand(ListAuditLogsCmd)
}
