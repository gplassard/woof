package observability_pipelines

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreatePipelineCmd = &cobra.Command{
	Use: "create-pipeline",

	Short: "Create a new pipeline",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.CreatePipeline(client.NewContext(apiKey, appKey, site), datadogV2.ObservabilityPipelineSpec{})
		cmdutil.HandleError(err, "failed to create-pipeline")

		cmd.Println(cmdutil.FormatJSON(res, "observability_pipelines"))
	},
}

func init() {
	Cmd.AddCommand(CreatePipelineCmd)
}
