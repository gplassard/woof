package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCostAzureUCConfigCmd = &cobra.Command{
	Use:   "deletecostazureucconfig",
	Short: "Delete Cloud Cost Management Azure config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cost/azure_uc_config/{cloud_account_id}")
		fmt.Println("OperationID: DeleteCostAzureUCConfig")
	},
}

func init() {
	Cmd.AddCommand(DeleteCostAzureUCConfigCmd)
}
