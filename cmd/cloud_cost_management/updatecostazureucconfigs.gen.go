package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCostAzureUCConfigsCmd = &cobra.Command{
	Use:   "updatecostazureucconfigs",
	Short: "Update Cloud Cost Management Azure config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/cost/azure_uc_config/{cloud_account_id}")
		fmt.Println("OperationID: UpdateCostAzureUCConfigs")
	},
}

func init() {
	Cmd.AddCommand(UpdateCostAzureUCConfigsCmd)
}
