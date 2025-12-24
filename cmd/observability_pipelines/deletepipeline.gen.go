package observability_pipelines

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeletePipelineCmd = &cobra.Command{
	Use:   "deletepipeline",
	Short: "Delete a pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/remote_config/products/obs_pipelines/pipelines/{pipeline_id}")
		fmt.Println("OperationID: DeletePipeline")
	},
}

func init() {
	Cmd.AddCommand(DeletePipelineCmd)
}
