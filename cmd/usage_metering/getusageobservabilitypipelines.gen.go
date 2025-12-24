package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetUsageObservabilityPipelinesCmd = &cobra.Command{
	Use:   "getusageobservabilitypipelines",
	Short: "Get hourly usage for observability pipelines",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/usage/observability_pipelines")
		fmt.Println("OperationID: GetUsageObservabilityPipelines")
	},
}

func init() {
	Cmd.AddCommand(GetUsageObservabilityPipelinesCmd)
}
