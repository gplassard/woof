package gcp_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateGCPSTSAccountCmd = &cobra.Command{
	Use:   "updategcpstsaccount",
	Short: "Update STS Service Account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integration/gcp/accounts/{account_id}")
		fmt.Println("OperationID: UpdateGCPSTSAccount")
	},
}

func init() {
	Cmd.AddCommand(UpdateGCPSTSAccountCmd)
}
