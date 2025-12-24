package cloud_cost_management

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UploadCustomCostsFileCmd = &cobra.Command{
	Use:   "uploadcustomcostsfile",
	Short: "Upload Custom Costs file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/cost/custom_costs")
		fmt.Println("OperationID: UploadCustomCostsFile")
	},
}

func init() {
	Cmd.AddCommand(UploadCustomCostsFileCmd)
}
