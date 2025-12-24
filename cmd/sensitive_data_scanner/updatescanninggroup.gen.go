package sensitive_data_scanner

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateScanningGroupCmd = &cobra.Command{
	Use:   "updatescanninggroup",
	Short: "Update Scanning Group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/sensitive-data-scanner/config/groups/{group_id}")
		fmt.Println("OperationID: UpdateScanningGroup")
	},
}

func init() {
	Cmd.AddCommand(UpdateScanningGroupCmd)
}
