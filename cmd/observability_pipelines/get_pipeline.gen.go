package observability_pipelines

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetPipelineCmd = &cobra.Command{
	Use: "get-pipeline [pipeline_id]",

	Short: "Get a specific pipeline",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.GetPipeline(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-pipeline")

		cmd.Println(cmdutil.FormatJSON(res, "observability_pipelines"))
	},
}

func init() {
	Cmd.AddCommand(GetPipelineCmd)
}
