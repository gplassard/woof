package observability_pipelines

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetPipelineCmd = &cobra.Command{
	Use:   "getpipeline [pipeline_id]",
	Short: "Get a specific pipeline",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.GetPipeline(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getpipeline: %v", err)
		}

		cmdutil.PrintJSON(res, "observability_pipelines")
	},
}

func init() {
	Cmd.AddCommand(GetPipelineCmd)
}
