package workflow_automation

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListWorkflowInstancesCmd = &cobra.Command{
	Use:     "list-workflow-instances [workflow_id]",
	Aliases: []string{"list-instances"},
	Short:   "List workflow instances",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err := api.ListWorkflowInstances(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-workflow-instances")

		cmd.Println(cmdutil.FormatJSON(res, "workflow_automation"))
	},
}

func init() {
	Cmd.AddCommand(ListWorkflowInstancesCmd)
}
