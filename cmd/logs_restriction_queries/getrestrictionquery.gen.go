package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetRestrictionQueryCmd = &cobra.Command{
	Use:   "getrestrictionquery",
	Short: "Get a restriction query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/restriction_queries/{restriction_query_id}")
		fmt.Println("OperationID: GetRestrictionQuery")
	},
}

func init() {
	Cmd.AddCommand(GetRestrictionQueryCmd)
}
