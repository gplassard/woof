package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAWSIntegrationIAMPermissionsCmd = &cobra.Command{
	Use:   "getawsintegrationiampermissions",
	Short: "Get AWS integration IAM permissions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/aws/iam_permissions")
		fmt.Println("OperationID: GetAWSIntegrationIAMPermissions")
	},
}

func init() {
	Cmd.AddCommand(GetAWSIntegrationIAMPermissionsCmd)
}
