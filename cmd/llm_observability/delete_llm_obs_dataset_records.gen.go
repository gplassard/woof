package llm_observability

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteLLMObsDatasetRecordsCmd = &cobra.Command{
	Use: "delete-llm-obs-dataset-records [project_id] [dataset_id]",

	Short: "Delete LLM Observability dataset records",
	Long: `Delete LLM Observability dataset records
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#delete-llm-obs-dataset-records`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.LLMObsDeleteDatasetRecordsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteLLMObsDatasetRecords(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to delete-llm-obs-dataset-records")

	},
}

func init() {

	DeleteLLMObsDatasetRecordsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DeleteLLMObsDatasetRecordsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DeleteLLMObsDatasetRecordsCmd)
}
