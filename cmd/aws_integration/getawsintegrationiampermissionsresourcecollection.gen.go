package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAWSIntegrationIAMPermissionsResourceCollectionCmd = &cobra.Command{
	Use:   "getawsintegrationiampermissionsresourcecollection",
	Short: "Get resource collection IAM permissions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/aws/iam_permissions/resource_collection")
		fmt.Println("OperationID: GetAWSIntegrationIAMPermissionsResourceCollection")
	},
}

func init() {
	Cmd.AddCommand(GetAWSIntegrationIAMPermissionsResourceCollectionCmd)
}
