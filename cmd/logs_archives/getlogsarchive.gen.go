package logs_archives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetLogsArchiveCmd = &cobra.Command{
	Use:   "getlogsarchive",
	Short: "Get an archive",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/archives/{archive_id}")
		fmt.Println("OperationID: GetLogsArchive")
	},
}

func init() {
	Cmd.AddCommand(GetLogsArchiveCmd)
}
