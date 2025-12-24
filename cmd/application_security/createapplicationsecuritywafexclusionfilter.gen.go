package application_security

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:   "createapplicationsecuritywafexclusionfilter",
	Short: "Create a WAF exclusion filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/remote_config/products/asm/waf/exclusion_filters")
		fmt.Println("OperationID: CreateApplicationSecurityWafExclusionFilter")
	},
}

func init() {
	Cmd.AddCommand(CreateApplicationSecurityWafExclusionFilterCmd)
}
