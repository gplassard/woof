package workflow_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListWorkflowInstancesCmd = &cobra.Command{
	Use:     "list-workflow-instances [workflow_id]",
	Aliases: []string{"list-instances"},
	Short:   "List workflow instances",
	Long: `List workflow instances
Documentation: https://docs.datadoghq.com/api/latest/workflow-automation/#list-workflow-instances`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.WorkflowListInstancesResponse
		var err error

		optionalParams := datadogV2.NewListWorkflowInstancesOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-length") {
			val, _ := cmd.Flags().GetInt64("page-length")
			optionalParams.WithPageNumber(val)
		}

		api := datadogV2.NewWorkflowAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListWorkflowInstances(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to list-workflow-instances")

		cmd.Println(cmdutil.FormatJSON(res, "workflow_automation"))
	},
}

func init() {

	ListWorkflowInstancesCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListWorkflowInstancesCmd.Flags().Int64("page-length", 0, "Specific page number to return.")

	Cmd.AddCommand(ListWorkflowInstancesCmd)
}
