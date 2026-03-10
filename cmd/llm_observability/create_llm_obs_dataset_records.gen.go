package llm_observability

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateLLMObsDatasetRecordsCmd = &cobra.Command{
	Use: "create-llm-obs-dataset-records [project_id] [dataset_id]",

	Short: "Append records to an LLM Observability dataset",
	Long: `Append records to an LLM Observability dataset
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#create-llm-obs-dataset-records`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LLMObsDatasetRecordsMutationResponse
		var err error

		var body datadogV2.LLMObsDatasetRecordsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateLLMObsDatasetRecords(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to create-llm-obs-dataset-records")

		fmt.Println(cmdutil.FormatJSON(res, "llm_observability"))
	},
}

func init() {

	CreateLLMObsDatasetRecordsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateLLMObsDatasetRecordsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateLLMObsDatasetRecordsCmd)
}
