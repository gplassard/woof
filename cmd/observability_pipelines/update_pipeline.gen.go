package observability_pipelines

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdatePipelineCmd = &cobra.Command{
	Use: "update-pipeline [pipeline_id] [payload]",

	Short: "Update a pipeline",
	Long: `Update a pipeline
Documentation: https://docs.datadoghq.com/api/latest/observability-pipelines/#update-pipeline`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ObservabilityPipeline
		var err error

		var body datadogV2.ObservabilityPipeline
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		res, _, err = api.UpdatePipeline(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-pipeline")

		cmd.Println(cmdutil.FormatJSON(res, "observability_pipelines"))
	},
}

func init() {
	Cmd.AddCommand(UpdatePipelineCmd)
}
