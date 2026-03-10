package llm_observability

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLLMObsDatasetRecordsCmd = &cobra.Command{
	Use: "list-llm-obs-dataset-records [project_id] [dataset_id]",

	Short: "List LLM Observability dataset records",
	Long: `List LLM Observability dataset records
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#list-llm-obs-dataset-records`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LLMObsDatasetRecordsListResponse
		var err error

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListLLMObsDatasetRecords(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to list-llm-obs-dataset-records")

		fmt.Println(cmdutil.FormatJSON(res, "llm_observability"))
	},
}

func init() {

	Cmd.AddCommand(ListLLMObsDatasetRecordsCmd)
}
