package rum

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AggregateRUMEventsCmd = &cobra.Command{
	Use:   "aggregaterumevents",
	Short: "Aggregate RUM events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/rum/analytics/aggregate")
		fmt.Println("OperationID: AggregateRUMEvents")
	},
}

func init() {
	Cmd.AddCommand(AggregateRUMEventsCmd)
}
