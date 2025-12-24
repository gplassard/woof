package observability_pipelines

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdatePipelineCmd = &cobra.Command{
	Use:   "updatepipeline",
	Short: "Update a pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/remote_config/products/obs_pipelines/pipelines/{pipeline_id}")
		fmt.Println("OperationID: UpdatePipeline")
	},
}

func init() {
	Cmd.AddCommand(UpdatePipelineCmd)
}
