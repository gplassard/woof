package workflow_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateWorkflowCmd = &cobra.Command{
	Use:     "create-workflow",
	Aliases: []string{"create"},
	Short:   "Create a Workflow",
	Long: `Create a Workflow
Documentation: https://docs.datadoghq.com/api/latest/workflow-automation/#create-workflow`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateWorkflowResponse
		var err error

		var body datadogV2.CreateWorkflowRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateWorkflow(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-workflow")

		cmd.Println(cmdutil.FormatJSON(res, "workflows"))
	},
}

func init() {

	CreateWorkflowCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateWorkflowCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateWorkflowCmd)
}
