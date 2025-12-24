package ci_visibility_pipelines

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCIAppPipelineEventsCmd = &cobra.Command{
	Use:   "listciapppipelineevents",
	Short: "Get a list of pipelines events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/ci/pipelines/events")
		fmt.Println("OperationID: ListCIAppPipelineEvents")
	},
}

func init() {
	Cmd.AddCommand(ListCIAppPipelineEventsCmd)
}
