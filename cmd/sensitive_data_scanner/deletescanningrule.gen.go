package sensitive_data_scanner

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteScanningRuleCmd = &cobra.Command{
	Use:   "deletescanningrule",
	Short: "Delete Scanning Rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/sensitive-data-scanner/config/rules/{rule_id}")
		fmt.Println("OperationID: DeleteScanningRule")
	},
}

func init() {
	Cmd.AddCommand(DeleteScanningRuleCmd)
}
