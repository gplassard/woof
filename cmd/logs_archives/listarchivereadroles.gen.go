package logs_archives

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListArchiveReadRolesCmd = &cobra.Command{
	Use:   "listarchivereadroles",
	Short: "List read roles for an archive",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/archives/{archive_id}/readers")
		fmt.Println("OperationID: ListArchiveReadRoles")
	},
}

func init() {
	Cmd.AddCommand(ListArchiveReadRolesCmd)
}
