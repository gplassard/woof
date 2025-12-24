package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateAWSAccountCmd = &cobra.Command{
	Use:   "createawsaccount",
	Short: "Create an AWS integration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integration/aws/accounts")
		fmt.Println("OperationID: CreateAWSAccount")
	},
}

func init() {
	Cmd.AddCommand(CreateAWSAccountCmd)
}
