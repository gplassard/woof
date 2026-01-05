package workflow_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateWorkflowInstanceCmd = &cobra.Command{
	Use:     "create-workflow-instance [workflow_id] [payload]",
	Aliases: []string{"create-instance"},
	Short:   "Execute a workflow",
	Long: `Execute a workflow
Documentation: https://docs.datadoghq.com/api/latest/workflow-automation/#create-workflow-instance`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.WorkflowInstanceCreateResponse
		var err error

		var body datadogV2.WorkflowInstanceCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err = api.CreateWorkflowInstance(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-workflow-instance")

		cmd.Println(cmdutil.FormatJSON(res, "workflow_automation"))
	},
}

func init() {
	Cmd.AddCommand(CreateWorkflowInstanceCmd)
}
