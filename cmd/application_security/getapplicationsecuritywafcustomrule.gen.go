package application_security

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:   "getapplicationsecuritywafcustomrule",
	Short: "Get a WAF custom rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/asm/waf/custom_rules/{custom_rule_id}")
		fmt.Println("OperationID: GetApplicationSecurityWafCustomRule")
	},
}

func init() {
	Cmd.AddCommand(GetApplicationSecurityWafCustomRuleCmd)
}
