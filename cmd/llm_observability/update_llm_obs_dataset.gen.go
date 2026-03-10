package llm_observability

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateLLMObsDatasetCmd = &cobra.Command{
	Use: "update-llm-obs-dataset [project_id] [dataset_id]",

	Short: "Update an LLM Observability dataset",
	Long: `Update an LLM Observability dataset
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#update-llm-obs-dataset`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LLMObsDatasetResponse
		var err error

		var body datadogV2.LLMObsDatasetUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateLLMObsDataset(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-llm-obs-dataset")

		fmt.Println(cmdutil.FormatJSON(res, "datasets"))
	},
}

func init() {

	UpdateLLMObsDatasetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateLLMObsDatasetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateLLMObsDatasetCmd)
}
