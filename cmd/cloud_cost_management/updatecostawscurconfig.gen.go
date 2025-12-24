package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCostAWSCURConfigCmd = &cobra.Command{
	Use:   "updatecostawscurconfig",
	Short: "Update Cloud Cost Management AWS CUR config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/cost/aws_cur_config/{cloud_account_id}")
		fmt.Println("OperationID: UpdateCostAWSCURConfig")
	},
}

func init() {
	Cmd.AddCommand(UpdateCostAWSCURConfigCmd)
}
