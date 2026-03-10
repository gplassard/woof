package llm_observability

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateLLMObsExperimentCmd = &cobra.Command{
	Use: "update-llm-obs-experiment [experiment_id]",

	Short: "Update an LLM Observability experiment",
	Long: `Update an LLM Observability experiment
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#update-llm-obs-experiment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LLMObsExperimentResponse
		var err error

		var body datadogV2.LLMObsExperimentUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateLLMObsExperiment(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-llm-obs-experiment")

		fmt.Println(cmdutil.FormatJSON(res, "experiments"))
	},
}

func init() {

	UpdateLLMObsExperimentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateLLMObsExperimentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateLLMObsExperimentCmd)
}
