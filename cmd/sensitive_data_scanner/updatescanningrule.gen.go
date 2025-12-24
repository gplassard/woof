package sensitive_data_scanner

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateScanningRuleCmd = &cobra.Command{
	Use:   "updatescanningrule",
	Short: "Update Scanning Rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/sensitive-data-scanner/config/rules/{rule_id}")
		fmt.Println("OperationID: UpdateScanningRule")
	},
}

func init() {
	Cmd.AddCommand(UpdateScanningRuleCmd)
}
