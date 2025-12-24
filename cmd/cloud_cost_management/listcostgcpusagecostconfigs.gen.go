package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCostGCPUsageCostConfigsCmd = &cobra.Command{
	Use:   "listcostgcpusagecostconfigs",
	Short: "List Google Cloud Usage Cost configs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/gcp_uc_config")
		fmt.Println("OperationID: ListCostGCPUsageCostConfigs")
	},
}

func init() {
	Cmd.AddCommand(ListCostGCPUsageCostConfigsCmd)
}
