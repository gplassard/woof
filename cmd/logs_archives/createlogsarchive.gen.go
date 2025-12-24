package logs_archives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateLogsArchiveCmd = &cobra.Command{
	Use:   "createlogsarchive",
	Short: "Create an archive",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/logs/config/archives")
		fmt.Println("OperationID: CreateLogsArchive")
	},
}

func init() {
	Cmd.AddCommand(CreateLogsArchiveCmd)
}
