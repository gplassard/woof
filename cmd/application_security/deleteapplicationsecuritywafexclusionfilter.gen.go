package application_security

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:   "deleteapplicationsecuritywafexclusionfilter",
	Short: "Delete a WAF exclusion filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/remote_config/products/asm/waf/exclusion_filters/{exclusion_filter_id}")
		fmt.Println("OperationID: DeleteApplicationSecurityWafExclusionFilter")
	},
}

func init() {
	Cmd.AddCommand(DeleteApplicationSecurityWafExclusionFilterCmd)
}
