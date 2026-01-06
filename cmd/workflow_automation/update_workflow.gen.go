package workflow_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateWorkflowCmd = &cobra.Command{
	Use:     "update-workflow [workflow_id]",
	Aliases: []string{"update"},
	Short:   "Update an existing Workflow",
	Long: `Update an existing Workflow
Documentation: https://docs.datadoghq.com/api/latest/workflow-automation/#update-workflow`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateWorkflowResponse
		var err error

		var body datadogV2.UpdateWorkflowRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateWorkflow(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-workflow")

		cmd.Println(cmdutil.FormatJSON(res, "workflows"))
	},
}

func init() {

	UpdateWorkflowCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateWorkflowCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateWorkflowCmd)
}
