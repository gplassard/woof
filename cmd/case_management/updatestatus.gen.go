package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateStatusCmd = &cobra.Command{
	Use:   "updatestatus",
	Short: "Update case status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/status")
		fmt.Println("OperationID: UpdateStatus")
	},
}

func init() {
	Cmd.AddCommand(UpdateStatusCmd)
}
