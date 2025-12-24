package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AssignCaseCmd = &cobra.Command{
	Use:   "assigncase",
	Short: "Assign case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/assign")
		fmt.Println("OperationID: AssignCase")
	},
}

func init() {
	Cmd.AddCommand(AssignCaseCmd)
}
