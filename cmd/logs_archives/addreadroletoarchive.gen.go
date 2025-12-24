package logs_archives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AddReadRoleToArchiveCmd = &cobra.Command{
	Use:   "addreadroletoarchive",
	Short: "Grant role to an archive",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/logs/config/archives/{archive_id}/readers")
		fmt.Println("OperationID: AddReadRoleToArchive")
	},
}

func init() {
	Cmd.AddCommand(AddReadRoleToArchiveCmd)
}
