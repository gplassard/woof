package error_tracking

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchIssuesCmd = &cobra.Command{
	Use:   "searchissues",
	Short: "Search error tracking issues",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/error-tracking/issues/search")
		fmt.Println("OperationID: SearchIssues")
	},
}

func init() {
	Cmd.AddCommand(SearchIssuesCmd)
}
