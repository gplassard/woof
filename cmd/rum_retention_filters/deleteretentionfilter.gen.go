package rum_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteRetentionFilterCmd = &cobra.Command{
	Use:   "deleteretentionfilter",
	Short: "Delete a RUM retention filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/rum/applications/{app_id}/retention_filters/{rf_id}")
		fmt.Println("OperationID: DeleteRetentionFilter")
	},
}

func init() {
	Cmd.AddCommand(DeleteRetentionFilterCmd)
}
