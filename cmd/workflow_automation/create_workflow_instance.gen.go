package workflow_automation

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateWorkflowInstanceCmd = &cobra.Command{
	Use:   "create-workflow-instance [workflow_id]",
	Aliases: []string{ "create-instance", },
	Short: "Execute a workflow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err := api.CreateWorkflowInstance(client.NewContext(apiKey, appKey, site), args[0], datadogV2.WorkflowInstanceCreateRequest{})
		cmdutil.HandleError(err, "failed to create-workflow-instance")

		cmd.Println(cmdutil.FormatJSON(res, "workflow_automation"))
	},
}

func init() {
	Cmd.AddCommand(CreateWorkflowInstanceCmd)
}
