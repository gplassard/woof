package observability_pipelines

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ValidatePipelineCmd = &cobra.Command{
	Use:   "validatepipeline",
	Short: "Validate an observability pipeline",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/remote_config/products/obs_pipelines/pipelines/validate")
		fmt.Println("OperationID: ValidatePipeline")
	},
}

func init() {
	Cmd.AddCommand(ValidatePipelineCmd)
}
