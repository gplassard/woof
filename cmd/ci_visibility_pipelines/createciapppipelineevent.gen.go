package ci_visibility_pipelines

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCIAppPipelineEventCmd = &cobra.Command{
	Use:   "createciapppipelineevent",
	Short: "Send pipeline event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/ci/pipeline")
		fmt.Println("OperationID: CreateCIAppPipelineEvent")
	},
}

func init() {
	Cmd.AddCommand(CreateCIAppPipelineEventCmd)
}
