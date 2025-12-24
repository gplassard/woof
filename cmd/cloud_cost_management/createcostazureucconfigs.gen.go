package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCostAzureUCConfigsCmd = &cobra.Command{
	Use:   "createcostazureucconfigs",
	Short: "Create Cloud Cost Management Azure configs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/cost/azure_uc_config")
		fmt.Println("OperationID: CreateCostAzureUCConfigs")
	},
}

func init() {
	Cmd.AddCommand(CreateCostAzureUCConfigsCmd)
}
