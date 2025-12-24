package ci_visibility_tests

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AggregateCIAppTestEventsCmd = &cobra.Command{
	Use:   "aggregateciapptestevents",
	Short: "Aggregate tests events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/ci/tests/analytics/aggregate")
		fmt.Println("OperationID: AggregateCIAppTestEvents")
	},
}

func init() {
	Cmd.AddCommand(AggregateCIAppTestEventsCmd)
}
