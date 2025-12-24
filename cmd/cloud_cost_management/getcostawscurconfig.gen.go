package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCostAWSCURConfigCmd = &cobra.Command{
	Use:   "getcostawscurconfig",
	Short: "Get cost AWS CUR config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/aws_cur_config/{cloud_account_id}")
		fmt.Println("OperationID: GetCostAWSCURConfig")
	},
}

func init() {
	Cmd.AddCommand(GetCostAWSCURConfigCmd)
}
