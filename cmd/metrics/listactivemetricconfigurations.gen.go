package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListActiveMetricConfigurationsCmd = &cobra.Command{
	Use:   "listactivemetricconfigurations",
	Short: "List active tags and aggregations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/metrics/{metric_name}/active-configurations")
		fmt.Println("OperationID: ListActiveMetricConfigurations")
	},
}

func init() {
	Cmd.AddCommand(ListActiveMetricConfigurationsCmd)
}
