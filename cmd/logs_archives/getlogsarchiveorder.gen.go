package logs_archives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetLogsArchiveOrderCmd = &cobra.Command{
	Use:   "getlogsarchiveorder",
	Short: "Get archive order",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/archive-order")
		fmt.Println("OperationID: GetLogsArchiveOrder")
	},
}

func init() {
	Cmd.AddCommand(GetLogsArchiveOrderCmd)
}
