package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSuppressionVersionHistoryCmd = &cobra.Command{
	Use:   "getsuppressionversionhistory",
	Short: "Get a suppression's version history",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security_monitoring/configuration/suppressions/{suppression_id}/version_history")
		fmt.Println("OperationID: GetSuppressionVersionHistory")
	},
}

func init() {
	Cmd.AddCommand(GetSuppressionVersionHistoryCmd)
}
