package application_security

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListApplicationSecurityWafExclusionFiltersCmd = &cobra.Command{
	Use:   "listapplicationsecuritywafexclusionfilters",
	Short: "List all WAF exclusion filters",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/asm/waf/exclusion_filters")
		fmt.Println("OperationID: ListApplicationSecurityWafExclusionFilters")
	},
}

func init() {
	Cmd.AddCommand(ListApplicationSecurityWafExclusionFiltersCmd)
}
