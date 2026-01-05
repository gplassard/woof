package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetProjectsCmd = &cobra.Command{
	Use: "get-projects",

	Short: "Get all projects",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.GetProjects(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-projects")

		cmd.Println(cmdutil.FormatJSON(res, "project"))
	},
}

func init() {
	Cmd.AddCommand(GetProjectsCmd)
}
