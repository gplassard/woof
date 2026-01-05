package workflow_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetWorkflowInstanceCmd = &cobra.Command{
	Use:     "get-workflow-instance [workflow_id] [instance_id]",
	Aliases: []string{"get-instance"},
	Short:   "Get a workflow instance",
	Long: `Get a workflow instance
Documentation: https://docs.datadoghq.com/api/latest/workflow-automation/#get-workflow-instance`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.WorklflowGetInstanceResponse
		var err error

		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err = api.GetWorkflowInstance(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-workflow-instance")

		cmd.Println(cmdutil.FormatJSON(res, "workflow_automation"))
	},
}

func init() {
	Cmd.AddCommand(GetWorkflowInstanceCmd)
}
