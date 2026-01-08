package observability_pipelines

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ValidatePipelineCmd = &cobra.Command{
	Use: "validate-pipeline",

	Short: "Validate an observability pipeline",
	Long: `Validate an observability pipeline
Documentation: https://docs.datadoghq.com/api/latest/observability-pipelines/#validate-pipeline`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ValidationResponse
		var err error

		var body datadogV2.ObservabilityPipelineSpec
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewObservabilityPipelinesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ValidatePipeline(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-pipeline")

		fmt.Println(cmdutil.FormatJSON(res, "pipeline"))
	},
}

func init() {

	ValidatePipelineCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ValidatePipelineCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ValidatePipelineCmd)
}
