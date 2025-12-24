package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteCustomCostsFileCmd = &cobra.Command{
	Use:   "deletecustomcostsfile",
	Short: "Delete Custom Costs file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/cost/custom_costs/{file_id}")
		fmt.Println("OperationID: DeleteCustomCostsFile")
	},
}

func init() {
	Cmd.AddCommand(DeleteCustomCostsFileCmd)
}
