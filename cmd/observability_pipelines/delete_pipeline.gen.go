package observability_pipelines

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeletePipelineCmd = &cobra.Command{
	Use: "delete-pipeline [pipeline_id]",

	Short: "Delete a pipeline",
	Long: `Delete a pipeline
Documentation: https://docs.datadoghq.com/api/latest/observability-pipelines/#delete-pipeline`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeletePipeline(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-pipeline")

	},
}

func init() {

	Cmd.AddCommand(DeletePipelineCmd)
}
