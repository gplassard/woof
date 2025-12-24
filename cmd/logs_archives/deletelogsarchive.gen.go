package logs_archives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteLogsArchiveCmd = &cobra.Command{
	Use:   "deletelogsarchive",
	Short: "Delete an archive",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/logs/config/archives/{archive_id}")
		fmt.Println("OperationID: DeleteLogsArchive")
	},
}

func init() {
	Cmd.AddCommand(DeleteLogsArchiveCmd)
}
