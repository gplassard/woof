package workflow_automation

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetWorkflowCmd = &cobra.Command{
	Use:     "get-workflow [workflow_id]",
	Aliases: []string{"get"},
	Short:   "Get an existing Workflow",
	Long: `Get an existing Workflow
Documentation: https://docs.datadoghq.com/api/latest/workflow-automation/#get-workflow`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetWorkflowResponse
		var err error

		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetWorkflow(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-workflow")

		fmt.Println(cmdutil.FormatJSON(res, "workflow"))
	},
}

func init() {

	Cmd.AddCommand(GetWorkflowCmd)
}
