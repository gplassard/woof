package observability_pipelines

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListPipelinesCmd = &cobra.Command{
	Use:   "listpipelines",
	Short: "List pipelines",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/remote_config/products/obs_pipelines/pipelines")
		fmt.Println("OperationID: ListPipelines")
	},
}

func init() {
	Cmd.AddCommand(ListPipelinesCmd)
}
