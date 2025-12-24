package application_security

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:   "updateapplicationsecuritywafexclusionfilter",
	Short: "Update a WAF exclusion filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/remote_config/products/asm/waf/exclusion_filters/{exclusion_filter_id}")
		fmt.Println("OperationID: UpdateApplicationSecurityWafExclusionFilter")
	},
}

func init() {
	Cmd.AddCommand(UpdateApplicationSecurityWafExclusionFilterCmd)
}
