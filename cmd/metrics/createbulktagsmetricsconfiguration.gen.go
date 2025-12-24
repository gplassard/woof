package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateBulkTagsMetricsConfigurationCmd = &cobra.Command{
	Use:   "createbulktagsmetricsconfiguration",
	Short: "Configure tags for multiple metrics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/metrics/config/bulk-tags")
		fmt.Println("OperationID: CreateBulkTagsMetricsConfiguration")
	},
}

func init() {
	Cmd.AddCommand(CreateBulkTagsMetricsConfigurationCmd)
}
