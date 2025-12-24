package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListSecurityFiltersCmd = &cobra.Command{
	Use:   "listsecurityfilters",
	Short: "Get all security filters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/configuration/security_filters")
		fmt.Println("OperationID: ListSecurityFilters")
	},
}

func init() {
	Cmd.AddCommand(ListSecurityFiltersCmd)
}
