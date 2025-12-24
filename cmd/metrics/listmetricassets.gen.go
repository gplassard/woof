package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListMetricAssetsCmd = &cobra.Command{
	Use:   "listmetricassets",
	Short: "Related Assets to a Metric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/metrics/{metric_name}/assets")
		fmt.Println("OperationID: ListMetricAssets")
	},
}

func init() {
	Cmd.AddCommand(ListMetricAssetsCmd)
}
