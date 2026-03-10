package llm_observability

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateLLMObsDatasetCmd = &cobra.Command{
	Use: "create-llm-obs-dataset [project_id]",

	Short: "Create an LLM Observability dataset",
	Long: `Create an LLM Observability dataset
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#create-llm-obs-dataset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LLMObsDatasetResponse
		var err error

		var body datadogV2.LLMObsDatasetRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateLLMObsDataset(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-llm-obs-dataset")

		fmt.Println(cmdutil.FormatJSON(res, "datasets"))
	},
}

func init() {

	CreateLLMObsDatasetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateLLMObsDatasetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateLLMObsDatasetCmd)
}
