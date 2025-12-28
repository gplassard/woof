package observability_pipelines

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreatePipelineCmd = &cobra.Command{
	Use:   "create-pipeline",
	Short: "Create a new pipeline",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.CreatePipeline(client.NewContext(apiKey, appKey, site), datadogV2.ObservabilityPipelineSpec{})
		if err != nil {
			log.Fatalf("failed to create-pipeline: %v", err)
		}

		cmdutil.PrintJSON(res, "observability_pipelines")
	},
}

func init() {
	Cmd.AddCommand(CreatePipelineCmd)
}
