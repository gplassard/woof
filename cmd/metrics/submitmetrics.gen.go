package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SubmitMetricsCmd = &cobra.Command{
	Use:   "submitmetrics",
	Short: "Submit metrics",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/series")
		fmt.Println("OperationID: SubmitMetrics")
	},
}

func init() {
	Cmd.AddCommand(SubmitMetricsCmd)
}
