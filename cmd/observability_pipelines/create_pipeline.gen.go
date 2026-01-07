package observability_pipelines

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreatePipelineCmd = &cobra.Command{
	Use: "create-pipeline",

	Short: "Create a new pipeline",
	Long: `Create a new pipeline
Documentation: https://docs.datadoghq.com/api/latest/observability-pipelines/#create-pipeline`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ObservabilityPipeline
		var err error

		var body datadogV2.ObservabilityPipelineSpec
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreatePipeline(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-pipeline")

		fmt.Println(cmdutil.FormatJSON(res, "observability_pipelines"))
	},
}

func init() {

	CreatePipelineCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreatePipelineCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreatePipelineCmd)
}
