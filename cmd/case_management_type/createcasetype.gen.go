package case_management_type

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCaseTypeCmd = &cobra.Command{
	Use:   "createcasetype",
	Short: "Create a case type",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases/types")
		fmt.Println("OperationID: CreateCaseType")
	},
}

func init() {
	Cmd.AddCommand(CreateCaseTypeCmd)
}
