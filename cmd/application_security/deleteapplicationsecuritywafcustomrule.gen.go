package application_security

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteApplicationSecurityWafCustomRuleCmd = &cobra.Command{
	Use:   "deleteapplicationsecuritywafcustomrule",
	Short: "Delete a WAF Custom Rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/remote_config/products/asm/waf/custom_rules/{custom_rule_id}")
		fmt.Println("OperationID: DeleteApplicationSecurityWafCustomRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteApplicationSecurityWafCustomRuleCmd)
}
