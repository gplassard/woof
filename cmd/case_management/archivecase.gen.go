package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ArchiveCaseCmd = &cobra.Command{
	Use:   "archivecase",
	Short: "Archive case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/archive")
		fmt.Println("OperationID: ArchiveCase")
	},
}

func init() {
	Cmd.AddCommand(ArchiveCaseCmd)
}
