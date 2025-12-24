package test_optimization

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchFlakyTestsCmd = &cobra.Command{
	Use:   "searchflakytests",
	Short: "Search flaky tests",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/test/flaky-test-management/tests")
		fmt.Println("OperationID: SearchFlakyTests")
	},
}

func init() {
	Cmd.AddCommand(SearchFlakyTestsCmd)
}
