package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCaseDescriptionCmd = &cobra.Command{
	Use:   "updatecasedescription",
	Short: "Update case description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/{case_id}/description")
		fmt.Println("OperationID: UpdateCaseDescription")
	},
}

func init() {
	Cmd.AddCommand(UpdateCaseDescriptionCmd)
}
