package case_management_type

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCaseTypeCmd = &cobra.Command{
	Use:   "deletecasetype",
	Short: "Delete a case type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cases/types/{case_type_id}")
		fmt.Println("OperationID: DeleteCaseType")
	},
}

func init() {
	Cmd.AddCommand(DeleteCaseTypeCmd)
}
