package workflow_automation

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateWorkflowInstanceCmd = &cobra.Command{
	Use:     "create-workflow-instance [workflow_id]",
	Aliases: []string{"create-instance"},
	Short:   "Execute a workflow",
	Long: `Execute a workflow
Documentation: https://docs.datadoghq.com/api/latest/workflow-automation/#create-workflow-instance`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.WorkflowInstanceCreateResponse
		var err error

		var body datadogV2.WorkflowInstanceCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateWorkflowInstance(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-workflow-instance")

		fmt.Println(cmdutil.FormatJSON(res, "workflow_automation"))
	},
}

func init() {

	CreateWorkflowInstanceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateWorkflowInstanceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateWorkflowInstanceCmd)
}
