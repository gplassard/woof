package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCostAWSCURConfigCmd = &cobra.Command{
	Use:   "createcostawscurconfig",
	Short: "Create Cloud Cost Management AWS CUR config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cost/aws_cur_config")
		fmt.Println("OperationID: CreateCostAWSCURConfig")
	},
}

func init() {
	Cmd.AddCommand(CreateCostAWSCURConfigCmd)
}
