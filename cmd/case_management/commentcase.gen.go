package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CommentCaseCmd = &cobra.Command{
	Use:   "commentcase",
	Short: "Comment case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/comment")
		fmt.Println("OperationID: CommentCase")
	},
}

func init() {
	Cmd.AddCommand(CommentCaseCmd)
}
