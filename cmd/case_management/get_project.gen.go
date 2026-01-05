package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetProjectCmd = &cobra.Command{
	Use: "get-project [project_id]",

	Short: "Get the details of a project",
	Long: `Get the details of a project
Documentation: https://docs.datadoghq.com/api/latest/case-management/#get-project`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ProjectResponse
		var err error

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err = api.GetProject(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-project")

		cmd.Println(cmdutil.FormatJSON(res, "project"))
	},
}

func init() {
	Cmd.AddCommand(GetProjectCmd)
}
