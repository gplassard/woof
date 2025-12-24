package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdatePriorityCmd = &cobra.Command{
	Use:   "updatepriority",
	Short: "Update case priority",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/priority")
		fmt.Println("OperationID: UpdatePriority")
	},
}

func init() {
	Cmd.AddCommand(UpdatePriorityCmd)
}
