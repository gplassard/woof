package observability_pipelines

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreatePipelineCmd = &cobra.Command{
	Use:   "createpipeline",
	Short: "Create a new pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/remote_config/products/obs_pipelines/pipelines")
		fmt.Println("OperationID: CreatePipeline")
	},
}

func init() {
	Cmd.AddCommand(CreatePipelineCmd)
}
