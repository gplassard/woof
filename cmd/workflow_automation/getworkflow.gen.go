package workflow_automation

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetWorkflowCmd = &cobra.Command{
	Use:   "getworkflow [workflow_id]",
	Short: "Get an existing Workflow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err := api.GetWorkflow(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getworkflow: %v", err)
		}

		cmdutil.PrintJSON(res, "workflow_automation")
	},
}

func init() {
	Cmd.AddCommand(GetWorkflowCmd)
}
