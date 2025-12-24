package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteBulkTagsMetricsConfigurationCmd = &cobra.Command{
	Use:   "deletebulktagsmetricsconfiguration",
	Short: "Delete tags for multiple metrics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/metrics/config/bulk-tags")
		fmt.Println("OperationID: DeleteBulkTagsMetricsConfiguration")
	},
}

func init() {
	Cmd.AddCommand(DeleteBulkTagsMetricsConfigurationCmd)
}
