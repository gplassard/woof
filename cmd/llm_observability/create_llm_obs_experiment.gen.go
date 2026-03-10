package llm_observability

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateLLMObsExperimentCmd = &cobra.Command{
	Use: "create-llm-obs-experiment",

	Short: "Create an LLM Observability experiment",
	Long: `Create an LLM Observability experiment
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#create-llm-obs-experiment`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LLMObsExperimentResponse
		var err error

		var body datadogV2.LLMObsExperimentRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateLLMObsExperiment(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-llm-obs-experiment")

		fmt.Println(cmdutil.FormatJSON(res, "experiments"))
	},
}

func init() {

	CreateLLMObsExperimentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateLLMObsExperimentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateLLMObsExperimentCmd)
}
