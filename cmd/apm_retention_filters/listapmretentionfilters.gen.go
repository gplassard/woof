package apm_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListApmRetentionFiltersCmd = &cobra.Command{
	Use:   "listapmretentionfilters",
	Short: "List all APM retention filters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/apm/config/retention-filters")
		fmt.Println("OperationID: ListApmRetentionFilters")
	},
}

func init() {
	Cmd.AddCommand(ListApmRetentionFiltersCmd)
}
