package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCustomCostsFilesCmd = &cobra.Command{
	Use:   "listcustomcostsfiles",
	Short: "List Custom Costs files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost/custom_costs")
		fmt.Println("OperationID: ListCustomCostsFiles")
	},
}

func init() {
	Cmd.AddCommand(ListCustomCostsFilesCmd)
}
