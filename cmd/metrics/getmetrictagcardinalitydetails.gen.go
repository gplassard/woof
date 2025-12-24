package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetMetricTagCardinalityDetailsCmd = &cobra.Command{
	Use:   "getmetrictagcardinalitydetails",
	Short: "Get tag key cardinality details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/metrics/{metric_name}/tag-cardinalities")
		fmt.Println("OperationID: GetMetricTagCardinalityDetails")
	},
}

func init() {
	Cmd.AddCommand(GetMetricTagCardinalityDetailsCmd)
}
