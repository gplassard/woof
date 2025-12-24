package restriction_policies

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteRestrictionPolicyCmd = &cobra.Command{
	Use:   "deleterestrictionpolicy",
	Short: "Delete a restriction policy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/restriction_policy/{resource_id}")
		fmt.Println("OperationID: DeleteRestrictionPolicy")
	},
}

func init() {
	Cmd.AddCommand(DeleteRestrictionPolicyCmd)
}
