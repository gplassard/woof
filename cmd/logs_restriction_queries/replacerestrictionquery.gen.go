package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ReplaceRestrictionQueryCmd = &cobra.Command{
	Use:   "replacerestrictionquery",
	Short: "Replace a restriction query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/logs/config/restriction_queries/{restriction_query_id}")
		fmt.Println("OperationID: ReplaceRestrictionQuery")
	},
}

func init() {
	Cmd.AddCommand(ReplaceRestrictionQueryCmd)
}
