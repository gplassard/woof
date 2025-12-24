package ci_visibility_pipelines

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchCIAppPipelineEventsCmd = &cobra.Command{
	Use:   "searchciapppipelineevents",
	Short: "Search pipelines events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/ci/pipelines/events/search")
		fmt.Println("OperationID: SearchCIAppPipelineEvents")
	},
}

func init() {
	Cmd.AddCommand(SearchCIAppPipelineEventsCmd)
}
