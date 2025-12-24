package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AddRoleToRestrictionQueryCmd = &cobra.Command{
	Use:   "addroletorestrictionquery",
	Short: "Grant role to a restriction query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/logs/config/restriction_queries/{restriction_query_id}/roles")
		fmt.Println("OperationID: AddRoleToRestrictionQuery")
	},
}

func init() {
	Cmd.AddCommand(AddRoleToRestrictionQueryCmd)
}
