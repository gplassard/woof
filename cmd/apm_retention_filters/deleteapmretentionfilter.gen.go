package apm_retention_filters

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteApmRetentionFilterCmd = &cobra.Command{
	Use:   "deleteapmretentionfilter",
	Short: "Delete a retention filter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/apm/config/retention-filters/{filter_id}")
		fmt.Println("OperationID: DeleteApmRetentionFilter")
	},
}

func init() {
	Cmd.AddCommand(DeleteApmRetentionFilterCmd)
}
