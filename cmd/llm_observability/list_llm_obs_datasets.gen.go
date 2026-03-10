package llm_observability

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLLMObsDatasetsCmd = &cobra.Command{
	Use: "list-llm-obs-datasets [project_id]",

	Short: "List LLM Observability datasets",
	Long: `List LLM Observability datasets
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#list-llm-obs-datasets`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LLMObsDatasetsResponse
		var err error

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListLLMObsDatasets(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-llm-obs-datasets")

		fmt.Println(cmdutil.FormatJSON(res, "datasets"))
	},
}

func init() {

	Cmd.AddCommand(ListLLMObsDatasetsCmd)
}
