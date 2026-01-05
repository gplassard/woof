package workflow_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateWorkflowCmd = &cobra.Command{
	Use:     "create-workflow [payload]",
	Aliases: []string{"create"},
	Short:   "Create a Workflow",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateWorkflowResponse
		var err error

		var body datadogV2.CreateWorkflowRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err = api.CreateWorkflow(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-workflow")

		cmd.Println(cmdutil.FormatJSON(res, "workflows"))
	},
}

func init() {
	Cmd.AddCommand(CreateWorkflowCmd)
}
