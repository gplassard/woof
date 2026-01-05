package observability_pipelines

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ValidatePipelineCmd = &cobra.Command{
	Use: "validate-pipeline",

	Short: "Validate an observability pipeline",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		res, _, err := api.ValidatePipeline(client.NewContext(apiKey, appKey, site), datadogV2.ObservabilityPipelineSpec{})
		cmdutil.HandleError(err, "failed to validate-pipeline")

		cmd.Println(cmdutil.FormatJSON(res, "observability_pipelines"))
	},
}

func init() {
	Cmd.AddCommand(ValidatePipelineCmd)
}
