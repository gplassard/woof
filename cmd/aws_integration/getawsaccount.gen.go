package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAWSAccountCmd = &cobra.Command{
	Use:   "getawsaccount",
	Short: "Get an AWS integration by config ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/aws/accounts/{aws_account_config_id}")
		fmt.Println("OperationID: GetAWSAccount")
	},
}

func init() {
	Cmd.AddCommand(GetAWSAccountCmd)
}
