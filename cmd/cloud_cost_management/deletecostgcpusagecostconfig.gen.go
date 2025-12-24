package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCostGCPUsageCostConfigCmd = &cobra.Command{
	Use:   "deletecostgcpusagecostconfig",
	Short: "Delete Google Cloud Usage Cost config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cost/gcp_uc_config/{cloud_account_id}")
		fmt.Println("OperationID: DeleteCostGCPUsageCostConfig")
	},
}

func init() {
	Cmd.AddCommand(DeleteCostGCPUsageCostConfigCmd)
}
