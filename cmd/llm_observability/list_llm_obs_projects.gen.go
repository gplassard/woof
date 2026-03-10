package llm_observability

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListLLMObsProjectsCmd = &cobra.Command{
	Use: "list-llm-obs-projects",

	Short: "List LLM Observability projects",
	Long: `List LLM Observability projects
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#list-llm-obs-projects`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.LLMObsProjectsResponse
		var err error

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListLLMObsProjects(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-llm-obs-projects")

		fmt.Println(cmdutil.FormatJSON(res, "projects"))
	},
}

func init() {

	Cmd.AddCommand(ListLLMObsProjectsCmd)
}
