package workflow_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteWorkflowCmd = &cobra.Command{
	Use:     "delete-workflow [workflow_id]",
	Aliases: []string{"delete"},
	Short:   "Delete an existing Workflow",
	Long: `Delete an existing Workflow
Documentation: https://docs.datadoghq.com/api/latest/workflow-automation/#delete-workflow`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteWorkflow(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-workflow")

	},
}

func init() {

	Cmd.AddCommand(DeleteWorkflowCmd)
}
