package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCostAzureUCConfigsCmd = &cobra.Command{
	Use:   "listcostazureucconfigs",
	Short: "List Cloud Cost Management Azure configs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/azure_uc_config")
		fmt.Println("OperationID: ListCostAzureUCConfigs")
	},
}

func init() {
	Cmd.AddCommand(ListCostAzureUCConfigsCmd)
}
