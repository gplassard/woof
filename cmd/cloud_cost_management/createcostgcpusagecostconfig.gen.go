package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCostGCPUsageCostConfigCmd = &cobra.Command{
	Use:   "createcostgcpusagecostconfig",
	Short: "Create Google Cloud Usage Cost config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cost/gcp_uc_config")
		fmt.Println("OperationID: CreateCostGCPUsageCostConfig")
	},
}

func init() {
	Cmd.AddCommand(CreateCostGCPUsageCostConfigCmd)
}
