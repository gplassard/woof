package restriction_policies

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateRestrictionPolicyCmd = &cobra.Command{
	Use:   "updaterestrictionpolicy",
	Short: "Update a restriction policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/restriction_policy/{resource_id}")
		fmt.Println("OperationID: UpdateRestrictionPolicy")
	},
}

func init() {
	Cmd.AddCommand(UpdateRestrictionPolicyCmd)
}
