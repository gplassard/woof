package restriction_policies

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetRestrictionPolicyCmd = &cobra.Command{
	Use:   "getrestrictionpolicy",
	Short: "Get a restriction policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/restriction_policy/{resource_id}")
		fmt.Println("OperationID: GetRestrictionPolicy")
	},
}

func init() {
	Cmd.AddCommand(GetRestrictionPolicyCmd)
}
