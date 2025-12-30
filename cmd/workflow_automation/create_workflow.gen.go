package workflow_automation

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateWorkflowCmd = &cobra.Command{
	Use:     "create-workflow",
	Aliases: []string{"create"},
	Short:   "Create a Workflow",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err := api.CreateWorkflow(client.NewContext(apiKey, appKey, site), datadogV2.CreateWorkflowRequest{})
		cmdutil.HandleError(err, "failed to create-workflow")

		cmd.Println(cmdutil.FormatJSON(res, "workflows"))
	},
}

func init() {
	Cmd.AddCommand(CreateWorkflowCmd)
}
