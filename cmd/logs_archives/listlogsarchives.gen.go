package logs_archives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListLogsArchivesCmd = &cobra.Command{
	Use:   "listlogsarchives",
	Short: "Get all archives",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/archives")
		fmt.Println("OperationID: ListLogsArchives")
	},
}

func init() {
	Cmd.AddCommand(ListLogsArchivesCmd)
}
