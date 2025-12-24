package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCostAWSCURConfigCmd = &cobra.Command{
	Use:   "deletecostawscurconfig",
	Short: "Delete Cloud Cost Management AWS CUR config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cost/aws_cur_config/{cloud_account_id}")
		fmt.Println("OperationID: DeleteCostAWSCURConfig")
	},
}

func init() {
	Cmd.AddCommand(DeleteCostAWSCURConfigCmd)
}
