package workflow_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CancelWorkflowInstanceCmd = &cobra.Command{
	Use:     "cancel-workflow-instance [workflow_id] [instance_id]",
	Aliases: []string{"cancel-instance"},
	Short:   "Cancel a workflow instance",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err := api.CancelWorkflowInstance(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to cancel-workflow-instance")

		cmd.Println(cmdutil.FormatJSON(res, "workflow_automation"))
	},
}

func init() {
	Cmd.AddCommand(CancelWorkflowInstanceCmd)
}
