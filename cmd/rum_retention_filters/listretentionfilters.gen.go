package rum_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListRetentionFiltersCmd = &cobra.Command{
	Use:   "listretentionfilters",
	Short: "Get all RUM retention filters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/rum/applications/{app_id}/retention_filters")
		fmt.Println("OperationID: ListRetentionFilters")
	},
}

func init() {
	Cmd.AddCommand(ListRetentionFiltersCmd)
}
