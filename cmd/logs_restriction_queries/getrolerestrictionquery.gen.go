package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetRoleRestrictionQueryCmd = &cobra.Command{
	Use:   "getrolerestrictionquery",
	Short: "Get restriction query for a given role",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/logs/config/restriction_queries/role/{role_id}")
		fmt.Println("OperationID: GetRoleRestrictionQuery")
	},
}

func init() {
	Cmd.AddCommand(GetRoleRestrictionQueryCmd)
}
