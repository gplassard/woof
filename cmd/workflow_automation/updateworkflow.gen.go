package workflow_automation

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateWorkflowCmd = &cobra.Command{
	Use:   "updateworkflow [workflow_id]",
	Short: "Update an existing Workflow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err := api.UpdateWorkflow(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateWorkflowRequest{})
		if err != nil {
			log.Fatalf("failed to updateworkflow: %v", err)
		}

		cmdutil.PrintJSON(res, "workflow_automation")
	},
}

func init() {
	Cmd.AddCommand(UpdateWorkflowCmd)
}
