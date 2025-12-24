package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCostGCPUsageCostConfigCmd = &cobra.Command{
	Use:   "getcostgcpusagecostconfig",
	Short: "Get Google Cloud Usage Cost config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/gcp_uc_config/{cloud_account_id}")
		fmt.Println("OperationID: GetCostGCPUsageCostConfig")
	},
}

func init() {
	Cmd.AddCommand(GetCostGCPUsageCostConfigCmd)
}
