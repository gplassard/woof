package workflow_automation

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteWorkflowCmd = &cobra.Command{
	Use:     "delete-workflow [workflow_id]",
	Aliases: []string{"delete"},
	Short:   "Delete an existing Workflow",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		_, err := api.DeleteWorkflow(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-workflow")

	},
}

func init() {
	Cmd.AddCommand(DeleteWorkflowCmd)
}
