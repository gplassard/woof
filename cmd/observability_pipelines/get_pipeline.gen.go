package observability_pipelines

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetPipelineCmd = &cobra.Command{
	Use: "get-pipeline [pipeline_id]",

	Short: "Get a specific pipeline",
	Long: `Get a specific pipeline
Documentation: https://docs.datadoghq.com/api/latest/observability-pipelines/#get-pipeline`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ObservabilityPipeline
		var err error

		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetPipeline(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-pipeline")

		fmt.Println(cmdutil.FormatJSON(res, "pipeline"))
	},
}

func init() {

	Cmd.AddCommand(GetPipelineCmd)
}
