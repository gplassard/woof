package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListUserRestrictionQueriesCmd = &cobra.Command{
	Use:   "listuserrestrictionqueries",
	Short: "Get all restriction queries for a given user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/restriction_queries/user/{user_id}")
		fmt.Println("OperationID: ListUserRestrictionQueries")
	},
}

func init() {
	Cmd.AddCommand(ListUserRestrictionQueriesCmd)
}
