package observability_pipelines

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeletePipelineCmd = &cobra.Command{
	Use:   "deletepipeline [pipeline_id]",
	Short: "Delete a pipeline",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		_, err := api.DeletePipeline(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletepipeline: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeletePipelineCmd)
}
