package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCaseTitleCmd = &cobra.Command{
	Use:   "updatecasetitle",
	Short: "Update case title",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/title")
		fmt.Println("OperationID: UpdateCaseTitle")
	},
}

func init() {
	Cmd.AddCommand(UpdateCaseTitleCmd)
}
