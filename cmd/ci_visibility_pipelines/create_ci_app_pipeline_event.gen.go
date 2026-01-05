package ci_visibility_pipelines

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateCIAppPipelineEventCmd = &cobra.Command{
	Use: "create-ci-app-pipeline-event [payload]",

	Short: "Send pipeline event",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res interface{}
		var err error

		var body datadogV2.CIAppCreatePipelineEventRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCIVisibilityPipelinesApi(client.NewAPIClient())
		res, _, err = api.CreateCIAppPipelineEvent(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-ci-app-pipeline-event")

		cmd.Println(cmdutil.FormatJSON(res, "ci_visibility_pipelines"))
	},
}

func init() {
	Cmd.AddCommand(CreateCIAppPipelineEventCmd)
}
