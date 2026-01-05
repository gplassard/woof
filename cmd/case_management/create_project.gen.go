package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateProjectCmd = &cobra.Command{
	Use: "create-project",

	Short: "Create a project",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.CreateProject(client.NewContext(apiKey, appKey, site), datadogV2.ProjectCreateRequest{})
		cmdutil.HandleError(err, "failed to create-project")

		cmd.Println(cmdutil.FormatJSON(res, "project"))
	},
}

func init() {
	Cmd.AddCommand(CreateProjectCmd)
}
