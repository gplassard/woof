package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetRuleVersionHistoryCmd = &cobra.Command{
	Use:   "getruleversionhistory",
	Short: "Get a rule's version history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/rules/{rule_id}/version_history")
		fmt.Println("OperationID: GetRuleVersionHistory")
	},
}

func init() {
	Cmd.AddCommand(GetRuleVersionHistoryCmd)
}
