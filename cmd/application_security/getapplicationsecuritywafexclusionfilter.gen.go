package application_security

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:   "getapplicationsecuritywafexclusionfilter",
	Short: "Get a WAF exclusion filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/asm/waf/exclusion_filters/{exclusion_filter_id}")
		fmt.Println("OperationID: GetApplicationSecurityWafExclusionFilter")
	},
}

func init() {
	Cmd.AddCommand(GetApplicationSecurityWafExclusionFilterCmd)
}
