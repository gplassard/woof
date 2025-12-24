package rum_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteRumMetricCmd = &cobra.Command{
	Use:   "deleterummetric",
	Short: "Delete a rum-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/rum/config/metrics/{metric_id}")
		fmt.Println("OperationID: DeleteRumMetric")
	},
}

func init() {
	Cmd.AddCommand(DeleteRumMetricCmd)
}
