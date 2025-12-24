package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCostAzureUCConfigCmd = &cobra.Command{
	Use:   "getcostazureucconfig",
	Short: "Get cost Azure UC config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/azure_uc_config/{cloud_account_id}")
		fmt.Println("OperationID: GetCostAzureUCConfig")
	},
}

func init() {
	Cmd.AddCommand(GetCostAzureUCConfigCmd)
}
