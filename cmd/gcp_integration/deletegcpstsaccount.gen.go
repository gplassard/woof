package gcp_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteGCPSTSAccountCmd = &cobra.Command{
	Use:   "deletegcpstsaccount",
	Short: "Delete an STS enabled GCP Account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integration/gcp/accounts/{account_id}")
		fmt.Println("OperationID: DeleteGCPSTSAccount")
	},
}

func init() {
	Cmd.AddCommand(DeleteGCPSTSAccountCmd)
}
