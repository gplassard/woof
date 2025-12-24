package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCaseCmd = &cobra.Command{
	Use:   "createcase",
	Short: "Create a case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cases")
		fmt.Println("OperationID: CreateCase")
	},
}

func init() {
	Cmd.AddCommand(CreateCaseCmd)
}
