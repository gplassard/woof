package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateAWSAccountCmd = &cobra.Command{
	Use:   "updateawsaccount",
	Short: "Update an AWS integration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/integration/aws/accounts/{aws_account_config_id}")
		fmt.Println("OperationID: UpdateAWSAccount")
	},
}

func init() {
	Cmd.AddCommand(UpdateAWSAccountCmd)
}
