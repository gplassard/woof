package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListTagsByMetricNameCmd = &cobra.Command{
	Use:   "listtagsbymetricname",
	Short: "List tags by metric name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/metrics/{metric_name}/all-tags")
		fmt.Println("OperationID: ListTagsByMetricName")
	},
}

func init() {
	Cmd.AddCommand(ListTagsByMetricNameCmd)
}
