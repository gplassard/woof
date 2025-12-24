package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteRestrictionQueryCmd = &cobra.Command{
	Use:   "deleterestrictionquery",
	Short: "Delete a restriction query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/logs/config/restriction_queries/{restriction_query_id}")
		fmt.Println("OperationID: DeleteRestrictionQuery")
	},
}

func init() {
	Cmd.AddCommand(DeleteRestrictionQueryCmd)
}
