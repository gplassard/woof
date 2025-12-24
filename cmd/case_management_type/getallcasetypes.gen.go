package case_management_type

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAllCaseTypesCmd = &cobra.Command{
	Use:   "getallcasetypes",
	Short: "Get all case types",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cases/types")
		fmt.Println("OperationID: GetAllCaseTypes")
	},
}

func init() {
	Cmd.AddCommand(GetAllCaseTypesCmd)
}
