package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetCustomCostsFileCmd = &cobra.Command{
	Use:   "getcustomcostsfile",
	Short: "Get Custom Costs file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/custom_costs/{file_id}")
		fmt.Println("OperationID: GetCustomCostsFile")
	},
}

func init() {
	Cmd.AddCommand(GetCustomCostsFileCmd)
}
