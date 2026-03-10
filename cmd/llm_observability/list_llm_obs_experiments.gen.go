package llm_observability

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLLMObsExperimentsCmd = &cobra.Command{
	Use: "list-llm-obs-experiments",

	Short: "List LLM Observability experiments",
	Long: `List LLM Observability experiments
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#list-llm-obs-experiments`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LLMObsExperimentsResponse
		var err error

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListLLMObsExperiments(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-llm-obs-experiments")

		fmt.Println(cmdutil.FormatJSON(res, "experiments"))
	},
}

func init() {

	Cmd.AddCommand(ListLLMObsExperimentsCmd)
}
