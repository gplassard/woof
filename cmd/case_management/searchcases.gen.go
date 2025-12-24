package case_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchCasesCmd = &cobra.Command{
	Use:   "searchcases",
	Short: "Search cases",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cases")
		fmt.Println("OperationID: SearchCases")
	},
}

func init() {
	Cmd.AddCommand(SearchCasesCmd)
}
