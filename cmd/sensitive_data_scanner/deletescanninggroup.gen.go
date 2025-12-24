package sensitive_data_scanner

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteScanningGroupCmd = &cobra.Command{
	Use:   "deletescanninggroup",
	Short: "Delete Scanning Group",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/sensitive-data-scanner/config/groups/{group_id}")
		fmt.Println("OperationID: DeleteScanningGroup")
	},
}

func init() {
	Cmd.AddCommand(DeleteScanningGroupCmd)
}
