package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UnarchiveCaseCmd = &cobra.Command{
	Use:   "unarchivecase",
	Short: "Unarchive case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/unarchive")
		fmt.Println("OperationID: UnarchiveCase")
	},
}

func init() {
	Cmd.AddCommand(UnarchiveCaseCmd)
}
