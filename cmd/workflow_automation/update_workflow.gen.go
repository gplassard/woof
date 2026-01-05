package workflow_automation

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateWorkflowCmd = &cobra.Command{
	Use:     "update-workflow [workflow_id]",
	Aliases: []string{"update"},
	Short:   "Update an existing Workflow",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err := api.UpdateWorkflow(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateWorkflowRequest{})
		cmdutil.HandleError(err, "failed to update-workflow")

		cmd.Println(cmdutil.FormatJSON(res, "workflows"))
	},
}

func init() {
	Cmd.AddCommand(UpdateWorkflowCmd)
}
