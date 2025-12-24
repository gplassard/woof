package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListRestrictionQueriesCmd = &cobra.Command{
	Use:   "listrestrictionqueries",
	Short: "List restriction queries",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/restriction_queries")
		fmt.Println("OperationID: ListRestrictionQueries")
	},
}

func init() {
	Cmd.AddCommand(ListRestrictionQueriesCmd)
}
