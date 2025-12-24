package spans

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AggregateSpansCmd = &cobra.Command{
	Use:   "aggregatespans",
	Short: "Aggregate spans",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/spans/analytics/aggregate")
		fmt.Println("OperationID: AggregateSpans")
	},
}

func init() {
	Cmd.AddCommand(AggregateSpansCmd)
}
