package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCaseCmd = &cobra.Command{
	Use:   "getcase",
	Short: "Get the details of a case",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cases/{case_id}")
		fmt.Println("OperationID: GetCase")
	},
}

func init() {
	Cmd.AddCommand(GetCaseCmd)
}
