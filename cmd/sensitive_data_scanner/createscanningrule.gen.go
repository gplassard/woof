package sensitive_data_scanner

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateScanningRuleCmd = &cobra.Command{
	Use:   "createscanningrule",
	Short: "Create Scanning Rule",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/sensitive-data-scanner/config/rules")
		fmt.Println("OperationID: CreateScanningRule")
	},
}

func init() {
	Cmd.AddCommand(CreateScanningRuleCmd)
}
