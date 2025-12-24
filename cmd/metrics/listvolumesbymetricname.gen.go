package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListVolumesByMetricNameCmd = &cobra.Command{
	Use:   "listvolumesbymetricname",
	Short: "List distinct metric volumes by metric name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/metrics/{metric_name}/volumes")
		fmt.Println("OperationID: ListVolumesByMetricName")
	},
}

func init() {
	Cmd.AddCommand(ListVolumesByMetricNameCmd)
}
