package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCostGCPUsageCostConfigCmd = &cobra.Command{
	Use:   "updatecostgcpusagecostconfig",
	Short: "Update Google Cloud Usage Cost config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/cost/gcp_uc_config/{cloud_account_id}")
		fmt.Println("OperationID: UpdateCostGCPUsageCostConfig")
	},
}

func init() {
	Cmd.AddCommand(UpdateCostGCPUsageCostConfigCmd)
}
