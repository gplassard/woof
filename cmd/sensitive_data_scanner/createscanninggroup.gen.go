package sensitive_data_scanner

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateScanningGroupCmd = &cobra.Command{
	Use:   "createscanninggroup",
	Short: "Create Scanning Group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/sensitive-data-scanner/config/groups")
		fmt.Println("OperationID: CreateScanningGroup")
	},
}

func init() {
	Cmd.AddCommand(CreateScanningGroupCmd)
}
