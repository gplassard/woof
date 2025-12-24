package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateRestrictionQueryCmd = &cobra.Command{
	Use:   "updaterestrictionquery",
	Short: "Update a restriction query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/logs/config/restriction_queries/{restriction_query_id}")
		fmt.Println("OperationID: UpdateRestrictionQuery")
	},
}

func init() {
	Cmd.AddCommand(UpdateRestrictionQueryCmd)
}
