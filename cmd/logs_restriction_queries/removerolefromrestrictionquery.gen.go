package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RemoveRoleFromRestrictionQueryCmd = &cobra.Command{
	Use:   "removerolefromrestrictionquery",
	Short: "Revoke role from a restriction query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/logs/config/restriction_queries/{restriction_query_id}/roles")
		fmt.Println("OperationID: RemoveRoleFromRestrictionQuery")
	},
}

func init() {
	Cmd.AddCommand(RemoveRoleFromRestrictionQueryCmd)
}
