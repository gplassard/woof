package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetActiveBillingDimensionsCmd = &cobra.Command{
	Use:   "getactivebillingdimensions",
	Short: "Get active billing dimensions for cost attribution",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/cost_by_tag/active_billing_dimensions")
		fmt.Println("OperationID: GetActiveBillingDimensions")
	},
}

func init() {
	Cmd.AddCommand(GetActiveBillingDimensionsCmd)
}
