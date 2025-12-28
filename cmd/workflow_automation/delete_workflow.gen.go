package workflow_automation

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteWorkflowCmd = &cobra.Command{
	Use:   "delete-workflow [workflow_id]",
	
	Short: "Delete an existing Workflow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		_, err := api.DeleteWorkflow(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-workflow: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteWorkflowCmd)
}
