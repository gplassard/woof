package workflow_automation

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetWorkflowInstanceCmd = &cobra.Command{
	Use:   "getworkflowinstance [workflow_id] [instance_id]",
	Short: "Get a workflow instance",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		res, _, err := api.GetWorkflowInstance(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to getworkflowinstance: %v", err)
		}

		cmdutil.PrintJSON(res, "workflow_automation")
	},
}

func init() {
	Cmd.AddCommand(GetWorkflowInstanceCmd)
}
