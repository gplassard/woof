package logs

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListLogsCmd = &cobra.Command{
	Use:   "listlogs",
	Short: "Search logs (POST)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/logs/events/search")
		fmt.Println("OperationID: ListLogs")
	},
}

func init() {
	Cmd.AddCommand(ListLogsCmd)
}
