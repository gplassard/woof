package rum_metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateRumMetricCmd = &cobra.Command{
	Use:   "createrummetric",
	Short: "Create a rum-based metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/rum/config/metrics")
		fmt.Println("OperationID: CreateRumMetric")
	},
}

func init() {
	Cmd.AddCommand(CreateRumMetricCmd)
}
