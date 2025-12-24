package application_security

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:   "updateapplicationsecuritywafcustomrule",
	Short: "Update a WAF Custom Rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/remote_config/products/asm/waf/custom_rules/{custom_rule_id}")
		fmt.Println("OperationID: UpdateApplicationSecurityWafCustomRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateApplicationSecurityWafCustomRuleCmd)
}
