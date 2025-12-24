package sensitive_data_scanner

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListScanningGroupsCmd = &cobra.Command{
	Use:   "listscanninggroups",
	Short: "List Scanning Groups",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/sensitive-data-scanner/config")
		fmt.Println("OperationID: ListScanningGroups")
	},
}

func init() {
	Cmd.AddCommand(ListScanningGroupsCmd)
}
