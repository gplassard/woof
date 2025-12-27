package observability_pipelines

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdatePipelineCmd = &cobra.Command{
	Use:   "updatepipeline [pipeline_id]",
	Short: "Update a pipeline",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.UpdatePipeline(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ObservabilityPipeline{})
		if err != nil {
			log.Fatalf("failed to updatepipeline: %v", err)
		}

		cmdutil.PrintJSON(res, "observability_pipelines")
	},
}

func init() {
	Cmd.AddCommand(UpdatePipelineCmd)
}
