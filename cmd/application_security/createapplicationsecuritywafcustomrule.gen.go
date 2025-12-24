package application_security

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:   "createapplicationsecuritywafcustomrule",
	Short: "Create a WAF custom rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/remote_config/products/asm/waf/custom_rules")
		fmt.Println("OperationID: CreateApplicationSecurityWafCustomRule")
	},
}

func init() {
	Cmd.AddCommand(CreateApplicationSecurityWafCustomRuleCmd)
}
