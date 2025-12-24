package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var QueryTimeseriesDataCmd = &cobra.Command{
	Use:   "querytimeseriesdata",
	Short: "Query timeseries data across multiple products",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/query/timeseries")
		fmt.Println("OperationID: QueryTimeseriesData")
	},
}

func init() {
	Cmd.AddCommand(QueryTimeseriesDataCmd)
}
