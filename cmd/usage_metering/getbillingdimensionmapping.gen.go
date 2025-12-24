package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetBillingDimensionMappingCmd = &cobra.Command{
	Use:   "getbillingdimensionmapping",
	Short: "Get billing dimension mapping for usage endpoints",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/usage/billing_dimension_mapping")
		fmt.Println("OperationID: GetBillingDimensionMapping")
	},
}

func init() {
	Cmd.AddCommand(GetBillingDimensionMappingCmd)
}
