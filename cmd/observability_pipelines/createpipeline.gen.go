package observability_pipelines

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreatePipelineCmd = &cobra.Command{
	Use:   "createpipeline",
	Short: "Create a new pipeline",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.CreatePipeline(client.NewContext(apiKey, appKey, site), datadogV2.ObservabilityPipelineSpec{})
		if err != nil {
			log.Fatalf("failed to createpipeline: %v", err)
		}

		cmdutil.PrintJSON(res, "observability_pipelines")
	},
}

func init() {
	Cmd.AddCommand(CreatePipelineCmd)
}
