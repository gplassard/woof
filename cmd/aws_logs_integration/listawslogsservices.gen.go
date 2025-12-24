package aws_logs_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAWSLogsServicesCmd = &cobra.Command{
	Use:   "listawslogsservices",
	Short: "Get list of AWS log ready services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/integration/aws/logs/services")
		fmt.Println("OperationID: ListAWSLogsServices")
	},
}

func init() {
	Cmd.AddCommand(ListAWSLogsServicesCmd)
}
