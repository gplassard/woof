package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UnassignCaseCmd = &cobra.Command{
	Use:   "unassigncase",
	Short: "Unassign case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/unassign")
		fmt.Println("OperationID: UnassignCase")
	},
}

func init() {
	Cmd.AddCommand(UnassignCaseCmd)
}
