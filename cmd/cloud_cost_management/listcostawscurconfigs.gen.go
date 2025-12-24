package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCostAWSCURConfigsCmd = &cobra.Command{
	Use:   "listcostawscurconfigs",
	Short: "List Cloud Cost Management AWS CUR configs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/aws_cur_config")
		fmt.Println("OperationID: ListCostAWSCURConfigs")
	},
}

func init() {
	Cmd.AddCommand(ListCostAWSCURConfigsCmd)
}
