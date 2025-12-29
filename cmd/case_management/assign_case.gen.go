package case_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AssignCaseCmd = &cobra.Command{
	Use:   "assign-case [case_id]",
	
	Short: "Assign case",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.AssignCase(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseAssignRequest{})
		cmdutil.HandleError(err, "failed to assign-case")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {
	Cmd.AddCommand(AssignCaseCmd)
}
