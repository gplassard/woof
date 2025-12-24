package rum_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var OrderRetentionFiltersCmd = &cobra.Command{
	Use:   "orderretentionfilters",
	Short: "Order RUM retention filters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/rum/applications/{app_id}/relationships/retention_filters")
		fmt.Println("OperationID: OrderRetentionFilters")
	},
}

func init() {
	Cmd.AddCommand(OrderRetentionFiltersCmd)
}
