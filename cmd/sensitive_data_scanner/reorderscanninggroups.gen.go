package sensitive_data_scanner

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ReorderScanningGroupsCmd = &cobra.Command{
	Use:   "reorderscanninggroups",
	Short: "Reorder Groups",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/sensitive-data-scanner/config")
		fmt.Println("OperationID: ReorderScanningGroups")
	},
}

func init() {
	Cmd.AddCommand(ReorderScanningGroupsCmd)
}
