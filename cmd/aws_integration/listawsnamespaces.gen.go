package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAWSNamespacesCmd = &cobra.Command{
	Use:   "listawsnamespaces",
	Short: "List available namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/aws/available_namespaces")
		fmt.Println("OperationID: ListAWSNamespaces")
	},
}

func init() {
	Cmd.AddCommand(ListAWSNamespacesCmd)
}
