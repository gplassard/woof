package logs_archives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RemoveRoleFromArchiveCmd = &cobra.Command{
	Use:   "removerolefromarchive",
	Short: "Revoke role from an archive",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/logs/config/archives/{archive_id}/readers")
		fmt.Println("OperationID: RemoveRoleFromArchive")
	},
}

func init() {
	Cmd.AddCommand(RemoveRoleFromArchiveCmd)
}
