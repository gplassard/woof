package logs_archives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateLogsArchiveOrderCmd = &cobra.Command{
	Use:   "updatelogsarchiveorder",
	Short: "Update archive order",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/logs/config/archive-order")
		fmt.Println("OperationID: UpdateLogsArchiveOrder")
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsArchiveOrderCmd)
}
