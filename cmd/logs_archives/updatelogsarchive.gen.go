package logs_archives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateLogsArchiveCmd = &cobra.Command{
	Use:   "updatelogsarchive",
	Short: "Update an archive",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/logs/config/archives/{archive_id}")
		fmt.Println("OperationID: UpdateLogsArchive")
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsArchiveCmd)
}
