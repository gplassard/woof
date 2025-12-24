package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAWSIntegrationIAMPermissionsStandardCmd = &cobra.Command{
	Use:   "getawsintegrationiampermissionsstandard",
	Short: "Get AWS integration standard IAM permissions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/aws/iam_permissions/standard")
		fmt.Println("OperationID: GetAWSIntegrationIAMPermissionsStandard")
	},
}

func init() {
	Cmd.AddCommand(GetAWSIntegrationIAMPermissionsStandardCmd)
}
