package ci_visibility_pipelines

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AggregateCIAppPipelineEventsCmd = &cobra.Command{
	Use:   "aggregateciapppipelineevents",
	Short: "Aggregate pipelines events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/ci/pipelines/analytics/aggregate")
		fmt.Println("OperationID: AggregateCIAppPipelineEvents")
	},
}

func init() {
	Cmd.AddCommand(AggregateCIAppPipelineEventsCmd)
}
