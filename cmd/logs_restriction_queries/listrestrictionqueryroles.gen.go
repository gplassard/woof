package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListRestrictionQueryRolesCmd = &cobra.Command{
	Use:   "listrestrictionqueryroles",
	Short: "List roles for a restriction query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/restriction_queries/{restriction_query_id}/roles")
		fmt.Println("OperationID: ListRestrictionQueryRoles")
	},
}

func init() {
	Cmd.AddCommand(ListRestrictionQueryRolesCmd)
}
