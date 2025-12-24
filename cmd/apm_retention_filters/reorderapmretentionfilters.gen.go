package apm_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ReorderApmRetentionFiltersCmd = &cobra.Command{
	Use:   "reorderapmretentionfilters",
	Short: "Re-order retention filters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/apm/config/retention-filters-execution-order")
		fmt.Println("OperationID: ReorderApmRetentionFilters")
	},
}

func init() {
	Cmd.AddCommand(ReorderApmRetentionFiltersCmd)
}
