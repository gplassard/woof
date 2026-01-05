package observability_pipelines

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ValidatePipelineCmd = &cobra.Command{
	Use: "validate-pipeline [payload]",

	Short: "Validate an observability pipeline",
	Long: `Validate an observability pipeline
Documentation: https://docs.datadoghq.com/api/latest/observability-pipelines/#validate-pipeline`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ValidationResponse
		var err error

		var body datadogV2.ObservabilityPipelineSpec
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		res, _, err = api.ValidatePipeline(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-pipeline")

		cmd.Println(cmdutil.FormatJSON(res, "observability_pipelines"))
	},
}

func init() {
	Cmd.AddCommand(ValidatePipelineCmd)
}
