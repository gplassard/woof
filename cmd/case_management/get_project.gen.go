package case_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetProjectCmd = &cobra.Command{
	Use:   "get-project [project_id]",
	
	Short: "Get the details of a project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.GetProject(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-project")

		cmdutil.PrintJSON(res, "project")
	},
}

func init() {
	Cmd.AddCommand(GetProjectCmd)
}
