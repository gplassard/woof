package application_security

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListApplicationSecurityWAFCustomRulesCmd = &cobra.Command{
	Use:   "listapplicationsecuritywafcustomrules",
	Short: "List all WAF custom rules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/asm/waf/custom_rules")
		fmt.Println("OperationID: ListApplicationSecurityWAFCustomRules")
	},
}

func init() {
	Cmd.AddCommand(ListApplicationSecurityWAFCustomRulesCmd)
}
