package logs

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListLogsGetCmd = &cobra.Command{
	Use:   "listlogsget",
	Short: "Search logs (GET)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/events")
		fmt.Println("OperationID: ListLogsGet")
	},
}

func init() {
	Cmd.AddCommand(ListLogsGetCmd)
}
