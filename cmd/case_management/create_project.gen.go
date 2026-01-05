package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateProjectCmd = &cobra.Command{
	Use: "create-project [payload]",

	Short: "Create a project",
	Long: `Create a project
Documentation: https://docs.datadoghq.com/api/latest/case-management/#create-project`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ProjectResponse
		var err error

		var body datadogV2.ProjectCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err = api.CreateProject(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-project")

		cmd.Println(cmdutil.FormatJSON(res, "project"))
	},
}

func init() {
	Cmd.AddCommand(CreateProjectCmd)
}
