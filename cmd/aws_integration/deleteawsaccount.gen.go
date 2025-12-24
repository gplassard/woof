package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteAWSAccountCmd = &cobra.Command{
	Use:   "deleteawsaccount",
	Short: "Delete an AWS integration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/integration/aws/accounts/{aws_account_config_id}")
		fmt.Println("OperationID: DeleteAWSAccount")
	},
}

func init() {
	Cmd.AddCommand(DeleteAWSAccountCmd)
}
