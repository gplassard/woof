package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCaseCommentCmd = &cobra.Command{
	Use:   "deletecasecomment",
	Short: "Delete case comment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cases/{case_id}/comment/{cell_id}")
		fmt.Println("OperationID: DeleteCaseComment")
	},
}

func init() {
	Cmd.AddCommand(DeleteCaseCommentCmd)
}
