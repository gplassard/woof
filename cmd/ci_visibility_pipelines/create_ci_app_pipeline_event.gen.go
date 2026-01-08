package ci_visibility_pipelines

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCIAppPipelineEventCmd = &cobra.Command{
	Use: "create-ci-app-pipeline-event",

	Short: "Send pipeline event",
	Long: `Send pipeline event
Documentation: https://docs.datadoghq.com/api/latest/ci-visibility-pipelines/#create-ci-app-pipeline-event`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res interface{}
		var err error

		var body datadogV2.CIAppCreatePipelineEventRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCIAppPipelineEvent(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-ci-app-pipeline-event")

		fmt.Println(cmdutil.FormatJSON(res, "c_i_app_pipeline_event"))
	},
}

func init() {

	CreateCIAppPipelineEventCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCIAppPipelineEventCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCIAppPipelineEventCmd)
}
