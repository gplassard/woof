package llm_observability

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteLLMObsProjectsCmd = &cobra.Command{
	Use: "delete-llm-obs-projects",

	Short: "Delete LLM Observability projects",
	Long: `Delete LLM Observability projects
Documentation: https://docs.datadoghq.com/api/latest/llm-observability/#delete-llm-obs-projects`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.LLMObsDeleteProjectsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewLLMObservabilityApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteLLMObsProjects(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to delete-llm-obs-projects")

	},
}

func init() {

	DeleteLLMObsProjectsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	DeleteLLMObsProjectsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(DeleteLLMObsProjectsCmd)
}
