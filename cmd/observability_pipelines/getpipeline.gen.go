package observability_pipelines

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetPipelineCmd = &cobra.Command{
	Use:   "getpipeline",
	Short: "Get a specific pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/obs_pipelines/pipelines/{pipeline_id}")
		fmt.Println("OperationID: GetPipeline")
	},
}

func init() {
	Cmd.AddCommand(GetPipelineCmd)
}
